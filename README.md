motorola project is an application written for
[Motorola Science Cup](https://science-cup.pl/) competition

# Description

TODO

# Documentation

Reffer [to docs](./docs)

# STATUS

In reference to [this document](https://science-cup.pl/wp-content/uploads/2022/11/MSC3_2022_Bioinformatyka.pdf)
- [X] Genetic code reading (35 pts)
- [X] Proteins presentation (25 pts)
- [X] More plots and charts (65 pts)
- [ ] Additional things:
    - [ ] "Application architecture" (10 pts)
    - [ ] documentation (10 pts) (**CAUTION** the current documentation needs to be rewriten! It contains bad language IIRC.)
    - [X] UI (25 pts)
    - [ ] Unittesting :smile: (5 pts)

# Installation instruction

## Pre-requirements

- [go](https://go.dev)
- GCC
- mingw (**for cross-platform compilation only**)
- Python 3.11 version **with C headers** (you can test if another versions works)

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

## Alternative - `Docker`

As compilation in the way described above may be a bit painful on some operating systems (especially Windows :smile:)
We've introduced another way to run our application - [Docker](https://docker.io).

```sh
# start immedietly
docker-compose up
# you can also construct a large docker command like this:
docker run --name=motorola_app_1 \
--security-opt label:type:container_runtime_t \
--network bridge \
-e DISPLAY=:0 \
-v /tmp/.X11-unix:/tmp/.X11-unix .
```

**NOTE** commands above are tested for [podman](https://podman.io),
but since it has the same api as docker, everything should work.

### important linux note

you need to disable access control for your X envirouemnt, otherwise it will not run:
```sh
xhost +
```
