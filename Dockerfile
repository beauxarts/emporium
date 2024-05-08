FROM golang:alpine as build
RUN apk add --no-cache --update git
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get ./...
RUN go build \
    -a -tags timetzdata \
    -o emp \
    -ldflags="-s -w -X 'github.com/boggydigital/emporium/cli.GitTag=`git describe --tags --abbrev=0`'" \
    main.go

# adding emporium
FROM alpine:latest
COPY --from=build /go/src/app/emp /usr/bin/emp
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 1927

# backups
VOLUME /usr/share/emporium/backups
# input
VOLUME /usr/share/emporium/share
# metadata
VOLUME /usr/share/emporium/metadata

ENTRYPOINT ["/usr/bin/emp"]
CMD ["serve","-port", "1927", "-stderr"]

