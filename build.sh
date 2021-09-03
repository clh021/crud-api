#!/bin/bash
set -ex
echo "正在准备编译标记版本"
gitTime=$(git log -1 --format=%at | xargs -I{} date -d @{} +%Y%m%d_%H%M%S)
gitCID=`git rev-parse HEAD`

echo "正在生成静态文件缓存"
go mod tidy
go generate




go build -ldflags "-X main.build=${gitTime}.${gitCID}" -o "bin.${gitTime}"
