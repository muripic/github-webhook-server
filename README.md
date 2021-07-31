# github-webhook-server

Process GitHub events for fun and profit.

## Development

### Environment

For the next steps, you need to have Docker and psql installed.

You also need to create a [ngrok][] account and generate an auth token.

1. Create a docker volume named `ìssuesdb` and copy the `structure.sql` file there. This volume will later be mounted to the container and create the structure in the database automatically.

```bash
docker volume create issuesdb
sudo cp structure.sql /var/lib/docker/volumes/issuesdb/_data/structure.sql
```

2. Run docker-compose. This will create 3 containers: [ngrok][], postgresql and the github-webhook-server.

```bash
export NGROK_TOKEN=<your token goes here>
docker-compose up
```

3. If you want to make sure that the database was created correctly, or change your password, you can connect to it with psql:

```bash
psql issues -h localhost -U admin
```

## Roadmap

- Add proper logging.
- Add unit and integration tests.
- Make actions configurable.
- Learn how to actually code in GO and refactor everything xD

## Acknowledgements

Tutorials and articles that helped me build this project:

- https://groob.io/tutorial/go-github-webhook/
- https://github.com/olliefr/docker-gs-ping

[ngrok]: https://ngrok.com/
