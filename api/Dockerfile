FROM golang:1.16

WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

COPY . .

# go build -tags -o ../bin/$(BINARY)-$@.mipsle
RUN go build -o api ./cmd/api/api.go

EXPOSE 8080

CMD [ "./api" ]
