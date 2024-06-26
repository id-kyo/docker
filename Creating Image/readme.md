Criando uma imagem base a partir da documentação oficial do Dockerfile.

Como começar, a partir desse Dockerfile? Simple! Siga os passos: (o "$" significa que é dentro do cli)

1.) na pasta do dockerfile faça build da imagem:

$ docker build -t aplicacao-teste .

3.) Em seguida inicie sua aplicação, usando:

$ docker run --name aplicacao-teste -dp 127.0.0.1:8080:8080 aplicacao-teste

5.) agora vamos entrar no container

$ docker container exec -it aplicacao-teste /bin/sh

7.) se tudo estiver certo você está dentro do container :)

8.) teste a conectividade dentro do seu container usando curl

$ curl 127.0.0.1:8080

10.) você deve receber o retorno do Nginx.
