# Web App



# App Details
A sample appliccation to demo CI/CD of an application with autoscaling using CirlceCI, ArgoCD and Kubernetes

.

# Proposed Stack

Golang


# Top-level directory layout

    📦k8s-argo-ci
        ┣ 📜.gitignore
        ┣ 📜go.mod
        ┣ 📜main_test.go
        ┣ 📜main.go
        ┣ 📜README.md


# How to setup project and run locally

### Fork the repository 

```
git clone https://github.com/{github_username}/k8s-argo-ci.git
```
### Checking Out
The App is deployed on the ``` main ``` branch you will need to checkout to the app branch using:

```
git checkout main
```


### Run the application

1. Run the the application using
```
go run main.go
```
3. In your browser, navigate to

```
http://localhost:3000/
```




NOTE: [Golang](https://golang.org/doc/install) needs to be installed to run this application