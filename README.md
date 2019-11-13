# AutoPilot - The Service Mesh Automator

**AutoPilot** is an SDK and toolkit for developing and deploying [service mesh operators](). 

**AutoPilot** generates scaffolding, builds, and deploys Operators which run against a local or remote Kubernetes cluster installed with a Service Mesh. 

**AutoPilot** generated code and libraries provide an easy way to automate configuration and monitoring of a service mesh (and other Kubernetes/infra resources) via the [Kubernetes Operator](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) pattern.

# How does it work?

Developers define an `autopilot.yaml` and `autopilot-operator.yaml` which specify the skeleton and configuration of an *AutoPilot Operator*.

AutoPilot makes use of these files to (re-)generate the project skeleton, build, deploy, and manage the lifecycle of the operator via the `ap` CLI.

The `ap` CLI is designed to manage the full lifecycle of the operator, but this can be done with standard tooling (`go`, `docker`, `kubectl`, `helm`, etc.). 

# How is it different from SDKs like Operator Framework and Kubebuilder?

The [Operator Framework](https://github.com/operator-framework) and [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder) are open-ended SDKs that take a far less opinionated approach to building Kubernetes software.

**AutoPilot** provides a more opinionated control loop via a generated *scheduler* that implements the [Controller-Runtime Reconciler interface](https://github.com/kubernetes-sigs/controller-runtime/blob/master/pkg/reconcile/reconcile.go#L80), for which users write stateless Work functions for various states of their top-level CRD. State information is stored
 on the *status* of the CRD, promoting a stateless design for AutoPilot operators.
 
**AutoPilot** additionally provides primitives, generated code, and helper functions for interacting with a variery of service meshes. While AutoPilot can be used to build operators that do not configure or monitor a mesh, much of *AutoPilot*'s design has been oriented to facilitate easy integration with popular service meshes.

Finally, **AutoPilot** favors simplicity over flexibility, though it is the intention of the project to support the vast majority of DevOps workflows built on top of Kubernetes+Service mesh.

# Hello World



# Roadmap
- Support for managing multiple (remote) clusters.

## scrap

AutoPilot provides an opinionated structure 
for executing an operator's 
workflow. Read more about the 
[AutoPilot Architecture]() to learn about 
how AutoPilot Operators schedule and execute work.

Code generation can also be invoked from Go code using the `codegen` package. 

AutoPilot is composed of 3 components:
- `cli`
- `codegen` package
- `pkg` libraries



# todo

# cleanup
- example
- docs 
- refactor metrics. make a factory for it
- code generated headers in the right places
- improve docs generation template
- bake templates into cli
- add many more aliases/parameter types
- clean up CLI messages
- generate CLI docs
- idempotent generation of rbac yaml (rule ordering not idempotent)

## test
- e2e metrics test with istio

## features
- git ops
- define custom metrics queries in autopilot.yaml
- validate method for project config
    - check operatorName is kube compliant
    - apiVerson, kind, phases are correct
    - customParameters
    - final phase with i/o
- add user config to configmap with config settings
- curl script to download
- builders
- ap undeploy
- label everything for easy deletion/listing
- expose garbage collection func to workers
    - rollback the phase when something ensure fails? (option in config)
- multiple crds

## punt
- schema generation
- interactive cli
- automatic metrics for worker syncs
- automatic traces for worker syncs







# done 
* works across namespaces..


# code guide:

explain where all the existing things are - the templated queries and clients

folder -> what does it do