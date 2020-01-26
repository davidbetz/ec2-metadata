FROM ubuntu as build

RUN apt-get update

RUN apt install git golang -y

WORKDIR /src/aws

RUN go get -u github.com/aws/aws-sdk-go/...

COPY metadata.go .

RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"'  -o metadata

FROM scratch

COPY --from=build /src/aws/metadata /

ENTRYPOINT ["./metadata"]
