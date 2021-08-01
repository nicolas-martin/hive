FROM golang:1.16

WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

COPY . .

# go build -tags -o ../bin/$(BINARY)-$@.mipsle
RUN go build -o scheduler ./cmd/main.go

EXPOSE 8081

CMD [ "./scheduler" ]
