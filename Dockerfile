FROM golang:1.9.2-alpine3.7

WORKDIR /go/src/app
COPY . .

RUN apk --no-cache add -t build-deps build-base git \
	&& apk --no-cache add ca-certificates \
  && git config --global http.https://gopkg.in.followRedirects true 

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
