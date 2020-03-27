xnotify
=======

xnotify is a filesystem event based developer workflow automation daemon.
It keeps track of configured directories and runs templated scripts based on event type.
For example, a `directory` `create` event can trigger a script to configure that new directory for event monitoring.
Similarly, a `file` `update` event can trigger a script to write a git commit for that update.

## Index


## Supported Platforms

### FreeBSD

### GNU/Linux

### Mac OS

### Windows


## Installation

### FreeBSD

### GNU/Linux

### Mac OS

### Windows


## Build From Source

The `Makefile`'s default action is to clean and run a build.

### Requirements

The build scripts in `hack/` test for the existence of required build tools on the system.
If the tools are not found, `podman` is used to run a container with the tool.
Thus, if the system does not have any of the required tools, the only requirement is `podman`.
Docker can be used instead by aliasing podman to docker.
```sh
alias podman="docker"
```

> TODO: set up pipelines for creating custom images which only contain the tools required per build step.
> TODO: set up pipelines for updating the tools in the custom images
> TODO: set up pipeline for updating the referenced version of images in the `hack/` scripts.


## Usage

`xnotify` is both a client for configuration and a daemon for monitoring filesystem events.


## Existing Integrations


## Config Reference


## Template Reference
