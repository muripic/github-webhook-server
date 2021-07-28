# github-webhook-server

Process GitHub events for fun and profit.

## Development

### Environment

For the next steps, you need to have Docker and psql installed.

1. Create issue-db docker volume and copy `structure.sql` file. This volume will later be mounted to the container and create the structure in the database automatically.

```bash
docker volume create issue-db
sudo cp structure.sql /var/lib/docker/volumes/issue-db/_data/structure.sql
```

2. Create postgres database container:

```bash
docker run --name issues-db -p 5432:5432 -d \
    -e POSTGRES_PASSWORD=changeme \
    -e POSTGRES_USER=admin \
    -e POSTGRES_DB=issues \
    -v pgdata:/var/lib/postgresql/data \
    -v issue-db:/docker-entrypoint-initdb.d \
    postgres
```

3. Connect to the database:

```bash
psql issues -h localhost -U admin
```

## Roadmap

- Add proper logging.

## Acknowledgements

Tutorials and articles that helped me build this project:

- https://groob.io/tutorial/go-github-webhook/
- https://github.com/olliefr/docker-gs-ping
