#!/bin/bash

# 執行專案的目錄
WORK_PATH=$(dirname $(readlink -f $0))
# 專案名稱(取當前資料夾路徑最後一個資料夾名稱)
PROJECT_NAME=${WORK_PATH##*/}
# 當前環境
ENV="local"

# 第一個參數為 LineBot Channel Secret
if [ -z "$1" ]
then
  echo "CHANNEL_SECRET is required arguments"
  exit
fi

# 第二個參數為 LineBot Channel Token
if [ -z "$2" ]
then
  echo "CHANNEL_TOKEN is required arguments"
  exit
fi

# 第三個參數為 ChatGPT Key
if [ -z "$3" ]
then
  echo "GPT_KEY is required arguments"
  exit
fi

# 存入.env
echo "ENV=$ENV">.env
echo "PROJECT_NAME=$PROJECT_NAME">>.env
echo "CHANNEL_SECRET=$1">>.env
echo "CHANNEL_TOKEN=$2">>.env
echo "GPT_KEY=$3">>.env

# 啟動容器服務
docker-compose up -d