# Doing Kubernetes CI/CD right with ArgoCD, Kustomize and Github Actions

<img src="assets/banner.png">

## Why we need CI/CD?
**CI/CD** reduces the gap between your system infrastructure and your development process thus allowing new features to be tested, deployed and delivered much more quickly. CI/CD is a must because:
- Manual deployments can be error prone
- Allow graceful changes like `Blue-Green Deployments` and `Canary Testing`

## CI/CD and Kubernetes
For deploying cloud-native applications safely and efficiently on kubernetes, we need something different from traditional **CI/CD pipelines**. It should follow the same declarative way used by `k8s` to manage applications
- Injecting kubernetes objects through a declarative document as YAML or JSON
- Kubernetes Operators processing endlessly evaluating the difference between the submitted objects and their real state in the cluster.

## Why GitOps?
GitOps is especially suited to deploy cloud-native applications on Kubernetes following the above methodology while retaining it's underlying philosophy:
- Use git repo as single source of truth
- changes are triggered with a git commit
- When the state of repo changes, same changes are enforced inside the cluster.

## Our GitOps flow

<img src="assets/flowchart.png">

**Tools**<br>
- **ArgoCD** - it lives inside our cluster and is responsible for deploying changes commited to repo
- **Github Actions** - it is responsible for building & pushing image to docker hub, and commit the latest image tag back to infra repo.
- **Kustomize** - it describes our application infrastructure and specification.

