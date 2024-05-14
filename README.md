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

- [ ] Add new and already deployed stacks
- [ ] Update all stacks to newest version
- [x] Environment variables how-to
- [x] Requirements and installation how-to
- [ ] Reusable parts
- [ ] Migrate projects description and metadata to separate `README.md` files
- [ ] Readme files generator based on config files
- [ ] Deployment environment agnostic projects
- [ ] Enable configuration of traefik router rules
- [ ] Enable more fine-grained configuration of used volumes

## Setup and prerequisites

Before using this repository's Docker Compose projects please first make sure
you have prepared the target machine with required
[environment variables](ENVIRONMENT.md) and went through the
[setup process](SETUP.md).

## Compose projects statuses

### Projects list

| Project name | Last updated[^1] |
| ----- | ----- |
| ✔️ actual | 2024-05-14 18:53:59 |
| ✔️ adguardhome | 2024-05-02 20:59:07 |
| ✔️ anonaddy | 2024-05-05 16:28:54 |
| ✔️ archivebox | 2024-02-22 20:01:14 |
| ✔️ aria2 | 2024-04-03 18:16:34 |
| ✔️ arr | 2024-04-04 00:05:08 |
| ✔️ authentik | 2024-05-14 13:39:08 |
| ✔️ diun | 2024-03-12 10:39:05 |
| ✔️ docker-mailserver | 2024-03-22 12:39:23 |
| 🗄️ dockge | 2024-03-01 16:26:25 |
| 🗄️ dokemon | 2024-03-01 19:21:16 |
| ✔️ mosquitto | 2024-03-12 10:37:47 |
| ✔️ esphome | 2024-05-02 19:46:36 |
| ✔️ frigate | 2024-03-22 10:56:49 |
| 🏗️ ghost | 2024-04-09 19:15:07 |
| ✔️ home-assistant | 2024-05-03 21:13:03 |
| ✔️ homebox | 2024-04-16 22:33:45 |
| 🗄️ homepage | 2024-05-03 21:13:03 |
| ✔️ jellyfin | 2024-04-29 19:06:27 |
| 🏗️ lancommander | 2024-03-31 20:46:37 |
| 🏗️ maker-management-platform | 2024-03-19 21:00:22 |
| 🗄️ mastodon-glitch | 2024-02-01 21:16:27 |
| ✔️ minio | 2024-04-16 23:10:06 |
| ✔️ mqttx | 2024-02-09 21:23:03 |
| 🏗️ netbootxyz | 2024-02-11 02:36:49 |
| ✔️ netdata | 2024-02-09 21:21:29 |
| 🏗️ node-red | 2024-02-19 15:51:50 |
| 🏗️ noisedash | 2024-02-12 13:29:22 |
| ✔️ ntfy | 2024-03-13 14:37:55 |
| 🏗️ paperless-ngx | 2024-03-05 20:29:37 |
| ✔️ penpotapp | 2024-04-26 20:56:05 |
| ✔️ portainer | 2024-05-01 22:08:00 |
| 🗄️ rabbitmq | 2024-01-18 18:28:09 |
| ✔️ reactive-resume | 2024-05-05 17:00:13 |
| ✔️ red-discordbot | 2023-10-13 09:01:08 |
| 🗄️ redis | 2024-02-09 21:20:02 |
| 🏗️ registry | 2024-03-01 15:44:55 |
| 🏗️ restic | 2024-04-22 22:08:55 |
| 🏗️ romm | 2024-05-03 20:35:26 |
| 🗄️ siyuan | 2024-02-22 15:30:13 |
| ✔️ stirling-pdf | 2024-02-09 21:18:53 |
| ✔️ traefik | 2024-04-02 21:16:06 |
| 🏗️ upsnap | 2024-02-11 02:35:30 |
| ✔️ uptime-kuma | 2024-02-19 15:51:28 |
| ✔️ vaultwarden | 2024-04-11 17:32:00 |
| ✔️ vscode | 2024-01-29 19:02:36 |
| ✔️ wallabag | 2024-04-29 22:05:08 |
| ✔️ web-check | 2024-03-19 22:59:34 |
| ✔️ zigbee2mqtt | 2024-02-09 21:17:49 |

### Status icons explanation

| Emoji | Status name | Description |
| ----- | ----------- | ----------- |
|   ✔️   | Working     | Proven working in live environment, was or even still is in active use |
|   🏗️   | Draft      | Project is being actively worked on (untested, unreleased or broken by updates) |
|   🗄️   | Archived   | Project was once deployed and working, but since then its updates were abandoned, or it never reached maturity |

## Copyrights

Project is shared under CC0 1.0 Universal license. For details please check
[`LICENSE`](/LICENSE), but in short - do whatever you want with the contents of
this repository.

All software, product names, and their branding are subject to copyright and/or
trademark of their rightful owners, *and all that obligatory legal stuff*.

[^1]: Timezone: Europe/Warsaw (CET)
