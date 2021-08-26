# Go-deploy
An end to end delivery tool that creates deployments at the push of a button.

# Problem Statment

As the name suggests this is an easy to use deployment tool that updates the state of
the cluster to the desired state leveraging ArgoCD , Docker, Kubernetes API’s and the
Go to create an end to end deployment tool.
The users have options to update the current state of the cluster based on the desired
state of the git-repo, or to also replace the cluster state with an image state of their
choice by passing the right tag.
They will be provided with security options like the sealed-secrets and to store important
secrets on the cluster and other secrets can be stored locally in the keychain access to
make sure authentication is key.
Go-deploy as the name suggests is a server written in Go and a tool that makes end to
end deployments as simple as a click of a button.

# Flow-chart
![Go-deploy-flowchart](https://user-images.githubusercontent.com/48808456/130895976-922b553e-0816-453a-8a81-4a68d22d5582.jpg)
