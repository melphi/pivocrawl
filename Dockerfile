#
# Build phase
#

FROM golang:1.20-buster as build

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apt-get update && apt-get install -y \
  protobuf-compiler \
  golang-goprotobuf-dev \
  && rm -rf /var/lib/apt/lists/*

COPY ./pivocrawl /work/pivocrawl

WORKDIR /work/pivocrawl
RUN go mod download
RUN go clean && go build -o pivocrawl /work/pivocrawl/cmd/main.go

#
# Package phase
#

FROM golang:1.20-alpine 

# Required by some cloud environments.
RUN apk --no-cache add ca-certificates && update-ca-certificates  

WORKDIR /app
RUN chown 1001 /app \
    && chmod "g+rwX" /app \
    && chown 1001:root /app

COPY --from=build --chown=1001:root /work/pivocrawl /app/pivocrawl

EXPOSE 8080
USER 1001

ENTRYPOINT []
CMD [ "/app/pivocrawl" ]

