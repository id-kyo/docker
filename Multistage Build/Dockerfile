# Stage 1: Build
FROM golang:1.18.0-bullseye AS builder

# expõe a porta 80 do container
EXPOSE 80

# muda o workdir para /app
WORKDIR /app

# copia o app.go local para a imagem.
COPY app.go .

# roda o "go mod init [name]", necessário para fazer o go build.
RUN go mod init main && go build 

# roda o /.main no cmd
CMD [ "./main" ]
