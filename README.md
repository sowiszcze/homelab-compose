# homelab-compose

Docker Compose files for services used by me as part of my homelab.

## Project status

Current status of the project I would call "uder heavy development", and
so of course it is **not production ready**.

There are already some archived (or abandonned, call it however) compose
projects that got substituted for other solutions better fitting the needs of my
homelab, or they just didn't prove themselve useful. They are left in the repo
though, so people can have some idea how those can be deployed.

### Outstanding tasks

> [!NOTE]
> Not sorted in any way.

- [x] Add new and already deployed stacks
- [ ] Update all stacks to newest version
- [x] Environment variables how-to
- [x] Requirements and installation how-to
- [ ] Add healthcheck where possible and it makes sense
- [ ] Reusable parts
- [ ] Migrate projects description and metadata to separate files
- [ ] Readme files generator based on config files
- [ ] Deployment environment agnostic projects
- [ ] Enable configuration of traefik router rules
- [ ] Enable more fine-grained configuration of used volumes

## Setup and prerequisites

Before using this repository's Docker Compose projects please first make sure
you have prepared the target machine with required
[environment variables](ENVIRONMENT.md) and went through the
[setup process](SETUP.md).

## Copyrights

Project is shared under CC0 1.0 Universal license. For details please check
[`LICENSE`](/LICENSE), but in short - do whatever you want with the contents of
this repository.

All software, product names, and their branding are subject to copyright and/or
trademark of their rightful owners, *and all that obligatory legal stuff*.

[^1]: Timezone: Europe/Warsaw (CET)
