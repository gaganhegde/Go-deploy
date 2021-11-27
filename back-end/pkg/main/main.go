package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	argocd "github.com/gaganhegde/go-deploy/pkg/argoCD"
	deployments "github.com/gaganhegde/go-deploy/pkg/deployment"
	namespaces "github.com/gaganhegde/go-deploy/pkg/namespace"
	"github.com/gaganhegde/go-deploy/pkg/services"
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

var counter int
var mutex = &sync.Mutex{}

type file map[string]interface{}

//The service() is directed to by the /service end-point to create service files.
func service(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	services.CreateService(body)

}

//The deployment() function is directed to by the /deployment end point to create deployment files
func deployment(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	deployments.CreateDeployments(body)
}

//The namespace() function is directed to by the /deployment end point to create deployment files
func namespace(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	namespaces.CreateNamespace(body)
}

//The argoCD() function is directed to create the appropriate ArgoCD CRD
func argoCD(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	argocd.CreateArgoCD(body)
}
func main() {
	http.HandleFunc("/service", service)
	http.HandleFunc("/deployment", deployment)
	http.HandleFunc("/namespace", namespace)
	http.HandleFunc("/argocd", argoCD)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
