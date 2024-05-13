# *rr stack


## Jellyseerr

~~Page~~ |
[Repository](https://github.com/Fallenbagel/jellyseerr) |
[Image](https://github.com/orgs/hotio/packages/container/package/jellyseerr) |
[Docker Compose](https://hotio.dev/containers/jellyseerr/#__tabbed_1_2)

Jellyseerr is a free and open source software application for managing requests
for your media library. It is a fork of Overseerr built to bring support for
Jellyfin & Emby media servers!


## Radarr

[Page](https://radarr.video/) |
[Repository](https://github.com/Radarr/Radarr) |
[Image](https://github.com/hotio/radarr/pkgs/container/radarr) |
[Docker Compose](https://hotio.dev/containers/radarr/#__tabbed_1_2)

A movie collection manager for Usenet and BitTorrent users. It can monitor
multiple RSS feeds for new movies and will interface with clients and indexers
to grab, sort, and rename them. It can also be configured to automatically
upgrade the quality of existing files in the library when a better quality
format becomes available. Note that only one type of a given movie is supported.
If you want both an 4k version and 1080p version of a given movie you will need
multiple instances.


## Sonarr

[Page](https://sonarr.tv/) |
[Repository](https://github.com/sonarr/sonarr) |
[Image](https://github.com/orgs/hotio/packages/container/package/sonarr) |
[Docker Compose](https://hotio.dev/containers/sonarr/#__tabbed_1_2)

A PVR for Usenet and BitTorrent users. It can monitor multiple RSS feeds for new
episodes of your favorite shows and will grab, sort and rename them. It can also
be configured to automatically upgrade the quality of files already downloaded
when a better quality format becomes available.


## Prowlarr

~~Page~~ |
[Repository](https://github.com/prowlarr/prowlarr) |
[Image](https://github.com/orgs/hotio/packages/container/package/prowlarr) |
[Docker Compose](https://hotio.dev/containers/prowlarr/#__tabbed_1_2)

An indexer manager/proxy built on the popular *arr .net/reactjs base stack to
integrate with your various PVR apps. Prowlarr supports management of both
Torrent Trackers and Usenet Indexers. It integrates seamlessly with Lidarr,
Mylar3, Radarr, Readarr, and Sonarr offering complete management of your
indexers with no per app Indexer setup required (we do it all).


## Jackett

~~Page~~ |
[Repository](https://github.com/jackett/jackett) |
[Image](https://github.com/orgs/hotio/packages/container/package/jackett) |
[Docker Compose](https://hotio.dev/containers/jackett/#__tabbed_1_2)

Works as a proxy server: it translates queries from apps (Sonarr, Radarr,
SickRage, CouchPotato, Mylar3, Lidarr, DuckieTV, qBittorrent, Nefarious etc.)
into tracker-site-specific http queries, parses the html or json response, and
then sends results back to the requesting software. This allows for getting
recent uploads (like RSS) and performing searches. Jackett is a single
repository of maintained indexer scraping & translation logic - removing the
burden from other apps.

## FlareSolverr

~~Page~~ |
[Repository](https://github.com/FlareSolverr/FlareSolverr) |
[Image](https://github.com/orgs/FlareSolverr/packages/container/package/flaresolverr) |
[Docker Compose](https://github.com/FlareSolverr/FlareSolverr/blob/master/docker-compose.yml)

Starts a proxy server, and it waits for user requests in an idle state using few
resources. When some request arrives, it uses Selenium with the
undetected-chromedriver to create a web browser (Chrome). It opens the URL with
user parameters and waits until the Cloudflare challenge is solved (or timeout).
The HTML code and the cookies are sent back to the user, and those cookies can
be used to bypass Cloudflare using other HTTP clients.
