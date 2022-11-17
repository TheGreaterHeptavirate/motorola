motorola project is an application written for
[Motorola Science Cup](https://science-cup.pl/) competition

# Description

TODO

# Documentation

Reffer [to docs](./docs)

# Installation instruction

## Pre-requirements

- [go](https://go.dev)
- GCC
- mingw (**for cross-platform compilation only**)

## Source

```sh
# download source
git clone git@github.com:TheGreaterHeptavirate/motorola
# change-dir
cd motorola
# download go dependencies
go get -d ./...

# run app:
go run github.com/TheGreaterHeptavirate/motorola/cmd/motorola
# OR
cd cmd/motorola
go run .
```

