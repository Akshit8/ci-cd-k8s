## TODO
- create a single node microk8s cluster ✅
- setup kubectl for it ✅
- deploy argo-cd in it ✅
- connect the repo with ssh key-pair ✅
- create a simple golang app for testing ✅
- dockerize the app ✅
- write github action for it(to push on docker registry && commit change to repo ) ✅
- Test it all! ✅

## Installing Microk8s
[download-docs](https://microk8s.io/docs)
```bash
# install the binary with snap
sudo snap install microk8s --classic --channel=1.19

# add your current user to the group and gain access to the .kube caching directory
sudo usermod -a -G microk8s $USER
sudo chown -f -R $USER ~/.kube

# MicroK8s uses a namespaced kubectl command to prevent conflicts with any existing installs of kubectl
# to use kubectl add alias to ~/.bash_aliases
alias kubectl='microk8s kubectl'

# MicroK8s uses the minimum of components for a pure, lightweight Kubernetes
# it is recommended to add DNS management to facilitate communication between services
microk8s enable dns
```

## deploying app in argo-app namespace
```bash
kubectl create ns argo-app

kubectl -n argo-app kustomize ci-cd-k8s/infra/ | kubectl -n argo-app apply -f -

kubectl get all -n argo-app

kubectl port-forward --address 0.0.0.0 service/argo-app -n argo-app 3000:3000
```

## deploy argocd in the cluster
```bash
kubectl create ns argo

kubectl -n argo apply -f ci-cd-k8s/argo/install.yml  

kubectl port-forward --address 0.0.0.0 service/argocd-server -n argo 3001:443
```

## Connecting repo to argo cd
```bash
ssh-keygen

settings > repostiories > connect repo using ssh

name - argo-app
url - git@github.com:/Akshit8/ci-cd-k8s
ssh key private

click on connect
```

## creating a new application
```bash
app name - argo-app
project - default
sync policy - automatic(untick prune resources and self heal)

tick auto create namespace

repo url

revision HEAD

path - infra

cluster url - https://kubernetes.default.svc
namesapce - argo-app

the auto kustomize path
```