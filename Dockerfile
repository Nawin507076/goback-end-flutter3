FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o api

EXPOSE 8090

ENTRYPOINT [ "./api" ]

# docker build -t 507076/nawin-api . 
# docker run -itd -p 7201:8090 --name ake-api ake-img-api 