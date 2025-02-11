#!/bin/bash

# Usage: ./gen-docker-compose.sh -a <account_num> -f <output_config_file> -r <repo_url> -b <branch_name>
# -a: 계정 수
# -f: 출력 파일
# -r: 레포지토리 URL
# -b: 브랜치 이름
# -a 옵션은 필수이며, -f 옵션은 선택적
# -f 옵션을 지정하지 않으면, 현재 디렉토리에 docker-compose.yml 파일을 생성
# -f 옵션을 지정하면, 해당 경로에 docker-compose.yml 파일을 생성

# Parse command line arguments
while getopts "a:f:r:b:v:" opt; do
    case $opt in
        a)
            account_num=$OPTARG
            ;;
        f)
            output_file=$OPTARG
            ;;
        r)
            repo_url=$OPTARG
            ;;
        b)
            branch_name=$OPTARG
            ;;
        v)
            version=$OPTARG
            ;;
        \?)
            echo "Invalid option: -$OPTARG" >&2
            exit 1
            ;;
        :)
            echo "Option -$OPTARG requires an argument." >&2
            exit 1
            ;;
    esac
done

# Check if account_num is provided
if [ -z "$account_num" ]; then
    echo "Error: -a <account_num> is required."
    exit 1
fi

# Check if repo_url is provided
if [ -z "$repo_url" ]; then
    echo "Error: -r <repo_url> is required."
    exit 1
fi

# Check if branch_name is provided
if [ -z "$branch_name" ]; then
    echo "Error: -b <branch_name> is required."
    exit 1
fi

# Set default output file if not provided
if [ -z "$output_file" ]; then
    output_file="docker-compose.yml"
fi

# Set default version if not provided
if [ -z "$version" ]; then
    version="latest"
fi

# Generate docker-compose.yml content
cat <<EOF > "$output_file"
services:
EOF

for ((i=1; i<=account_num; i++)); do
    if [ $i -eq 1 ]; then
        cat <<EOF >> "$output_file"
  wemix-boot:
    build:
      context: .
      dockerfile: Dockerfile.boot.git
      args:
        REPO: $repo_url
        BRANCH: $branch_name
        NODE_NUM: $i
        UBUNTU_VERSION: $version
    image: wemix/node-boot:latest
    hostname: wemix-boot
    networks:
      wemix-dev-bridge:
        ipv4_address: 172.16.237.11
    restart: unless-stopped
    tty: true
    ports:
      - 8588:8588
      - 8589:8589
      - 8598:8598
    container_name: wemix-boot
EOF
    else
        cat <<EOF >> "$output_file"
  wemix-node$((i-1)):
    build:
      context: .
      dockerfile: Dockerfile.node.git
      args:
        REPO: $repo_url
        BRANCH: $branch_name
        NODE_NUM: $i
        UBUNTU_VERSION: $version
    image: wemix/node$((i-1)):latest
    hostname: wemix-node$((i-1))
    networks:
      wemix-dev-bridge:
        ipv4_address: 172.16.237.$((i+10))
    restart: unless-stopped
    tty: true
    depends_on:
      wemix-boot:
        condition: service_started
    container_name: wemix-node$((i-1))
EOF
    fi
done

cat <<EOF >> "$output_file"
networks:
    wemix-dev-bridge:
        name: wemix-dev-bridge
        driver: bridge
        ipam:
            driver: default
            config:
                - subnet: 172.16.237.0/24
EOF

# Echo the message that docker-compose.yml is created
echo "docker-compose.yml has been created."