#!/bin/bash

if [ $# -ne 1 ]; then
	echo "usage: $0 name"
	exit 1
fi

# change to the directory of this script
cd $(dirname "$0")

name=$(echo "$1" | tr '[:upper:]' '[:lower:]' | tr ' ' '-')

/bin/mkdir ${name}

/bin/cat > ${name}/README.md << EOF
## $1

TODO: description of the program

EOF

/bin/cat > ${name}/main.go << EOF
package main

func main() {

}

EOF

