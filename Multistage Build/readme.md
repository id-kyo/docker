# PASSOS PARA INICIAR ESSE DOCKERFILE
1. docker build -t multistage:golang .
2. docker run -d -p 8080:80 multistage:golang


# PASSOS PARA TESTAR A IMAGEM CRIADA
1. http://localhost:8080


A base da imagem está em binário, por conta do Multistaging, então você não consegue usar as funcionalidades da imagem golang, por exemplo. Apenas as funcionalidades que já existem na imagem Scratch.