#!/bin/bash

year=$1
quest=$2
part=$3
# input_file=$4

curr_dir=$(pwd)

ln -sf $curr_dir/$year/quest$quest/part1.go $curr_dir/lib/go/part1.go

part2_arg=""
if [ -f "$curr_dir/$year/quest$quest/part2.go" ]; then
  ln -sf $curr_dir/$year/quest$quest/part2.go $curr_dir/lib/go/part2.go
  part2_arg="lib/go/part2.go"
fi

part3_arg=""
if [ -f "$curr_dir/$year/quest$quest/part3.go" ]; then
  ln -sf $curr_dir/$year/quest$quest/part3.go $curr_dir/lib/go/part3.go
  part3_arg="lib/go/part3.go"
fi


helpers_arg=""
if [ -f "$curr_dir/$year/quest$quest/helpers.go" ]; then
  ln -sf $curr_dir/$year/quest$quest/helpers.go $curr_dir/lib/go/helpers.go
  helpers_arg="lib/go/helpers.go"
fi

input_file="$curr_dir/$year/quest$quest/part$part.txt"

go run lib/go/part1.go $part2_arg $part3_arg $helpers_arg lib/go/main.go $year $quest $part < "$input_file"