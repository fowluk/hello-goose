# Hello Goose

This section contains Kubernetes deployments for Hello Goose.

To work in Kubernetes we have to build a docker image. We offer different examples for how to do that.

* [buildpacks.io](buildpacks.io/): An example using [Cloud Native Buildpacks](https://buildpacks.io/)
* TBD

Each one contains two scripts:
* 1-build.sh:  Script to build the image and push it to the official Hello-Goose dockerhub.
* 2-push.sh:  A script that applies the example deployment and service files in k8s.

Additional documentation is contained in the README for each version. For simplicity, the k8s versions just push the python application.

