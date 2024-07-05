# Local Test Environment Setup Guide

This document explains how to set up a local testing environment based on the `go-wemix` repository.

## Prerequisites

- `make` must be installed in the local environment. It is required to use `gwemix`.

## Available Scripts

### local-docker.sh

- **Purpose**: Integrated running script for local testing based on the local repository.
- **Usage**: `./local-docker-env/local-docker.sh -a <account_num> -v <ubuntu_version>(optional)`

### local-docker-git.sh

- **Purpose**: Integrated running script for local testing based on the Git repository.
- **Usage**: `./local-docker-env/local-docker-git.sh -a <account_num> -b <branch> -r <repo> -v <ubuntu_version>(optional)`

### key-gen.sh

- **Purpose**: Generates an account in the `keystore` folder and creates a nodekey in the `nodekey` folder.
- **Usage**: `./key-gen.sh -a <account_num>`

### config-gen.sh

- **Purpose**: Generates the `config.json` file.
- **Usage**: `./config-gen.sh -a <account_num> -f <output_config_file>`

### docker-compose-gen.sh

- **Purpose**: Generates the `docker-compose.yml` file based on the local repository.
- **Usage**: `./docker-compose-gen.sh -a <account_num> -f <output_file_path> -v <ubuntu_version>(optional)`

### docker-compose-gen-git.sh

- **Purpose**: Generates the `docker-compose.yml` file based on the Git repository.
- **Usage**: `./docker-compose-gen-git.sh -a <account_num> -f <output_file_path> -r <repo_url> -b <branch_name> -v <ubuntu_version>(optional)`

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
./local-docker-env/local-docker.sh -a <account_num> -v <ubuntu_version>(optional)
```

### Execution based on the Git repository code

```bash
./local-docker-env/local-docker-git.sh -a <account_num> -b <branch> -r <repo> -v <ubuntu_version>(optional)
```
