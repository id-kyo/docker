

docker build -t multistage-build:golang .


docker run -v $'(pwd)':/app -w /app -it golang:1.18.0-bullseye sh


docker run -v "C:\Users\ROB - 01\Desktop\Docker\Docker-Git\Multistage Build":/app -w /app -it golang:1.18.0-bullseye sh