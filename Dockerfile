FROM golang:alpine as build
RUN apk add --no-cache --update git
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get ./...
RUN go build \
    -a -tags timetzdata \
    -o emporium \
    -ldflags="-s -w -X 'github.com/boggydigital/emporium/cli.GitTag=`git describe --tags --abbrev=0`'" \
    main.go

# adding emporium
FROM alpine:latest
COPY --from=build /go/src/app/emporium /usr/bin/emporium
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 1927

# backups
VOLUME /usr/share/emporium/backups
# shares
VOLUME /usr/share/emporium/shares
# metadata
VOLUME /usr/share/emporium/metadata

ENTRYPOINT ["/usr/bin/emporium"]
CMD ["serve","-port", "1927", "-stderr"]

