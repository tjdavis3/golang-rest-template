#!/bin/sh
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 {module}"
  exit 1
fi
[ -f go.mod ] && rm go.mod
go mod init $1
find . -name "*.go" -exec sed -i "" 's,boilerplate,'"$1"',g' {} \;
