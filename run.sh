#!/bin/bash
currDir=$(dirname $0)
(cd "$currDir"/"$1"/"$2" && go run $2.go)