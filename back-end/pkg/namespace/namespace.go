package namespace

import (
	"encoding/json"
	"fmt"

	"github.com/redhat-developer/kam/pkg/files"
	"github.com/redhat-developer/kam/pkg/utility"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Namespace1 struct contains the json structure of service yaml file
type Namespace1 struct {
	Path    string `json:"Path"`
	Name    string `json:"Name"`
	Version string `json:"Version"`
}

type finalNamespace struct {
	path     string
	typemeta metav1.TypeMeta
	metdata  metav1.ObjectMeta
}
type file map[string]interface{}

//CreateDeployments creates the service file for us
func CreateNamespace(body []byte) {
	Nm := Namespace1{}
	FNm := finalNamespace{}
	err := json.Unmarshal(body, &Nm)
	if err != nil {
		fmt.Println(err)
	}
	if Nm.Version != "" {
		typeMeta := utility.CreateTypeMeta(Nm.Version, "Namespace")
		FNm.typemeta = *typeMeta
	}
	if Nm.Name != "" {
		metadata := utility.CreateMetaData(Nm.Name, "", "", "")
		FNm.metdata = *metadata
	}
	fmt.Println("tis is the Fnm", Nm)
	deploymentFile := buildDeployments(FNm)
	resources := file{}
	resources["namespace.yaml"] = deploymentFile
	files.CreateFiles(Nm.Path, resources)

}

func buildDeployments(deployment finalNamespace) *corev1.Namespace {
	return &corev1.Namespace{
		TypeMeta:   deployment.typemeta,
		ObjectMeta: deployment.metdata,
	}
}
