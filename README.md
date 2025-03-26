# k8s-webhook-learning


## Conversion Webhook Notes
A basic example to start with can be found in main. It contains a boilerplate kubebuilder project with a single type called greeter. 
It will simply produce a log every 10 seconds. In the log it will print the people specified in the spec as a comma separated list. The status is updated to keep the number of times the reconcilation run and the timestamp of the last run. 

It serves as a base to build webhooks on top of that. Those will be kept in separate branches.

### Notes Conversion webhook
**Prequisites**

Create Cluster
```
# create new local cluster
kind create cluster
```

#### Worflow
Following steps are required. For each of them there is a commit in the related branch. 

**add second version**
```
kubebuilder create api --group friendly --version v1beta1 --kind Greeter
```

**set storage version annotation (to v1beta1) -> see commit**
```
// +kubebuilder:storageversion
```



Our Guidelines to be defined : 
- how to choose storage version?

