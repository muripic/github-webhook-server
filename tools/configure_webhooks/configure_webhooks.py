import logging
from typing import Any, Optional

import click
import yaml
from cerberus import Validator  # type: ignore
from github import Github

HOOK_SCHEMA = {
    "endpoint": {"type": "string", "required": True},
    "events": {"type": "list", "required": True},
}

CONFIG_SCHEMA = {
    "repository": {"type": "string", "required": True},
    "token": {"type": "string", "required": True},
    "base_url": {"type": "string", "required": True},
    "hooks": {"type": "list"},
}

logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")


class Config:
    def __init__(self):
        self._config = {}

    def load(self, path: str):
        with open(path) as config_file:
            self._config = yaml.safe_load(config_file)

    def _validate(self, data: dict, schema: dict):
        v = Validator(schema, allow_unknown=True)
        if not v.validate(data):
            raise ValueError(f"Config error: {v.errors}")

    def validate(self):
        self._validate(self._config, CONFIG_SCHEMA)
        hooks = self._config.get("hooks")
        if not hooks:
            logging.warning("No hooks configured.")
            return
        for hook_config in hooks:
            self._validate(hook_config, HOOK_SCHEMA)

    def __getattr__(self, key: str) -> Optional[Any]:
        return self._config.get(key)


class GitRepository:
    def __init__(self):
        self._repo = None

    def get(self):
        if not self._repo:
            github = Github(config.token)
            self._repo = github.get_repo(config.repository)

    def create_hook(self, url: str, events: list):
        config = {"url": url, "content_type": "json"}
        self._repo.create_hook("web", config, events, active=True)

    def create_hooks(self, hooks: list, base_url: str):
        self.get()
        logging.info("Creating new hooks for repository %s", config.repository)
        for hook in hooks:
            self.create_hook(
                f"{self._url(base_url, hook['endpoint'])}",
                hook["events"]
            )

    def _url(self, base_url: str, endpoint: str) -> str:
        return base_url.strip("/") + "/" + endpoint.strip("/")

    def delete_hooks(self):
        self.get()
        logging.info("Deleting all hooks from repository %s", config.repository)
        configured_hooks = [h for h in self._repo.get_hooks()]
        for h in configured_hooks:
            h.delete()


config = Config()
repo = GitRepository()


@click.group()
@click.option("--config-file", required=True, type=click.Path(exists=True))
def cli(config_file):
    config.load(config_file)
    config.validate()


@click.command()
def create():
    repo.create_hooks(config.hooks, config.base_url)
    logging.info("Hooks created successfully.")


@click.command()
def delete():
    repo.delete_hooks()
    logging.info("Hooks deleted successfully.")


cli.add_command(create)
cli.add_command(delete)


if __name__ == "__main__":
    cli()
