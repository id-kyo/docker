docker build -t multistage:golang .

docker run -d -p 80:8080 multistage:golang

docker run -v $'(pwd)':/app -w /app -it golang:1.18.0-bullseye sh


docker run -v "C:\Users\ROB - 01\Desktop\Docker\Docker-Git\Multistage Build":/app -w /app -it golang:1.18.0-bullseye sh

curl http://localhost:8080