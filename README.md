# demo-k6-in-docker

Project for demonstrating the how to execute tests and customize [k6](https://k6.io/) using [Docker](https://www.docker.com/); aka _"Look Ma, No Go!"_

This tutorial goes through the steps to show how you can create a customized k6 image using [xk6](https://github.com/grafana/xk6) with any number of available extensions without requiring a Go development environment. We also provide examples of running your customized k6 image against various targets.

This project expands on ["Build and run k6 with xk6 extensions using Docker"](https://gist.github.com/javaducky/6954ae9f67ec0b0420bfc21f9e8017b3).

## Prerequisites
* [Docker](https://docs.docker.com/get-docker/) - For building our custom k6 images and running the examples

## :raising_hand: How do I use this?
Clone the repo, then go through the lab in the order they're provided---some labs may build upon an earlier example.


## Lab 1: Build your first image
All Docker images we create will be built using the `Dockerfile` in the project root.

This is a _multi-stage_ Docker file which simply means we will use an initial _stage_ installing all of our build tools, i.e. Golang support within a Linux container.
This is where we will also install `xk6` then build our custom k6 binary with our desired extensions.
For this initial lab, we're creating an image that will have the `xk6-output-prometheus-remote` compiled in.

> :point_right: With v0.42.0 of k6, the extension is being provided to the k6 binary by default.

The second stage provides a _"clean"_ Linux container into which we add our new binary from the earlier stage.
This really helps on providing a slim image as the build tools can take up considerable space and are only necessary for the build itself.

Let's do this! Make sure your terminal is in the project root directory, then run the following command:
```shell
docker build -t k6-extended:latest .
```
This command tells Docker to build an image named `k6-extended` and _tagged_ as `latest` using the Dockerfile in the current directory.
In practice the `tag`-portion is _typically_ a semantic version, e.g. `grafana/k6:0.41.0`.
This may take a couple (or few) minutes initially, but as your local system caches more of the image layers, timings should improve.
> :point_right: For more of the gory details, check the [command-line reference](https://docs.docker.com/engine/reference/commandline/build/) for `docker build`.

Running the following lists all Docker images available on your system:
```shell
docker images
```
Keep in mind, these images are the basis for containers. Consider them to be the _template_ from which a running container is created.
