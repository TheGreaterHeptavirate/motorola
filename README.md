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
- [ ] More plots and charts (65 pts)
- [ ] Additional things:
    - [ ] "Application architecture" (10 pts)
    - [ ] documentation (10 pts) (**CAUTION** the current documentation needs to be rewriten! It contains bad language IIRC.)
    - [ ] UI (25 pts)
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
# setup python dependencies (ofc you can use virtualenv)
python3 -m pip install -r requirements.txt

# run app:
go run github.com/TheGreaterHeptavirate/motorola/cmd/motorola
# OR
cd cmd/motorola
go run .
```

