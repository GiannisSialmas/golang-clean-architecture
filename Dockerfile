FROM golang:1.13-alpine AS build

# Metadata for the build, compile only for linux target, smaller binary size
ENV GOOS linux
ENV GOARCH amd64
# Turn off CGO since that can result in dynamic links to libc/libmusl.
ENV CGO_ENABLED 0 

WORKDIR /app 

#The certificates are needed cause scratch will keep giving "x509: certificate signed by unknown authority" on http requests
RUN apk --no-cache add ca-certificates

# Copy `go.mod` for definitions and `go.sum` to invalidate the next layer
# in case of a change in the dependencies
COPY src/go.mod src/go.sum ./
# Download dependencies
RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

# Copy and build the app
COPY src/ .
# RUN go build -ldflags="-w -s" -o ./main ./main.go
# ENTRYPOINT CompileDaemon --build="go build -o ./main ./main.go" --command=./main
ENTRYPOINT CompileDaemon --build="go build -o /tmp/main ./main.go" --command=/tmp/main



FROM scratch
COPY --from=build /app/main /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./main"]  
