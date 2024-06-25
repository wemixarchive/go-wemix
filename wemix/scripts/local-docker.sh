#!/bin/bash

# 필수 인수 확인
if [ "$#" -ne 8 ]; then
  echo "Usage: $0 -a <account_num> -f <output_config_file> -b <branch> -r <repo>"
  exit 1
fi

# 옵션 파싱
while getopts "a:f:b:r:" opt; do
  case ${opt} in
  a)
    ACCOUNT_NUM=$OPTARG
    ;;
  f)
    OUTPUT_CONFIG_FILE=$OPTARG
    ;;
  b)
    BRANCH=$OPTARG
    ;;
  r)
    REPO=$OPTARG
    ;;
  \?)
    echo "Usage: $0 -a <account_num> -f <output_config_file> -b <branch> -r <repo>"
    exit 1
    ;;
  esac
done

# 필수 인수 확인
if [ -z "$ACCOUNT_NUM" ] || [ -z "$OUTPUT_CONFIG_FILE" ] || [ -z "$BRANCH" ] || [ -z "$REPO" ]; then
  exit 1
fi

# 초기화
echo '{"members":[], "accounts":[]}' >"$OUTPUT_CONFIG_FILE"

# 파일을 JSON 형식으로 로드
json_content=$(cat "$OUTPUT_CONFIG_FILE")

# 고정된 env 값을 포함한 JSON 객체
ENV_JSON='{
  "ballotDurationMin": 86400,
  "ballotDurationMax": 604800,
  "stakingMin": 1500000000000000000000000,
  "stakingMax": 1500000000000000000000000,
  "MaxIdleBlockInterval": 100,
  "blockCreationTime": 5000,
  "blockRewardAmount": 1000000000000000000,
  "maxPriorityFeePerGas": 100000000000,
  "rewardDistributionMethod": [4000, 1000, 2500, 2500],
  "maxBaseFee": 50000000000000,
  "blockGasLimit": 105000000,
  "baseFeeMaxChangeRate": 55,
  "gasTargetPercentage": 30
}'

json_content=$(echo "$json_content" | jq --argjson env "$ENV_JSON" \
  '.env = $env| .extraData = "chain for local test"')

mkdir keystore
mkdir nodekey

for ((i = 1; i <= ACCOUNT_NUM; i++)); do
  if [ -f "keystore/account$i" ]; then
    rm keystore/account$i
  fi
  yes "test" | head -n 2 | gwemix wemix new-account --out keystore/account$i
  
  if [ -f "nodekey/nodekey$i" ]; then
    rm nodekey/nodekey$i
  fi
  gwemix wemix new-nodekey --out nodekey/nodekey$i

  ids=$(gwemix wemix nodeid nodekey/nodekey$i)
  idv5=$(echo "$ids" | awk '/idv5:/ {print $2}')
  idv5="0x$idv5"
  ip="172.16.237.$((i+10))"

  # account1 파일에서 address 필드 값 추출
  address=$(jq -r '.address' "keystore/account$i")
  address="0x$address"
  name="localtest$i"

  json_content=$(echo "$json_content" | jq --arg addr "$address" --arg name "$name" --arg id "$idv5" --arg ip "$ip" --argjson index $((i - 1)) \
    '.members += [{"addr": $addr, "stake": 1500000000000000000000000, "name": $name, "id": $id, "ip": $ip, "port": 8589, "bootnode": false}] 
    | .accounts += [{"addr": $addr, "balance": 2000000000000000000000000}]')

  if [ $i -eq 1 ]; then
    json_content=$(echo "$json_content" | jq --arg addr "$address" \
      '.staker = $addr | .ecosystem = $addr | .maintanence = $addr | .feecollector = $addr | .members[0].bootnode = true')
  fi
done

# 수정된 내용을 파일에 저장
echo "$json_content" >"$OUTPUT_CONFIG_FILE"
echo "Updated config saved to $OUTPUT_CONFIG_FILE"

# BRANCH와 REPO 정보를 입력으로 받아 gen-docker-compose.sh 실행
./gen-docker-compose.sh -a "$ACCOUNT_NUM" -b "$BRANCH" -r "$REPO"

# docker-compose.yml 파일을 이용해 docker-compose build 및 up 실행
docker compose -f docker-compose.yml build --no-cache
docker compose -f docker-compose.yml up -d