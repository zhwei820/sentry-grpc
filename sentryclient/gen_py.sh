#!/usr/bin/env bash

# protoc --go_out=.  --go-grpc_out=. -I.  ./*.proto --go_opt paths=source_relative --go-grpc_opt paths=source_relative

python3 -m grpc_tools.protoc -I. ./*.proto --python_out=.  --grpc_python_out=. 