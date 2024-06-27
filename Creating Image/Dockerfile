FROM nginx:latest AS base

# Criação do diretório de trabalho
WORKDIR /app/dockercreated

# Criação de um arquivo e adição de texto nele
RUN touch test.txt && \
    echo "teste nginx" >> test.txt && \
    echo -e "are you seeking for the text? try on /app/dockercreated" &&\
    mkdir anotherone

# Adição de um arquivo local para o diretório de trabalho
ADD ./localfile.txt /app/dockercreated/localfile.txt

# Criação de outro diretório e cópia do arquivo para ele
WORKDIR /app/dockercreated/anotherone
COPY ./localfile.txt /app/dockercreated/anotherone/

# Exposição da porta 8080 para o container
EXPOSE 8080