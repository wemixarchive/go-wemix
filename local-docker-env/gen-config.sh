#!/bin/bash

# Check required arguments
if [ "$#" -ne 4 ]; then
  echo "Usage: $0 -a <node_num> -f <output_config_file>"
  exit 1
fi

# Option parsing
while getopts "a:f:" opt; do
  case ${opt} in
  a)
    NODE_NUM=$OPTARG
    ;;
  f)
    OUTPUT_CONFIG_FILE=$OPTARG
    ;;
  \?)
    echo "Usage: $0 -a <node_num> -f <output_config_file>"
    exit 1
    ;;
  esac
done

# Check required arguments
if [ -z "$NODE_NUM" ] || [ -z "$OUTPUT_CONFIG_FILE" ]; then
  exit 1
fi

# Initialize
echo '{"members":[], "accounts":[]}' >"$OUTPUT_CONFIG_FILE"

# Load file as JSON
json_content=$(cat "$OUTPUT_CONFIG_FILE")

# JSON object with fixed env values
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

# Generate JSON config file
for ((i = 1; i <= NODE_NUM; i++)); do
  ids=$(gwemix wemix nodeid local-docker-env/nodekey/nodekey$i) || { echo "Failed to get node ID"; exit 1; }
  idv5=$(echo "$ids" | awk '/idv5:/ {print $2}')
  idv5="0x$idv5"
  ip="172.16.237.$((i+10))"
  
  address=$(jq -r '.address' "local-docker-env/keystore/account$i") || { echo "Failed to get address"; exit 1; }
  address="0x$address"
  name="localtest$i"

  json_content=$(echo "$json_content" | jq --arg addr "$address" --arg name "$name" --arg id "$idv5" --arg ip "$ip" --argjson index $((i - 1)) \
    '.members += [{"addr": $addr, "stake": 1500000000000000000000000, "name": $name, "id": $id, "ip": $ip, "port": 8589, "bootnode": false}] 
    | .accounts += [{"addr": $addr, "balance": 2000000000000000000000000}]') || { echo "Failed to update JSON content"; exit 1; }

  if [ $i -eq 1 ]; then
    json_content=$(echo "$json_content" | jq --arg addr "$address" \
      '.staker = $addr | .ecosystem = $addr | .maintanence = $addr | .feecollector = $addr | .members[0].bootnode = true') || { echo "Failed to update JSON content"; exit 1; }
  fi
done

# Save modified content to file
echo "$json_content" >"$OUTPUT_CONFIG_FILE"
echo "Updated config saved to $OUTPUT_CONFIG_FILE"