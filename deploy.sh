docker build -t hwrbot:latest build/docker
docker run --rm --env-file .env -it hwrbot