# kube-operator-sidecar
通过kubebuilder实现创建自动注入sidecar的kube-operator

## Description
自我实践通过kubebuilder实现一个operator扩展，主要功能是模拟Openkruise的Sidecarset功能。为Pod自动注入sidecar


## 项目构建过程

### 安装 Kubebuilder
官方地址: https://book.kubebuilder.io/quick-start.html

```sh
curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)
chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
```

### 初始化项目
```shell
# 创建项目目录
mkdir -P $GOPATH/src/Kube-Operator-Sidecar

# 进入项目目录
cd $GOPATH/src/Kube-Operator-Sidecar

# 初始化项目, domain 指定了后续注册 CRD 对象的 Group 域名。
kubebuilder init --domain=kruise.io 
```

### 创建API(CRD)和Controller框架
```shell
kubebuilder create api --group apps --version v1alpha1 --kind SidecarSet --namespace=false
```
操作：

运行 "kubebuilder create api --group apps --version v1alpha1 --kind SidecarSet --namespace=false"
实际上不仅会创建 API，也就是 CRD，还会生成 Controller 的框架。

参数解读：
- group 加上之前的 domian 即此 CRD 的 Group: apps.kruise.io；
- version 一般分三种，按社区标准：
- v1alpha1: 此 api 不稳定，CRD 可能废弃、字段可能随时调整，不要依赖；
- v1beta1: api 已稳定，会保证向后兼容，特性可能会调整；
- v1: api 和特性都已稳定；
- kind: 此 CRD 的类型，类似于社区原生的 Service 的概念；
- namespaced: 此 CRD 是全局唯一还是 namespace 唯一，类似 node 和 Pod。

它的参数基本可以分为两类。group, version, kind 基本上对应了 CRD 元信息的三个重要组成部分。这里给出了一些常见的标准，大家实际使用的时候可以参考一下。namespaced 则用于指定我们刚刚创建的 CRD 时全局唯一的（如 node）还是 namespace 唯一的（如 Pod）。这里用了 false，即指定 SidecarSet 为全局唯一的。

效果解读：
生成了 CRD 和 controller 的框架，后面需要手工填充代码。
实际结果如下图所示：
```shell
├── api
│   └── v1alpha1
│       ├── groupversion_info.go
│       ├── sidecarset_types.go
│       └── zz_generated.deepcopy.go
├── bin
│   └── controller-gen
├── config
...
├── controllers
│   ├── sidecarset_controller.go
│   └── suite_test.go
├── Dockerfile
├── go.mod
├── go.sum
├── hack
│   └── boilerplate.go.txt
├── main.go
├── Makefile
├── PROJECT
└── README.md

```
我们重点关注下面两个文件。
- sidecarset_types.go 就是定义 CRD 的地方，需要我们填充。
- sidecarset_controller.go 则用于定义 controller，同样需要进行填充。

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:
	
```sh
make docker-build docker-push IMG=<some-registry>/kube-operator-sidecar:tag
```
	
3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/kube-operator-sidecar:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller to the cluster:

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/) 
which provides a reconcile function responsible for synchronizing resources untile the desired state is reached on the cluster 

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

