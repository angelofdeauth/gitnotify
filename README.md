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

### Requirements

The build scripts in `hack/` test for the existence of required build tools on the system.
If the tools are not found, `podman` is used to run a container with the tool.
Thus, if the system does not have any of the required tools, the only requirement is `podman`.
Docker can be used instead by aliasing podman to docker.
```sh
alias podman="docker"
```

### Instructions

The `Makefile`'s default action is to clean and run a build.
Run `make` from the reporitory root directory, the executable will be output to `reporoot/bin/${VERSION}/${OS_ARCH}/`.

To run the scripts individually, you must export all of the variables in the `Makefile`.

### New Commits

The existing build tools can aid in making new commits.
If there are tracked files with cached changes ( added via `git add 'filename'`), only the cached changes will be automatically used.
If there are tracked files with uncached changes, all changes will be added to the commit (including new files).
If there are only new files, all changes will be added to the commit.
`make commit` from the repository root directory will perform all necessary automation and prompt the user for a commit message.

### Tagging a new version

For those who choose to customize this tool and track it in their own repositories, tagging a new version using the existing build tools is simple.
Simply execute the following:
```sh
VERSION=DESIRED_VERSION make release
```

### TODO

> TODO: set up pipelines for creating custom images which only contain the tools required per build step.

> TODO: set up pipelines for updating the tools in the custom images

> TODO: set up pipeline for updating the referenced version of images in the `hack/` scripts.


## Usage

`xnotify` is both a client for configuration and a daemon for monitoring filesystem events.


## Existing Integrations


## Config Reference


## Template Reference
