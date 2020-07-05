# goprotoconv

A one-to-one converter generator between two protos in different packages used for Golang.

## Requirements

This project actively builds upon protobufs, so it's necessary to have proto compiler and go language generator installed. They must be accessible to the build system.

## Build

I use [please](https://please.build) build system. Run  ```plz build``` to build or ```plz test``` to test.
