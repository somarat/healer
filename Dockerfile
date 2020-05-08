FROM golang:1.12.3-stretch
MAINTAINER Will Schenk <wschenk@gmail.com>

# Get the TLS CA certificates, they're not provided by busybox.
RUN apt-get update && apt-get install -y ca-certificates

# Copy the single source file to the app directory
WORKDIR /go/src/app
COPY . .

# Install depenancies
RUN go get -d

# Build the app
RUN go build

# Switch to a small base image
FROM busybox:1-glibc
MAINTAINER Will Schenk <wschenk@gmail.com>

# Copy the binary over from the deploy container
COPY --from=0 /go/src/app/app /usr/bin/healer

# Get the TLS CA certificates from the build container, they're not provided by busybox.

COPY --from=0 /etc/ssl/certs /etc/ssl/certs

CMD healer
