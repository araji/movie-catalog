#!/bin/bash
oapi-codegen --package main --old-config-style --generate "types,server,spec" swagger.yaml > apiserver.generated.go
