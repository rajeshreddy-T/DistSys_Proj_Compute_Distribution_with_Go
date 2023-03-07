FROM golang:latest

WORKDIR /go/src/github.com/rajeshreddyt/GrpcServer

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["GrpcServer"]


# docker build -t mygrpcserver .
# docker run -p 50051:50051 mygrpcserver
# protoc --go_out=. path/to/your.proto

# protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative pkg/gopher/gopher.proto
# 
# run Hello.proto

# protoc --go_out=plugins=grpc:. --go_opt=paths=./ hello/Hello.proto

# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./Hello.proto

# github.com/rajeshreddyt/GrpcServer
# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative Hello/Hello.proto



# docker run -p 50051:50051 mygrpcserver