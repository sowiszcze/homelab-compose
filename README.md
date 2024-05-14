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
| ------------ | ---------------- |
| ✔️ [actual](actual) | 2024-05-14 18:53:59 |
| ✔️ [adguardhome](adguardhome) | 2024-05-14 19:24:06 |
| ✔️ [adguardhome-sync](adguardhome-sync) | 2024-05-03 10:38:19 |
| 🗄️ [affine](affine) | 2024-04-13 19:40:57 |
| ✔️ [anonaddy](anonaddy) | 2024-05-14 19:47:46 |
| 🏗️ [apprise](apprise) | 2024-05-12 14:09:33 |
| ✔️ [archivebox](archivebox) | 2024-05-14 19:52:03 |
| ✔️ [aria2](aria2) | 2024-04-03 18:16:34 |
| ✔️ [arr](arr) | 2024-04-04 00:05:08 |
| ✔️ [authentik](authentik) | 2024-05-14 19:56:53 |
| 🏗️ [backrest](backrest) | 2024-03-01 20:01:10 |
| 🏗️ [changedetection](changedetection) | 2024-04-16 20:51:31 |
| 🏗️ [cyberchef](cyberchef) | 2024-04-17 19:18:37 |
| ✔️ [czkawka](czkawka) | 2024-05-04 14:47:01 |
| ✔️ [dashy](dashy) | 2024-05-01 22:09:02 |
| 🗄️ [directus](directus) | 2024-04-15 23:17:44 |
| ✔️ [diun](diun) | 2024-03-12 10:39:05 |
| ✔️ [docker-mailserver](docker-mailserver) | 2024-03-22 12:39:23 |
| 🗄️ [dockge](dockge) | 2024-03-01 16:26:25 |
| 🗄️ [dokemon](dokemon) | 2024-03-01 19:21:16 |
| ✔️ [dozzle](dozzle) | 2024-05-03 17:00:30 |
| ✔️ [mosquitto](mosquitto) | 2024-03-12 10:37:47 |
| 🗄️ [elasticsearch](elasticsearch) | 2024-04-12 15:13:42 |
| 🏗️ [epicgames-freebies-claimer](epicgames-freebies-claimer) | 2024-05-10 20:36:10 |
| ✔️ [esphome](esphome) | 2024-05-02 19:46:36 |
| 🗄️ [filestash](filestash) | 2024-04-11 21:24:31 |
| 🗄️ [flame](flame) | 2024-04-06 16:38:53 |
| 🗄️ [flatnotes](flatnotes) | 2024-02-13 21:33:34 |
| 🗄️ [flyimg](flyimg) | 2024-04-17 21:16:48 |
| 🏗️ [free-games-claimer](free-games-claimer) | 2024-05-10 21:26:18 |
| ✔️ [frigate](frigate) | 2024-03-22 10:56:49 |
| 🏗️ [ghost](ghost) | 2024-04-09 19:15:07 |
| 🏗️ [glances](glances) | 2024-04-30 21:28:23 |
| ✔️ [headscale](headscale) | 2024-04-25 21:31:07 |
| 🗄️ [heimdall](heimdall) | 2024-04-06 15:35:52 |
| ✔️ [home-assistant](home-assistant) | 2024-05-03 21:13:03 |
| ✔️ [homebox](homebox) | 2024-04-16 22:33:45 |
| 🗄️ [homepage](homepage) | 2024-05-03 21:13:03 |
| 🗄️ [homer](homer) | 2024-04-06 15:52:51 |
| 🗄️ [hrconvert2](hrconvert2) | 2024-04-17 19:13:15 |
| 🗄️ [huginn](huginn) | 2024-04-16 21:35:48 |
| ✔️ [immich](immich) | 2024-04-30 12:54:37 |
| ✔️ [it-tools](it-tools) | 2024-04-09 22:01:32 |
| ✔️ [jellyfin](jellyfin) | 2024-04-29 19:06:27 |
| 🏗️ [lancommander](lancommander) | 2024-03-31 20:46:37 |
| ✔️ [languagetool](languagetool) | 2024-04-18 19:19:17 |
| 🗄️ [leantime](leantime) | 2024-05-13 13:48:18 |
| ✔️ [libretranslate](libretranslate) | 2024-05-04 21:24:37 |
| 🏗️ [linguacafe](linguacafe) | 2024-04-09 21:28:25 |
| 🏗️ [linkding](linkding) | 2024-04-17 20:19:59 |
| ✔️ [linkstack](linkstack) | 2024-05-09 18:32:05 |
| 🗄️ [logseq](logseq) | 2024-02-14 21:00:05 |
| 🗄️ [maddy](maddy) | 2024-03-11 14:42:44 |
| 🏗️ [maker-management-platform](maker-management-platform) | 2024-03-19 21:00:22 |
| 🗄️ [mastodon-glitch](mastodon-glitch) | 2024-02-01 21:16:27 |
| ✔️ [minio](minio) | 2024-04-16 23:10:06 |
| 🏗️ [monica](monica) | 2024-04-17 19:53:32 |
| ✔️ [mqttx](mqttx) | 2024-02-09 21:23:03 |
| ✔️ [music-assistant](music-assistant) | 2024-04-18 20:53:50 |
| 🗄️ [n8n](n8n) | 2024-04-15 22:34:18 |
| ✔️ [netalertx](netalertx) | 2024-04-24 16:08:28 |
| 🏗️ [netbootxyz](netbootxyz) | 2024-02-11 02:36:49 |
| ✔️ [netdata](netdata) | 2024-02-09 21:21:29 |
| ✔️ [nextcloud](nextcloud) | 2024-05-04 16:58:20 |
| 🏗️ [node-red](node-red) | 2024-02-19 15:51:50 |
| 🏗️ [noisedash](noisedash) | 2024-02-12 13:29:22 |
| ✔️ [ntfy](ntfy) | 2024-03-13 14:37:55 |
| ✔️ [octoprint](octoprint) | N/A |
| ✔️ [outline](outline) | 2024-04-12 21:56:02 |
| 🏗️ [paperless-ngx](paperless-ngx) | 2024-03-05 20:29:37 |
| ✔️ [penpotapp](penpotapp) | 2024-04-26 20:56:05 |
| 🗄️ [photoprism](photoprism) | 2024-04-29 18:25:42 |
| ✔️ [planka](planka) | 2024-05-13 18:57:11 |
| ✔️ [portainer](portainer) | 2024-05-01 22:08:00 |
| 🗄️ [postal](postal) | 2024-03-11 15:01:44 |
| 🗄️ [rabbitmq](rabbitmq) | 2024-01-18 18:28:09 |
| ✔️ [reactive-resume](reactive-resume) | 2024-05-05 17:00:13 |
| ✔️ [red-discordbot](red-discordbot) | 2023-10-13 09:01:08 |
| 🗄️ [redis](redis) | 2024-02-09 21:20:02 |
| 🏗️ [registry](registry) | 2024-03-01 15:44:55 |
| 🏗️ [restic](restic) | 2024-04-22 22:08:55 |
| 🏗️ [romm](romm) | 2024-05-03 20:35:26 |
| ✔️ [sabnzbd](sabnzbd) | 2024-04-03 21:26:49 |
| 🗄️ [siyuan](siyuan) | 2024-02-22 15:30:13 |
| 🏗️ [sshwifty](sshwifty) | 2024-04-18 19:50:25 |
| 🗄️ [stalwart-smtp](stalwart-smtp) | 2024-03-10 22:01:11 |
| ✔️ [stirling-pdf](stirling-pdf) | 2024-02-09 21:18:53 |
| ✔️ [string-is](string-is) | 2024-04-17 20:32:07 |
| 🏗️ [tooljet](tooljet) | 2024-04-15 22:03:30 |
| ✔️ [traefik](traefik) | 2024-04-02 21:16:06 |
| 🗄️ [trillium](trillium) | 2024-02-13 21:00:55 |
| 🏗️ [upsnap](upsnap) | 2024-02-11 02:35:30 |
| ✔️ [uptime-kuma](uptime-kuma) | 2024-02-19 15:51:28 |
| ✔️ [vaultwarden](vaultwarden) | 2024-04-11 17:32:00 |
| ✔️ [vscode](vscode) | 2024-01-29 19:02:36 |
| ✔️ [wallabag](wallabag) | 2024-04-29 22:05:08 |
| 🏗️ [watchtower](watchtower) | 2024-04-14 19:56:11 |
| ✔️ [web-check](web-check) | 2024-03-19 22:59:34 |
| 🗄️ [wekan](wekan) | 2024-05-13 17:38:36 |
| 🏗️ [whatsupdocker](whatsupdocker) | 2024-04-29 14:39:06 |
| 🗄️ [wikijs](wikijs) | 2024-04-12 20:42:09 |
| ✔️ [zigbee2mqtt](zigbee2mqtt) | 2024-02-09 21:17:49 |

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
