#!/bin/bash

# 设置编译输出目录
OUTPUT_DIR="./bin"

# 可执行文件名
EXECUTABLE="generate_code"

# 如果 bin 目录不存在，则创建
if [ ! -d "$OUTPUT_DIR" ]; then
  mkdir -p "$OUTPUT_DIR"
  echo "创建目录 $OUTPUT_DIR"
fi

# 切换到 source 目录，因为 go.mod 在这里
cd source

# 编译 Go 源码并输出到 ../bin 目录
go build -o "../$OUTPUT_DIR/$EXECUTABLE" ./main.go

# 检查编译是否成功
if [ $? -eq 0 ]; then
  echo "编译成功，可执行文件已生成: $OUTPUT_DIR/$EXECUTABLE"
else
  echo "编译失败"
  exit 1
fi
