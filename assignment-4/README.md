# Assignmet-4
Simple-Go-Web-Application deploy using Kubernetes Artifacts

## Prerequisites
```
Docker
Minikube
Kubectl
```

## Installation

1. [Docker](https://www.docker.com/) &nbsp; &nbsp; : If you don't have docker install, please install it from [Here](https://docs.docker.com/engine/install/).
2. [Minikube]() : You can install Minikube on docker as well on virtual machine[Here](https://v1-18.docs.kubernetes.io/docs/tasks/tools/install-minikube/) is link to install Minikube using virtual machine.
3. [Kubectl]() &nbsp; &nbsp;: You can install Kubectl from [Here](https://kubernetes.io/docs/tasks/tools/).



## Getting Started

1. Clone the project.
2. Go to assiggment-4/ repository
3. docker build . -t test
4. Start the minikube by executing `minikube start` command
5. Execute `kubectl apply -f .`
6. Grab the IP address of the Minikube by executing `kubectl cluster-info` command.
7. Enter `kubectl get svc` command to get the port of the service.
8. Open browser and enter `http://minikube-ip:port/pucsd`, replace minikube-ip and port from the output of the 6th and 7th command respectively.

## Built With Stack

* [Docker](https://www.docker.com/) &nbsp;&nbsp;&nbsp;&nbsp;- &nbsp;Used to build the custom images
* [Minikube](https://v1-18.docs.kubernetes.io/docs/tasks/tools/install-minikube/) - &nbsp;Used start the cluster for kubernetes
* [Kubectl](https://kubernetes.io/docs/tasks/tools/) &nbsp;&nbsp;&nbsp;- &nbsp;Used to execute `.yaml` files