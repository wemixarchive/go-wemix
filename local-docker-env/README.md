# Local Test Environment Setup Guide

This document explains how to set up a local testing environment based on the `go-wemix` repository.

## Prerequisites

- `make` must be installed in the local environment. It is required to use `gwemix`.

## Available Scripts

### local-docker.sh

- **Purpose**: Integrated running script for local testing based on the local repository.
- **Usage**: `./local-docker-env/local-docker.sh -a <account_num> -v <ubuntu_version>(optional) --keep-nodekey(optional)`

### local-docker-git.sh

- **Purpose**: Integrated running script for local testing based on the Git repository.
- **Usage**: `./local-docker-env/local-docker-git.sh -a <account_num> -b <branch> -r <repo> -v <ubuntu_version>(optional --keep-nodekey(optional)`

### gen-nodekey.sh

- **Purpose**: Script for generating a nodekey for the node.
- **Usage**: `./local-docker-env/gen-nodekey.sh -a <nodekey_num>`

### gen-account.sh

- **Purpose**: Script for generating a nodekey for the node.
- **Usage**: `./local-docker-env/gen-account.sh -a <account_num>`

### init-boot.sh

- **Purpose**: Script for deploying the governance contract and running the node on the boot node.

### init-node.sh

- **Purpose**: Script for downloading the genesis block and running the node on a regular node.

### set-nodekey.sh

- **Purpose**: Script for setting the nodekey to be used on the actual node.

## Dockerfile Description

- **Dockerfile.boot**: Docker file for running the boot node based on the local repository.
- **Dockerfile.node**: Docker file for running the node based on the local repository.
- **Dockerfile.boot.git**: Docker file for running the boot node based on the Git repository.
- **Dockerfile.node.git**: Docker file for running the node based on the Git repository.

## Execution Method

### Execution Location

- Run from the cloned path of the `go-wemix` repository.

### Execution based on the local repository code

```bash
./local-docker-env/local-docker.sh -a <account_num> -v <ubuntu_version> (optional) --keep-nodekey (optional)
```

### Execution based on the Git repository code

```bash
./local-docker-env/local-docker-git.sh -a <account_num> -b <branch> -r <repo> -v <ubuntu_version> (optional) --keep-nodekey (optional)
```
