package argocd

//Pipeline is the basic pipeleine struct
type Pipeline struct {
	Stages  []string `json:"stages" protobuf:"bytes,1,opt,name=stages"`
	Build   Build    `json:"build" protobuf:"bytes,1,opt,name=build"`
	Publish Publish  `json:"publish" protobuf:"bytes,1,opt,name=publish"`
	Deploy  Deploy   `json:"deploy" protobuf:"bytes,1,opt,name=deploy"`
}

// Build represents desired application state. Contains link to repository with application definition and additional parameters link definition revision.
type Build struct {
	// Source is a reference to the location ksonnet application definition
	Stage string `json:"stage" protobuf:"bytes,1,opt,name=stage"`
	// Destination overrides the kubernetes server and namespace defined in the environment ksonnet app.yaml
	Image Image `json:"image" protobuf:"bytes,2,name=image"`
	// Project is a application project name. Empty name means that application belongs to 'default' project.
	Script []string `json:"script" protobuf:"bytes,3,name=script"`
	// SyncPolicy controls when a sync will be performed
	Artifacts Artifacts `json:"artifacts,omitempty" protobuf:"bytes,4,name=artifacts"`
	// SyncPolicy controls when a sync will be performed
	Variables Variables `json:"variables,omitempty" protobuf:"bytes,4,name=variables"`
}

// Deploy represents desired application state. Contains link to repository with application definition and additional parameters link definition revision.
type Deploy struct {
	// Source is a reference to the location ksonnet application definition
	Stage string `json:"stage" protobuf:"bytes,1,opt,name=stage"`
	// Destination overrides the kubernetes server and namespace defined in the environment ksonnet app.yaml
	Image Image `json:"image" protobuf:"bytes,2,name=image"`
	// Project is a application project name. Empty name means that application belongs to 'default' project.
	BeforeScript []string `json:"before_script" protobuf:"bytes,3,name=script"`
	// Project is a application project name. Empty name means that application belongs to 'default' project.
	Script []string `json:"script" protobuf:"bytes,3,name=script"`
	// SyncPolicy controls when a sync will be performed
	Only []string `json:"only,omitempty" protobuf:"bytes,4,name=only"`
}

// Publish represents desired application state. Contains link to repository with application definition and additional parameters link definition revision.
type Publish struct {
	// Source is a reference to the location ksonnet application definition
	Stage string `json:"stage" protobuf:"bytes,1,opt,name=stage"`
	// Destination overrides the kubernetes server and namespace defined in the environment ksonnet app.yaml
	Image Image `json:"image" protobuf:"bytes,2,name=image"`
	// Project is a application project name. Empty name means that application belongs to 'default' project.
	Script []string `json:"script" protobuf:"bytes,3,name=script"`
	// SyncPolicy controls when a sync will be performed
	Dependencies []string `json:"dependencies,omitempty" protobuf:"bytes,4,name=dependencies"`
	// SyncPolicy controls when a sync will be performed
	Only []string `json:"only,omitempty" protobuf:"bytes,4,name=only"`
}

//Image source throws light on the type of source and the revision
type Image struct {
	// RepoURL is the repository URL of the application manifests
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
}

// Artifacts contains deployment destination information
type Artifacts struct {
	// Server overrides the environment server value in the ksonnet app.yaml
	Paths []string `json:"paths,omitempty" protobuf:"bytes,1,opt,name=paths"`
}

// Variables controls when a sync will be performed in response to updates in git
type Variables struct {
	// Automated will keep an application synced to the target revision
	CGOEnabled int `json:"CGO_ENABLED,omitempty" protobuf:"bytes,1,opt,name=CGO_ENABLED"`
}
