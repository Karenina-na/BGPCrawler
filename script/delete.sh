#!/bin/bash

# 获取输入的参数
target_date=$1
dir1=$2
dir2=$3

# 在dir1中找出文件名中日期小于target_date的文件
for file in $dir1/*; do
  filename=$(basename $file)
  date_in_file=$(echo $filename | cut -d. -f2)
  if [[ $date_in_file < $target_date ]]; then
    rm "$file"
  fi
done

# 在dir2中找出文件名中日期小于target_date的文件
for file in $dir2/*; do
  filename=$(basename $file)
  date_in_file=$(echo $filename | cut -d. -f2)
  if [[ $date_in_file < $target_date ]]; then
    rm "$file"
  fi  
done