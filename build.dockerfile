FROM fedora:latest

# /-==================================-\
# ||     Install dependencies         ||
# \-==================================-/

RUN dnf update -y

# install development tools required for minimal development (git/make)
RUN dnf install -y git make gcc g++

# install dependencies required to run giu application
RUN dnf install -y libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel libGL-devel libXxf86vm-devel
RUN dnf install -y gtk3 gtk3-devel

# install go
WORKDIR /
RUN dnf install -y wget
RUN wget https://go.dev/dl/go1.20.1.linux-amd64.tar.gz
RUN tar xvf go1.20.1.linux-amd64.tar.gz
ENV PATH="${PATH}:/go/bin:/gopath/bin"
ENV GOPATH="/gopath"

# install python
RUN dnf install -y python3 python3-devel

# install mingw
RUN dnf install -y mingw64-gcc mingw64-gcc-c++ mingw64-headers mingw64-python3 mingw64-python3-Cython

# set workidr
WORKDIR /app

# move all the stuff into working directory
ADD . /app

RUN mkdir -p /build

# /-==================================-\
# ||          linux build             ||
# \-==================================-/

# go-get pakcages (I recommend using go's vendoring-mode since it makes modules downloading super-fast
# as they are in fact already downloaded and stored by previous command)
RUN make setup

# pre-build binaries to make running them faster
RUN go build -o /build/motorola.bin github.com/TheGreaterHeptavirate/motorola/cmd/motorola

# /-==================================-\
# ||        widnows build             ||
# \-==================================-/

# re-generate flags
#RUN chmod 777 /usr/x86_64-w64-mingw32/sys-root/mingw/bin/python3-config
#RUN go run scripts/flags.go -o pkg/python_integration/flags.go -pycfg /usr/x86_64-w64-mingw32/sys-root/mingw/bin/python3-config
RUN go run scripts/flags.go -o pkg/python_integration/flags.go -ldflags "-static /usr/x86_64-w64-mingw32/sys-root/mingw/lib/libpython3.10.dll.a -L/usr/x86_64-w64-mingw32/sys-root/mingw/lib -lpthread -lm -lversion -lshlwapi -lm"  -cflags "-I/usr/x86_64-w64-mingw32/sys-root/mingw/include/python3.10 -I/usr/x86_64-w64-mingw32/sys-root/mingw/include/python3.10  -Wno-unused-result -Wsign-compare  -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions --param=ssp-buffer-size=4 -DNDEBUG -g -O3 -Wall"
ENV CGO_ENABLED=1
ENV CC=/usr/bin/x86_64-w64-mingw32-gcc
ENV CXX=/usr/bin/x86_64-w64-mingw32-g++
ENV GOOS=windows
ENV GOARCH=amd64
RUN go build -o /build/motorola.exe github.com/TheGreaterHeptavirate/motorola/cmd/motorola

WORKDIR /build
CMD bash
