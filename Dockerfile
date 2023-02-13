FROM docker.io/library/golang:1.19.3-buster

RUN apt-get update

# install dependencies required to run giu application
RUN apt-get install -y libgtk-3-dev libasound2-dev libxxf86vm-dev

RUN apt-get install -y git
RUN git clone https://github.com/python/cpython /python
WORKDIR /python
RUN ./configure
RUN make install

# set workidr
WORKDIR /app

# move all the stuff into working directory
ADD . /app

# go-get pakcages (I recommend using go's vendoring-mode since it makes modules downloading super-fast
# as they are in fact already downloaded and stored by previous command)
RUN make setup



# pre-build binaries to make running them faster
RUN go build github.com/TheGreaterHeptavirate/motorola/cmd/motorola

# define command to run
CMD go run github.com/gucio321/tic-tac-go/cmd/giu-game
