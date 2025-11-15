#!/bin/bash

year=$1
quest=$2

files=("$year/quest$quest/part1.go")

if [ -f "$year/quest$quest/part2.go" ]; then
    files+=("$year/quest$quest/part2.go")
fi

if [ -f "$year/quest$quest/part3.go" ]; then
    files+=("$year/quest$quest/part3.go")
fi

if [ -f "$year/quest$quest/helpers.go" ]; then
    files+=("$year/quest$quest/helpers.go")
fi

output=$(go vet "${files[@]}" 2>&1)

if [ $? -ne 0 ]; then
    echo "go vet failed for $year/quest$quest/"
    echo "$output"
    exit 1
else
    echo -e "go vet $year/quest$quest/ \033[32m✔\033[0m"
fi