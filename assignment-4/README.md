# Simple-Go-Web-Application built to deploy using Kubernetes Artifacts

We can either use the public image or local image in the deployment.yaml file.

As in this section we are using local image and if the pod fails with the ErrImagePull status, please execute the below commands:

1) minikube docker-env
o/p => 
export DOCKER_TLS_VERIFY=”1"
export DOCKER_HOST=”tcp://172.17.0.2:2376"
export DOCKER_CERT_PATH=”/home/user/.minikube/certs”
export MINIKUBE_ACTIVE_DOCKERD=”minikube”# To point your shell to minikube’s docker-daemon, run:
# eval $(minikube -p minikube docker-env)

2) eval $(minikube -p minikube docker-env)

3) docker build . -t test

4) And recreate the jobs once again by using `kubectl apply -f .`
You shall see the pods are running.

All this procedures are well tested on from my side.
If you still got some error please contact to `shubhamjagdhane1010@gmail.com`