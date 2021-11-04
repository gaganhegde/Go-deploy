package namespace

import (
	"encoding/json"
	"fmt"

	argo "github.com/redhat-developer/kam/pkg/argoCD/argostructs"
	pipelines "github.com/redhat-developer/kam/pkg/argoCD/pipelines"
	"github.com/redhat-developer/kam/pkg/files"
	"github.com/redhat-developer/kam/pkg/utility"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//ArgoCD struct contains the json structure of service yaml file
type ArgoCD struct {
	GitlabRepo      string `json:"GitlabRepo"`
	Username        string `json:"Username"`
	LocalPath       string `json:"LocalPath"`
	ApplicationName string `json:"ApplicationName"`
	Namespace       string `json:"Namespace"`
}

type finalNamespace struct {
	path     string
	typemeta metav1.TypeMeta
	metdata  metav1.ObjectMeta
	spec     argo.ApplicationSpec
}
type file map[string]interface{}

//CreateArgoCD creates the service file for us
func CreateArgoCD(body []byte) {
	Nm := ArgoCD{}
	FNm := finalNamespace{}
	err := json.Unmarshal(body, &Nm)
	if err != nil {
		fmt.Println(err)
	}
	typeMeta := utility.CreateTypeMeta("Application", "argoproj.io/v1alpha1")
	FNm.typemeta = *typeMeta
	metadata := createMetaData(Nm.ApplicationName, "argocd")
	FNm.metdata = *metadata
	spec := createApplicationSpec(Nm.GitlabRepo, Nm.Namespace)
	FNm.spec = *spec
	deploymentFile := buildDeployments(FNm)
	resources := file{}
	pipelineFile := buildPipelines(Nm.Username, Nm.GitlabRepo, Nm.Namespace)
	fmt.Println("This is the pipeline file", pipelineFile)
	resources["argoCD.yaml"] = deploymentFile
	resources[".gitlab-ci.yml"] = pipelineFile
	files.CreateFiles(Nm.LocalPath, resources)
	// createPipelineFille(Nm.LocalPath, pipelineFile)

}
func buildPipelines(username string, gitrepo string, namespace string) *pipelines.Pipeline {
	return &pipelines.Pipeline{
		Stages: []string{"builld", "publish"},
		Build: pipelines.Build{
			Stage: "build",
			Image: pipelines.Image{
				Name: "golang:1.13.1",
			},
			Script: []string{"go build -o main main.go"},
			Artifacts: pipelines.Artifacts{
				Paths: []string{"main"},
			},
			Variables: pipelines.Variables{
				CGOEnabled: 0,
			},
		},
		Publish: pipelines.Publish{
			Stage: "publish",
			Image: pipelines.Image{
				Name: "gcr.io/kaniko-project/executor:debug",
			},
			Script:       []string{"echo \"{\"auths\":{\"$CI_REGISTRY\":{\"username\":\"$CI_REGISTRY_USER\",\"password\":\"$CI_REGISTRY_PASSWORD\"}}}\" > /kaniko/.docker/config.json", "/kaniko/executor --context $CI_PROJECT_DIR --dockerfile ./Dockerfile --destination $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA"},
			Dependencies: []string{"build"},
			Only:         []string{"master"},
		},
		Deploy: pipelines.Deploy{
			Stage: "deploy",
			Image: pipelines.Image{
				Name: "alpine:3.8",
			},
			BeforeScript: []string{
				"apk add --no-cache git curl bash",
				"curl -s \"https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh\"  | bash",
				"mv kustomize /usr/local/bin/",
				"git remote set-url origin https://${CI_USERNAME}:${CI_PUSH_TOKEN}@" + gitrepo,
				"git config --global user.email \"gitlab@gitlab.com\"",
				"git config --global user.name \"GitLab CI/CD\"",
			},
			Script: []string{
				"git checkout -B master",
				"cd manifests/" + namespace,
				"kustomize edit set image $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA",
				"cat kustomization.yaml",
				"git commit -am '[skip ci] DEV image update'",
				"git push origin master",
			},
			Only: []string{"master"},
		},
	}
}
func createApplicationSpec(repo string, namespace string) *argo.ApplicationSpec {
	return &argo.ApplicationSpec{
		Project: "default",
		Source: argo.ApplicationSource{
			RepoURL:        repo,
			Path:           "manifest/" + namespace,
			TargetRevision: "HEAD",
		},
		Destination: argo.ApplicationDestination{
			Server:    "https://kubernetes.default.svc",
			Namespace: namespace,
		},
		SyncPolicy: &argo.SyncPolicy{
			Automated: &argo.SyncPolicyAutomated{
				Prune:    true,
				SelfHeal: true,
			},
		},
	}
}
func createMetaData(name string, namespace string) *metav1.ObjectMeta {
	return &metav1.ObjectMeta{
		Name:      name,
		Namespace: namespace,
	}
}

func buildDeployments(deployment finalNamespace) *argo.Application {
	return &argo.Application{
		TypeMeta:   deployment.typemeta,
		ObjectMeta: deployment.metdata,
		Spec:       deployment.spec,
	}
}
