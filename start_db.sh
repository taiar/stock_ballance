docker run --rm  --name pg-docker -e POSTGRES_PASSWORD=postgres -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres:11.2-alpine
