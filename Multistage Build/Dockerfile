# Stage 1: Build
FROM golang:1.18.0-bullseye AS builder

# muda o workdir para /app
WORKDIR /app

# copia o app.go local para a imagem.
COPY /app.go .

# roda o "go mod init [name]", necessário para fazer o go build.
RUN go mod init main && CGO_ENABLED=0 go build



# começando o stage
FROM scratch AS stage

# expõe a porta 80 do container
EXPOSE 80

# copia do golang a aplicação que foi criada pelo RUN, em binários.
COPY --from=builder /app/main /go/main

# vai fazer a execução do main. dessa forma vai criar o webserver
CMD ["/go/main"]