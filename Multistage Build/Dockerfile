# Etapa 1: Construção da aplicação
FROM node:14 AS builder
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .
RUN npm run build

# Etapa 2: Execução da aplicação
FROM scratch:latest
WORKDIR /app
#Chamando o primeiro "FROM" as Builder. Para formar as imagens em binário
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/package.json ./
RUN npm install --only=production
CMD ["node", "dist/index.js"]
