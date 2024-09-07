#!/bin/bash

# 检查是否提供了两个参数
if [ "$#" -ne 2 ]; then
    echo "用法: $0 <参数1> <参数2>"
    exit 1
fi

# 获取两个参数
PARAM1="$1"
PARAM2="$2"

# 可执行文件路径
EXECUTABLE="./bin/generate_code"

# 检查可执行文件是否存在
if [ -f "$EXECUTABLE" ]; then
    # 运行可执行文件并传递参数
    "$EXECUTABLE" "$PARAM1" "$PARAM2"
else
    echo "错误：找不到可执行文件 $EXECUTABLE，请先运行 build.sh 进行编译"
    exit 1
fi
