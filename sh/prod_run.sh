#!/bin/bash

go clean
go build && ./bilimger -env prod
