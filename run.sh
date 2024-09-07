#!/bin/bash

# 定义有效参数选项
VALID_OPTIONS="msg db ui"

# 检查是否提供了两个参数
if [ "$#" -ne 2 ]; then
    echo "用法: $0 <参数1> <参数2>"
    echo "参数1 必须是以下之一: $VALID_OPTIONS"
    echo "参数2 是待生成代码的核心命名"
    exit 1
fi

# 获取两个参数
PARAM1="$1"
PARAM2="$2"

# 验证参数1是否有效
if ! echo "$VALID_OPTIONS" | grep -w "$PARAM1" > /dev/null; then
    echo "错误: 参数1 '$PARAM1' 无效。有效选项是: $VALID_OPTIONS"
    exit 1
fi

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
