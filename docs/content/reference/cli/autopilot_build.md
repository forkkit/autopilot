---
title: "autopilot build"
weight: 5
---
## autopilot build

Compiles code and builds artifacts

### Synopsis

The operator-sdk build command compiles the Operator code into an executable binary
and generates the Dockerfile manifest.

<image> is the container image to be built, e.g. "quay.io/example/operator:v0.0.1".
This image will be automatically set in the deployment manifests.

After build completes, the image would be built locally in docker. Then it needs to
be pushed to remote registry.
For example:
	$ operator-sdk build quay.io/example/operator:v0.0.1
	$ docker push quay.io/example/operator:v0.0.1


```
autopilot build <image> [flags]
```

### Options

```
      --go-build-args string      Extra Go build arguments as one string such as "-ldflags -X=main.xyz=abc"
  -h, --help                      help for build
      --image-build-args string   Extra image build arguments as one string such as "--build-arg https_proxy=$https_proxy"
```

### Options inherited from parent commands

```
  -v, --verbose   Enable verbose logging
```

### SEE ALSO

* [autopilot](../autopilot)	 - An SDK for building Service Mesh Operators with ease

