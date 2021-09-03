#!/bin/bash
set -e
echo "正在准备编译标记版本"
gitTime=$(date +00%y%m%d%H%M%S)
gitCID=`git rev-parse HEAD`

echo "正在生成静态文件缓存"
go mod tidy
go generate

go build -ldflags "-X main.build=${gitTime}.${gitCID}" -o "bin.${gitTime}"

"./bin.${gitTime}"