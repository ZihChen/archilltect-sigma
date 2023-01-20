#!/bin/bash

# 執行專案的目錄
WORK_PATH=$(dirname $(readlink -f $0))
# 專案名稱(取當前資料夾路徑最後一個資料夾名稱)
PROJECT_NAME=${WORK_PATH##*/}
# 當前環境
ENV="local"

# 存入.env
echo "ENV=$ENV">.env
echo "PROJECT_NAME=$PROJECT_NAME">>.env

# 啟動容器服務
docker-compose up -d