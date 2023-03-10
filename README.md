# QuickBuild
QuickBuild is a controller that quickly creates kubernetes deployment resources. It can install deployment and service resources at one time by specifying parameters, and monitor their running status.

## Description

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
make docker-build docker-push IMG=<some-registry>/quickbuild:tag
```
	
3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/quickbuild:tag
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

## Records
The prototype of QuickBuild CR is as follows:
```shell
$ kubectl get qb -A -oyaml                                  
apiVersion: v1
items:
- apiVersion: app.mark8s.io/v1
  kind: QuickBuild
  metadata:
    annotations:
      kubectl.kubernetes.io/last-applied-configuration: |
        {"apiVersion":"app.mark8s.io/v1","kind":"QuickBuild","metadata":{"annotations":{},"labels":{"app.kubernetes.io/created-by":"quickbuild","app.kubernetes.io/instance":"quickbuild-sample","app.kubernetes.io/managed-by":"kustomize","app.kubernetes.io/name":"quickbuild","app.kubernetes.io/part-of":"quickbuild"},"name":"quickbuild-sample","namespace":"demo"},"spec":{"enableService":true,"image":"registry.cn-shenzhen.aliyuncs.com/solarmesh/httpbin","name":"quick","namespace":"demo","port":80,"replicas":1}}
    creationTimestamp: "2023-01-10T07:53:30Z"
    generation: 1
    labels:
      app.kubernetes.io/created-by: quickbuild
      app.kubernetes.io/instance: quickbuild-sample
      app.kubernetes.io/managed-by: kustomize
      app.kubernetes.io/name: quickbuild
      app.kubernetes.io/part-of: quickbuild
    name: quickbuild-sample
    namespace: demo
    resourceVersion: "5048829"
    uid: f773598d-2a27-4cd1-9f1b-3b9d717f8ef9
  spec:
    enableService: true
    image: registry.cn-shenzhen.aliyuncs.com/solarmesh/httpbin
    name: quick
    namespace: demo
    port: 80
    replicas: 1
  status:
    serviceIp: 10.96.198.161
    status: AllReady
kind: List
metadata:
  resourceVersion: ""
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

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

