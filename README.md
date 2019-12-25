# kenv

A [pyenv](https://github.com/pyenv/pyenv)-like CLI tool for managing and using multiple [kubectl](https://github.com/kubernetes/kubectl) versions.

## Installation

### Prerequisites

- UNIX-like operating system (e.g. macOS, Ubuntu)
- bash 

kenv may work with other similar shells or setups, though they are not (yet) supported.

### Manual

```shell script
{
    go get github.com/jameshiew/kenv
    kenv init  # initializes the `.kube/kenv` directory, where shims and kubectl versions will be stored
    eval "$(kenv init -)"  # activates kenv for this shell session
    echo 'command -v kenv && eval "$(kenv init -)"' >> ~/.bashrc  # (or ~/.bash_profile) - load kenv into your shell at start up
}
```

## Quickstart

### Install a version of kubectl

```shell script
kenv install 1.17.0
```

### Set the global (i.e. default)) kubectl version

```shell script
kenv global 1.17.0
```

If no global version is set with `kenv`, the system's kubectl will be used by default.

### Set a local kubectl version

```shell script
kenv local 1.17.0 
```

Creates a `.kenv.yml` file in the local directory, that will record the `kubectl` version to use.

### Set a kubectl version for a shell session

```shell script
kenv shell 1.17.0
```

Sets the `$KENV_VERSION` environment variable.

### Everything else

Run `kenv help` to see more help strings.