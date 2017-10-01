FROM golang:1.9

WORKDIR /go/src/github.com/DNewt/LifeTrack

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -i -o bin/linux/LifeTrack ./cmd

FROM scratch
MAINTAINER David Newton <DavidWillNewton@gmail.com>

COPY --from=0 /go/src/github.com/DNewt/LifeTrack/bin/linux/LifeTrack .

ENTRYPOINT ["./LifeTrack"]