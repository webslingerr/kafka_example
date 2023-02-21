# workspace (GOPATH) configured at /go
FROM golang:1.15 as builder


#
RUN mkdir -p $GOPATH/src/gitlab.udevs.io/car24/car24_go_car_service
WORKDIR $GOPATH/src/gitlab.udevs.io/car24/car24_go_car_service

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/car24_go_car_service /

FROM gitlab.udevs.io:5050/docker/envconsul:alpine
COPY --from=builder car24_go_car_service .
CMD ["envconsul", "-config", "./config.hcl", "./car24_go_car_service"]
