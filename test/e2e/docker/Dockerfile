# We need to build in a Linux environment to support C libraries, e.g. RocksDB.
# We use Debian instead of Alpine, so that we can use binary database packages
# instead of spending time compiling them.
FROM golang:1.18-bullseye

RUN apt-get -qq update -y && apt-get -qq upgrade -y >/dev/null
RUN apt-get -qq install -y cmake sudo libgmp-dev libleveldb-dev librocksdb-dev >/dev/null

# Set up build directory /src/tenderdash
WORKDIR /src/tenderdash

# Install DLV debugger
ENV DEBUG ""
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Fetch dependencies separately (for layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Install BLS library
COPY third_party ./third_party
COPY test/Makefile ./test/Makefile
COPY Makefile ./
RUN make install-bls

# Copy Tenderdash source
COPY . .

COPY test/e2e/docker/entrypoint* /usr/bin/

RUN cd test/e2e && make node && cp build/node /usr/bin/app

# Set up runtime directory. We don't use a separate runtime image since we need
# e.g. leveldb and rocksdb which are already installed in the build image.
WORKDIR /tenderdash
VOLUME /tenderdash
ENV TMHOME=/tenderdash
ENV GOTRACEBACK=crash

EXPOSE 26656 26657 26660 6060
ENTRYPOINT ["/usr/bin/entrypoint"]
CMD ["start"]
STOPSIGNAL SIGTERM
