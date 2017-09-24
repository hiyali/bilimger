#!/bin/bash

# FIXME must change ./router file name
go build *.go && ./router -env prod
