FROM golang:1.17.3-buster
WORKDIR /go/src/ppamo/server
COPY ./ /go/src/ppamo/server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server .

FROM scratch
COPY --from=0 /go/src/ppamo/server/bin/server /server
COPY ./static /static
CMD ["/server"]
