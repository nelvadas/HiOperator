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
Initialize the kubebuilder project with the domain `abyster.com` for eg
```
$ kubebuilder init --domain abyster.com
```


## Create a HiMessage resource and API
```
$ kubebuilder create api --group messaging --version v1 --kind HiMessage --resource --controller
```
The `--resource ` option generates the CRD without prompting
The `--controller `" options generates a controller without prompting

## Deploy the CRD in k8s cluster

While testing this sample, we are running a k8S 1.14 Cluster
```
$ kubectl get nodes
NAME          STATUS   ROLES    AGE    VERSION
ucpleader     Ready    master   132d   v1.14.3-docker-2
ucpworker-0   Ready    <none>   132d   v1.14.3-docker-2
ucpworker-1   Ready    <none>   132d   v1.14.3-docker-2
```

Kubebuilder provides a makefile with a target for most of our operations

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


## Working with Himessage instances

Check if we have any instance of the newly created CRD on the cluster
```
$ kubectl get himessages.messaging.abyster.com
No resources found.
```
The same request can be runned with the defualt plural form `himessages`
```
kubectl get himessages.messaging.abyster.com
No resources found.
```

Create a sample Himessage object
```
kubectl apply -f  config/samples/messaging_v1_himessage.yaml
himessage.messaging.abyster.com/himessage-sample created
```
Check the newly created object
```
$ kubectl get himessage
NAME               AGE
himessage-sample   31s
```

## Start a local Controller
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

## CRD Customization

What's if we want our HiMessage instances to be called with the shortname `hm/hi/him` ?
```
$ kubectl get him
error: the server doesn't have a resource type "him"
```

Customize the HiMessage structure with the kubuilder:resource Marker
```
 +kubebuilder:resource:categories=messaging,path=himessages,singular=himessage,shortName=hi;him;himesg
 HiMessage is the Schema for the himessages API
type HiMessage struct {
```

Regenerate the resources using
```
$ make manifests
```
Check the existing CRD using the defined shortName

```
$ kubectl get hi
NAME               AGE
himessage-sample   25m
```
## CRD Validation

Include the following marker to allow only 10 chars max in the Message
```
 \\\\+kubebuilder:validation:MaxLength:=10 Message string `json:"message,omitempty"`
```
Create a YAML resource that violates this constraint
```
apiVersion: messaging.abyster.com/v1
kind: HiMessage
metadata:
  name: himessage-invalid-size
spec:
  message: "Hello World Hello World HelloWorld"
  image: alpine


$ kubectl apply -f msg_invalidsize.yaml
The HiMessage "himessage-invalid-size" is invalid: []: Invalid value: map[string]interface {}{"apiVersion":"messaging.abyster.com/v1", "kind":"HiMessage", "metadata":map[string]interface {}{"annotations":map[string]interface {}{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"messaging.abyster.com/v1\",\"kind\":\"HiMessage\",\"metadata\":{\"annotations\":{},\"name\":\"himessage-invalid-size\",\"namespace\":\"test\"},\"spec\":{\"image\":\"alpine\",\"message\":\"Hello World Hello World HelloWorld\"}}\n"}, "creationTimestamp":"2020-01-01T21:21:56Z", "generation":1, "name":"himessage-invalid-size", "namespace":"test", "uid":"bcff318d-2cdc-11ea-a8e7-0242ac11000a"}, "spec":map[string]interface {}{"image":"alpine", "message":"Hello World Hello World HelloWorld"}}: validation failure list:
spec.message in body should be at most 10 chars long
```
