# Hi Operator
The hi operator creates a Pod that display a custom log message
whenever a HiMessage CRD resource is populated  on the cluster

## Software requirements
* kubebuilder 2.2.0+
* Go 1.12.14
* Kubernetes cluster 1.7+ for testing purpose
* Kubectl


## Create the project

### Cloning the repository
```
git clone https://github.com/nelvadas/HiOperator.git
```

### Fist time (optionnal)
 ```
  $ cd HiOperator
  $ go mod init github.com/nelvadas/HiOperator
    go: creating new go.mod: module github.com/nelvadas/HiOperator
```
Initialize the kubebuilder project
```
$ kubebuilder init --domain abyster.com
```


## Create a HiMessage resource and API
```
$ kubebuilder create api --group messaging --version v1 --kind HiMessage --resource --controller
```
The *--resource * option generates the CRD without prompting
The *--controller *" options generates a controller without prompting

## Install the CRD in the cluster
```
$ make install
customresourcedefinition.apiextensions.k8s.io/himessages.messaging.abyster.com created
```

The make install goal deploy the crd resources on the cluster

```
$ kubectl get crd -o wide
NAME                               CREATED AT
himessages.messaging.abyster.com   2020-01-01T14:40:23Z
```


## Working with himessage CRD instances

Check if we have any instance of the newly created CRD on the cluster
```
$ kubectl get himessages.messaging.abyster.com
No resources found.
```
```
kubectl get himessage.messaging.abyster.com
No resources found.
```

Create a sample himessage instance
```
kubectl apply -f  config/samples/messaging_v1_himessage.yaml
himessage.messaging.abyster.com/himessage-sample created
```
Check the newly create instance
```
$ kubectl get himessage
NAME               AGE
himessage-sample   31s
```

## Start a local controller
```
$ make run
/Users/elvadasnonowoguia/go/bin/controller-gen object:headerFile=./hack/boilerplate.go.txt paths="./..."
go fmt ./...
go vet ./...
/Users/elvadasnonowoguia/go/bin/controller-gen "crd:trivialVersions=true" rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases
go run ./main.go
2020-01-01T16:44:48.445+0100	INFO	controller-runtime.metrics	metrics server is starting to listen	{"addr": ":8080"}
2020-01-01T16:44:48.446+0100	INFO	setup	starting manager
2020-01-01T16:44:48.446+0100	INFO	controller-runtime.manager	starting metrics server	{"path": "/metrics"}
2020-01-01T16:44:48.550+0100	INFO	controller-runtime.controller	Starting EventSource	{"controller": "himessage", "source": "kind source: /, Kind="}
2020-01-01T16:44:48.654+0100	INFO	controller-runtime.controller	Starting Controller	{"controller": "himessage"}
2020-01-01T16:44:48.757+0100	INFO	controller-runtime.controller	Starting workers	{"controller": "himessage", "worker count": 1}
2020-01-01T16:44:48.757+0100	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "himessage", "request": "test/himessage-sample"}
```

Create another himessage item an you will see the following logs from the server output
```
$ kubectl apply -f himessage1.yaml
```
Server
```
020-01-01T16:47:30.170+0100	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "himessage", "request": "test/himessage-1"}
```

What's if we want our HiMessage instances to be called with the shortname hm/hi/him ?
```
$ kubectl get him
error: the server doesn't have a resource type "him"
```



// +kubebuilder:resource:path=services,shortName=hi;him
