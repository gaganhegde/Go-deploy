package services

import (
	"encoding/json"
	"fmt"

	"github.com/redhat-developer/kam/pkg/files"
	"github.com/redhat-developer/kam/pkg/utility"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Service1 struct contains the json structure of service yaml file
type Service1 struct {
	Path      string   `json:"servicePath"`
	Version   string   `json:"apiVersion"`
	Metadata  data     `json:"metadata"`
	SSelector selector `json:"SSelector"`
	Sports    []ports  `json:"Sports"`
}
type ports struct {
	Protocol   string `json:"protocol"`
	Port       int    `json:"port"`
	Targetport int    `json:"targetport"`
}
type selector struct {
	App  string `json:"app"`
	Role string `json:"role"`
	Tier string `json:"tier"`
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

type finalservice struct {
	path     string
	typemeta metav1.TypeMeta
	metdata  metav1.ObjectMeta
}
type file map[string]interface{}

//CreateService creates the service file for us
func CreateService(body []byte) {
	Serv := Service1{}
	FServ := finalservice{}
	err := json.Unmarshal(body, &Serv)
	if err != nil {
		fmt.Println(err)
	}
	FServ.path = Serv.Path

	if Serv.Version != "" {
		typeMeta := utility.CreateTypeMeta(Serv.Version, "Service")
		FServ.typemeta = *typeMeta
	}
	if Serv.Metadata.Name != "" {
		metadata := utility.CreateMetaData(Serv.Metadata.Name, Serv.Metadata.Labels.App, Serv.Metadata.Labels.Role, Serv.Metadata.Labels.Tier)
		FServ.metdata = *metadata
	}
	serviceFile := buildService(FServ)
	resources := file{}
	resources["service.yaml"] = serviceFile
	files.CreateFiles(FServ.path, resources)

}

func buildService(service finalservice) *corev1.Service {
	return &corev1.Service{
		TypeMeta:   service.typemeta,
		ObjectMeta: service.metdata,
	}
}
