FROM golang:1.15.2-alpine

WORKDIR $GOPATH/src/timescaledb-statistics-go-service

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN apk add --update g++
RUN go mod tidy
RUN cd $GOPATH/src/timescaledb-statistics-go-service/ && go build main.go

# This container exposes port 8187 to the outside world
EXPOSE 8187

RUN chmod 0766 $GOPATH/src/timescaledb-statistics-go-service/scripts/init.sh

# Run the executable
CMD ["./scripts/init.sh"]