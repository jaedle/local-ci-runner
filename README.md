# Local CI Runner

Run your CI pipeline locally

## Motivation

- Submit a workflow run to a local ci
- Ensure that no state is on your machine
- Run the same verification locally as on the remote CI

## Installation

As this is work in progress, the easiest way to install is to install from source.

```sh
git clone https://github.com/jaedle/local-ci-runner.git
cd local-ci-runner
task install
```

## Usage

### Bootstrap your machine

```sh
lcr bootstrap
```


### Submit a new CI run

```sh
lcr start
```

### List all CI runs

```sh
lcr list
```

### Clean previous CI runs

```sh
lcr clean
```
