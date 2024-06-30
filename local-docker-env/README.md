# Local Test 환경 구성 가이드

이 문서는 `go-wemix` 리포지토리를 기반으로 로컬 테스트 환경을 구성하는 방법에 대해 설명합니다.

## Prerequisites

- `make`가 로컬 환경에 설치되어 있어야 합니다. `gwemix`를 사용하기 위해 필요합니다.

## 사용 가능한 스크립트

### local-docker.sh

- **목적**: 로컬 리포지토리 기준으로 실행되는 로컬 테스트를 위한 통합 구동 스크립트입니다.
- **사용법**: `./local-docker-env/local-docker.sh -a <account_num>`

### local-docker-git.sh

- **목적**: Git 리포지토리 기준으로 실행되는 로컬 테스트를 위한 통합 구동 스크립트입니다.
- **사용법**: `./local-docker-env/local-docker-git.sh -a <account_num> -b <branch> -r <repo>`

### key-gen.sh

- **목적**: `keystore` 폴더에 account를 생성하고, `nodekey` 폴더에 nodekey를 생성합니다.
- **사용법**: `./key-gen.sh -a <account_num>`

### config-gen.sh

- **목적**: `config.json` 파일을 생성합니다.
- **사용법**: `./config-gen.sh -a <account_num> -f <output_config_file>`

### docker-compose-gen.sh

- **목적**: 로컬 리포지토리 기준으로 실행되는 `docker-compose.yml` 파일을 생성합니다.
- **사용법**: `./docker-compose-gen.sh -a <account_num> -f <output_file_path>`

### docker-compose-gen-git.sh

- **목적**: Git 리포지토리 기준으로 실행되는 `docker-compose.yml` 파일을 생성합니다.
- **사용법**: `./docker-compose-gen-git.sh -a <account_num> -f <output_file_path> -r <repo_url> -b <branch_name>`

### init-boot.sh

- **목적**: 부트 노드에서의 governance contract 배포 및 노드 구동 스크립트입니다.

### init-node.sh

- **목적**: 일반 노드에서의 genesis block download 및 구동 스크립트입니다.

### set-nodekey.sh

- **목적**: 실제 노드에서 사용할 nodekey setting 스크립트입니다.

## Dockerfile 설명

- **Dockerfile.boot**: 로컬 리포지토리 기준으로 실행되는 부트 노드 도커 파일입니다.
- **Dockerfile.node**: 로컬 리포지토리 기준으로 실행되는 노드 도커 파일입니다.
- **Dockerfile.boot.git**: Git 리포지토리 기준으로 실행되는 부트 노드 도커 파일입니다.
- **Dockerfile.node.git**: Git 리포지토리 기준으로 실행되는 노드 도커 파일입니다.

## 실행 방법

### 실행 위치

- `go-wemix` 리포지토리를 클론한 경로에서 실행합니다.

### 로컬 리포 코드 기준 실행

```bash
./local-docker-env/local-docker.sh -a <account_num>
```

### Git 리포 코드 기준 실행

```bash
./local-docker-env/local-docker-git.sh -a <account_num> -b <branch> -r <repo>
```
