# Development

This section walks through how to build and test the plugin.

## Requirements

| software | version | link                         |
|:---------|:--------|:-----------------------------|
| kubectl  |  ????   | [download][download-kubectl] |
| go       |  v1.22  | [download-go]                |

## IDE Setup (VS Code)

1. Install the [VS Code Go extension]
1. Import the KubeRay workspace configuration by using the file `kuberay.code-workspace` in the root
   of the KubeRay git repo: "File" -> "Open Workspace from File" -> "kuberay.code-workspace".

Setting up workspace configuration is required because KubeRay contains multiple Go modules. See the
[VS Code Go documentation] for details.

### Automated Tests

Run unit tests with the following command.

```console
cd kubectl-plugin
go test ./pkg/... -race -parallel 4
```

### Manual Tests

Run the plugin commands against a Kubernetes cluster where you have deployed the [KubeRay Operator].
Here's some example usage.

```console
❯ kubectl ray --context my-context version
kubectl ray plugin version: development
KubeRay operator version: v1.2.2@sha256:cc8ce713f3b4be3c72cca1f63ee78e3733bc7283472ecae367b47a128f7e4478

❯ kubectl ray --context my-context create cluster -n exoplanet portia
Created Ray Cluster: portia

❯ kubectl ray --context my-context get cluster -n exoplanet
NAME      NAMESPACE   DESIRED WORKERS   AVAILABLE WORKERS   CPUS   GPUS   TPUS   MEMORY   AGE
bianca    exoplanet   1                 1                   16     2      0      96Gi     21d
portia    exoplanet   1                 1                   4      0      0      8Gi      29m

❯ kubectl ray --context my-context delete -n exoplanet raycluster/portia
Are you sure you want to delete raycluster portia? (y/yes/n/no) y
Delete raycluster portia
```

[download-kubectl]: https://kubernetes.io/docs/tasks/tools/install-kubectl/
[download-go]: https://golang.org/dl/
[VS Code Go extension]: https://marketplace.visualstudio.com/items?itemName=golang.Go
[VS Code Go documentation]: https://github.com/golang/vscode-go/blob/master/README.md#setting-up-your-workspace
[KubeRay Operator]: https://docs.ray.io/en/latest/cluster/kubernetes/index.html
