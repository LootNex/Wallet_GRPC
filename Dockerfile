FROM golang

WORKDIR /wallet_grpc

COPY . .

RUN go mod tidy

RUN go build -o app ./cmd

EXPOSE 50051

CMD [ "./app" ]



