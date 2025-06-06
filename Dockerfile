FROM golang
ENV GOPROXY=https://goproxy.io,https://proxy.golang.org,https://gocenter.io,direct

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN make build

CMD ["./build/progy"]

