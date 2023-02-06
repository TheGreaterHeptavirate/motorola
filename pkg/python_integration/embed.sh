#!/bin/bash

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <directory>" >&2
  exit 1
fi

dir=$1

echo -e "package python\nvar pythonFiles = []string{"
for file in $(find $dir -type f -name "*.py"); do
  echo "  \"$file\","
done
echo "}"
