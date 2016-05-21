[![Build Status](https://travis-ci.org/open-devops/engine.svg?branch=master)](https://travis-ci.org/open-devops/engine)
[![codecov.io](https://codecov.io/github/open-devops/engine/coverage.svg?branch=master)](https://codecov.io/github/open-devops/engine?branch=master)

A lightweight dev op spec runtime.

# Official SDK's

[sdk-for-golang](https://github.com/open-devops/sdk-for-golang)

# Supported Use Cases
- add op
- add sub op
- get event stream
- get liveness
- kill op run
- list ops
- run op
- set description of op


# Prerequisites
- [docker](https://github.com/docker/docker) >= 1.10

Note: if using Windows or OSX, you need to update your docker-machine to use NFS instead of vboxfs 
(or suffer painfully slow performance). One recommended way to achieve this is via 
[docker-machine-nfs](https://github.com/adlogix/docker-machine-nfs). 
Your mileage may vary.

# Example Usage

### 1) Start dockerized engine
```SHELL
docker run -it --rm -v /var/run/docker.sock:/var/run/docker.sock -v /Users:/Users -p 42224:42224 devopspec/engine
```
explanation:

- `-it` interactive/tty
- `--rm` remove on exit
- `-v /var/run/docker.sock:/var/run/docker.sock` bind mount host docker socket
- `-v /Users:/Users` bind mount host `/Users` dir
- `-p 42224:42224` expose container port `42224` via docker-machine port `42224`
- `devopspec/engine` use latest [devopspec/engine](https://hub.docker.com/r/devopspec/engine/) image

### 2) Explore the engine API via Swagger UI

open your browser and navigate to the url returned from `echo $(docker-machine ip):42224`

# Releases
All releases will be [tagged](https://github.com/open-devops/engine/tags) and made available on the [releases](https://github.com/open-devops/engine/releases) page with release notes.

# Versioning
This project adheres to the [Semantic Versioning](http://semver.org/) specification

*Be advised: this project is currently at Major version zero. Per the semantic versioning spec: "Major version zero (0.y.z) is for initial development. Anything may change at any time. The public API should not be considered stable."*

# Contributing

refer to [CONTRIBUTING.md](CONTRIBUTING.md)
