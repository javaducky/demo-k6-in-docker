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

## Lab 2: Running your image
This test our newly built image and create a container to execute our `test-scripts/simple.js`.
```shell
docker run -i k6-extended:latest run - < test-scripts/simple.js
```
Congratulations, you've executed a k6 test using Docker!
Our command had Docker `run` a new container based upon the `k6-extended:latest` image, providing `run -` to let k6 know to read from `stdin`, to which we pipe the contents of `test-scripts/simple.js` using the `<` redirection operator.

But wait... did you happen to notice the output from your test run?
```shell
          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: -
     output: -
```
The script name is "-" although we ran `test-scripts/simple.js`? 
This is because we had k6 read our script content from `stdin` which our operating system did for use using the contents of `test-scripts/simple.js`.

Now, let's talk a little about housekeeping with Docker as well. 
Each execution of the above command creates a new _container_ based upon the named _image_.
The container has state from the test execution, e.g. logs and such.

Using `docker ps`, we can view the currently running containers. 
Our example runs pretty quick, so chances are they're no longer running; you won't see any containers based on the `k6-extended:latest` image.
Instead, let's execute `docker ps --all` which will show all containers whether currently running or not.
Depending upon how many times you ran your test, you'll see an entry for each execution.
Guess what? Each of these is unnecessarily taking up space on your disk drive!
```shell
CONTAINER ID   IMAGE                COMMAND      CREATED          STATUS                      PORTS     NAMES
1d42a900e312   k6-extended:latest   "k6 run -"   7 minutes ago    Exited (0) 7 minutes ago              mystifying_mayer
40d916d89e2f   k6-extended:latest   "k6 run -"   16 minutes ago   Exited (0) 16 minutes ago             sharp_nightingale
```
For each of these containers, let's clear them up by running `docker rm <container name>`, e.g. `docker rm mystifying_mayer`.

Let's improve our execution by having Docker clean up after itself and address the lack of a script name in our output.
```shell
docker run -v $PWD:/scripts -it --rm k6-extended:latest run /scripts/test-scripts/simple.js
```
Sweet! Now you should have noticed the script name shown in your output and that a quick check of `docker ps --all` shows no loitering containers from our test!

## Lab 3: k6-in-docker script

TODO - Overview of the wrapper functionality
* Overview of the `k6-in-docker.sh` script.
* Call out the `--env-file` and `--network` options
* Overview of the `script_vars.env`
* Describe how the values can be overridden without changing the script itself
* DEMO the script with the same scenario

## Lab 4: Docker Compose setup in the project

TODO - Overview of project layout for docker-compose files and dependencies.
* Overview of `grafana-prom.yml` config
* Start the containers using `docker-compose -f docker-compose/grafana-prom.yml up`
* Access the [Grafana instance](http://localhost:3000/)
* Call out the available dashboards
* Overview the `dependencies` directory and the grafana dashboards there
* Show output for `docker ps`

## Lab 5: Targeting environments with env variables

TODO - Run our test script against the running Prom container then Grafana Cloud
* Update `script_vars.env` for `K6_OUT` to xk6 remote write
* Execute our script against the local environment
* ??? Call out the `--network` override and `docker network ls` to list networks
* ??? Can show what happens when the network is not matching or `--network` is omitted
* Discuss other options in the example file
* Run against Grafana Cloud `ENV_FILE=.private/.env-grafanacloud ./k6-in-docker.sh test-scripts/simple.js`

