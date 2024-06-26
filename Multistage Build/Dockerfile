# Etapa 1: Construção da aplicação
FROM node:14 AS builder
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .
RUN npm run build

# Etapa 2: Execução da aplicação
FROM node:14-slim
WORKDIR /app
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/package.json ./
RUN npm install --only=production
CMD ["node", "dist/index.js"]
