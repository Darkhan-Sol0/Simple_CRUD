FROM golang:1.24-alpine
ENV GOPROXY=https://goproxy.io,https://proxy.golang.org,https://gocenter.io,direct

WORKDIR /app

COPY ["./go.mod", "./go.sum", "./"]
RUN go mod download

COPY ./ ./

RUN mkdir -p build
RUN	go build -o ./build/progy ./cmd/main.go

CMD ["./build/progy"]

