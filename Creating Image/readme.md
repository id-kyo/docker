Criando uma imagem base a partir da documentação oficial do Dockerfile.

Como começar, a partir desse Dockerfile? Simple! Siga os passos: (o "$" significa que é dentro do cli)

1.) na pasta do dockerfile faça build da imagem:
    1.1) $ docker build -t aplicacao-teste .

2.) Em seguida inicie sua aplicação, usando:
    2.1) $ docker run --name aplicacao-teste -dp 127.0.0.1:8080:8080 aplicacao-teste

3.) agora vamos entrar no container
    3.1) $ docker container exec -it aplicacao-teste /bin/sh
    3.2) se tudo estiver certo você está dentro do container :)

4.) teste a conectividade dentro do seu container usando curl
    4.1) $ curl 127.0.0.1:8080
    4.2) você deve receber o retorno do Nginx.