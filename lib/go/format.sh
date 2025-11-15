#!/bin/bash

year=$1
quest=$2

files=("part1.go" "part2.go" "part3.go" "helpers.go")

for file in "${files[@]}"; do
    if [ -f "$year/quest$quest/$file" ]; then
        gofmt -w $year/quest$quest/$file && echo -e "gofmt $year/quest$quest/$file \033[32m✔\033[0m"
    fi
done