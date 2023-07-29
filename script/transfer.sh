#!/bin/bash

if [ $# -ne 2 ]; then
  echo "Usage: $0 <input_dir> <output_dir>"
  exit 1
fi

input_dir="$1"
output_dir="$2" 

if [ ! -d "$input_dir" ]; then
  echo "Input directory $input_dir does not exist"
  exit 1  
fi

if [ ! -d "$output_dir" ]; then
  echo "Output directory $output_dir does not exist"
  exit 1
fi

for file in "$input_dir"/*; do
  if [ -f "$file" ] && [[ "$file" == *.bz2 ]]; then
    output_file="$output_dir/${file##*/}.txt"
    if [ -f "$output_file" ]; then
      echo "Skipping existing file $output_file"
      continue
    fi
    bgpscanner "$file" >> "$output_file"
    echo "Success: $output_file"
  fi
done