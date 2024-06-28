# Stage 1: Build
FROM golang:1.16 AS builder

# Definindo o diretório de trabalho no contêiner
WORKDIR /app

# Copiando os arquivos do projeto para o contêiner
COPY . .

# transformando o código em binário. parecido com um .exe
RUN RUN go build -o app.go

# Stage 2: Construção do Run
FROM scratch:latest AS runner

# Definindo o diretório de trabalho no contêiner
WORKDIR /root/

# copia o binário do builder para o runner
COPY --from=builder . .

# Definindo o comando de entrada do contêiner
CMD ["./myapp"]
