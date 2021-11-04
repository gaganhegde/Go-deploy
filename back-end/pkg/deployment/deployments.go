package deployment

import (
	"encoding/json"
	"fmt"

	"github.com/redhat-developer/kam/pkg/files"
	"github.com/redhat-developer/kam/pkg/utility"
	v1 "k8s.io/api/apps/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Service1 struct contains the json structure of service yaml file
type Deployment1 struct {
	Path      string      `json:"servicePath"`
	Version   string      `json:"apiVersion"`
	Metadata  data        `json:"metadata"`
	Container []container `json:"Containers"`
}
type container struct {
	Name      string `json:"Name"`
	Image     string `json:"Image"`
	Portvalue int    `json:"PortValue"`
}
type data struct {
	Name   string `json:"Name"`
	Labels label  `json:"Labels"`
}

type label struct {
	App  string `json:"App"`
	Role string `json:"Role"`
	Tier string `json:"Tier"`
}

type finaldeployment struct {
	path     string
	typemeta metav1.TypeMeta
	metdata  metav1.ObjectMeta
}
type file map[string]interface{}

//CreateDeployments creates the service file for us
func CreateDeployments(body []byte) {
	Dep := Deployment1{}
	FDep := finaldeployment{}
	err := json.Unmarshal(body, &Dep)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("This is the deployment parameters", Dep)
	FDep.path = Dep.Path

	if Dep.Version != "" {
		typeMeta := utility.CreateTypeMeta(Dep.Version, "Service")
		FDep.typemeta = *typeMeta
	}
	if Dep.Metadata.Name != "" {
		metadata := utility.CreateMetaData(Dep.Metadata.Name, Dep.Metadata.Labels.App, Dep.Metadata.Labels.Role, Dep.Metadata.Labels.Tier)
		FDep.metdata = *metadata
	}
	deploymentFile := buildDeployments(FDep)
	resources := file{}
	resources["deployment.yaml"] = deploymentFile
	files.CreateFiles(FDep.path, resources)

}

func buildDeployments(deployment finaldeployment) *v1.Deployment {
	return &v1.Deployment{
		TypeMeta:   deployment.typemeta,
		ObjectMeta: deployment.metdata,
	}
}
