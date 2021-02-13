## TODO
- create a single node microk8s cluster ✅
- setup kubectl for it ✅
- deploy argo-cd in it
- connect the repo with ssh key-pair
- create a simple golang app for testing
- dockerize the app
- write github action for it(to push on docker registry && commit change to repo )

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
```