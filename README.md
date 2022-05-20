# Web App



# App Details
A sample appliccation to demo CI/CD of an application with autoscaling using CirlceCI, ArgoCD and Kubernetes

.

# Proposed Stack

Golang


# Top-level directory layout

    📦k8s-argo-ci
        ┣ 📜.gitignore
        ┣ 📜Dockerfile
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

# How to setup project using docker

1.  Create the Docker image
    ```
    docker build -t your-username/appdemo:v1 .
    ```
    Example
    ```
    docker build -t adefemi171/appdemo:v1 .
    ```

3.  Run the docker file
    ```
    docker run -p 3000:3000 --name appdemo: YOURNAME/appdemo:v1
    ```
    Example
    ```
    docker run -p 3000:3000 --name appdemo adefemi171/appdemo:v1
    ```

4.  To confirm if your container was built, run:
    ```
    docker ps -a
    ```

5.  View the image in your localhost port by typing ```localhost:3000``` in your browser




NOTE: [Docker](https://docs.docker.com/get-docker/), [Golang](https://golang.org/doc/install) needs to be installed to run this application.