# Web App



# App Details
A sample appliccation to demo CI/CD of an application with autoscaling using CirlceCI, ArgoCD and Kubernetes

.

# Proposed Stack

Golang


# Top-level directory layout

    📦k8s-argo-ci
        📦.circleci
            ┗ 📜config.yml
        ┣ 📜.gitignore
        ┣ 📜Dockerfile
        ┣ 📜go.mod
        ┣ 📜main_test.go
        ┣ 📜main.go
        ┣ 📜README.md


## Folder structure

1. [.circleci](https://github.com/adefemi171/k8s-argo-ci/tree/main/.circleci): Contains configuration for deployment

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


### Connect your repository to CircleCI
- To run the CI part of the application you will need to link your repository with [CircleCI](https://circleci.com/) by [creating an account and connecting your repo](https://circleci.com/signup/). 
- To add the environment variables. Click on Dashboard, Select your pipeline, on the far right of your screen go to Project Settings -> Environment variables. Add `DOCKER_PASSWORD`, `DOCKER_USERNAME` and `GITHUB_PERSONAL_TOKEN`.
- Update the [.circleci/config.yml](https://github.com/adefemi171/k8s-argo-ci/blob/main/.circleci/config.yml) from line [32](https://github.com/adefemi171/k8s-argo-ci/blob/291e74a727be999407c83501d8b78640473802b8/.circleci/config.yml#L32) to line [53](https://github.com/adefemi171/k8s-argo-ci/blob/291e74a727be999407c83501d8b78640473802b8/.circleci/config.yml#L53) to reflect yur changes and your docker image.


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