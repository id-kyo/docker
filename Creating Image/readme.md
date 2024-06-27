# Criando uma Imagem Base a partir da Documentação Oficial do Dockerfile
Este guia proporciona um fluxo claro e fácil de seguir para criar e executar sua imagem Docker.

Como começar a partir desse Dockerfile? Simples! Siga os passos abaixo:

## Passos para Criar e Executar a Imagem Docker

1. Na pasta do Dockerfile, faça o build da imagem:  
    $ docker build -t aplicacao-teste .

2. Em seguida, inicie sua aplicação usando:  
    $ docker run --name aplicacao-teste -d -p 127.0.0.1:80:8080 aplicacao-teste

3. Agora, vamos entrar no container:  
    $ docker container exec -it aplicacao-teste /bin/sh

4. Se tudo estiver certo, você está dentro do container :)

5. Teste a conectividade dentro do seu container usando curl:  
    $ curl 127.0.0.1

Você deve receber o retorno do Nginx.