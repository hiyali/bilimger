#!/bin/bash

# FIXME must change ./router file name
go clean
go build *.go && ./router -env prod
