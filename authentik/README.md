# <img alt="" src="https://cdn.jsdelivr.net/gh/walkxcode/dashboard-icons/png/authentik.png" height="32px"> authentik stack

## Authentik

[Page](https://goauthentik.io/) |
[Repository](https://github.com/goauthentik/authentik) |
[Image](https://github.com/goauthentik/authentik/pkgs/container/server) |
[Docker Compose](https://goauthentik.io/docs/installation/docker-compose)

authentik is an open-source Identity Provider, focused on flexibility and
versatility. With authentik, site administrators, application developers, and
security engineers have a dependable and secure solution for authentication in
almost any type of environment. There are robust recovery actions available for
the users and applications, including user profile and password management. You
can quickly edit, deactivate, or even impersonate a user profile, and set a new
password for new users or reset an existing password.

### Traefik's ``cors-allow-all@file`` middleware

Example of the middleware settings is presented below. If you want to use
different source name, for example because you store it in something else than
file, then use ``AUTHENTIK_CORS_MIDDLEWARE`` environment variable to overwrite
middleware used.

```yaml
http:
  middlewares:
    cors-allow-all:
      headers:
        accessControlAllowMethods:
          - "*"
        accessControlAllowHeaders: "*"
        accessControlAllowOriginList:
          - "*"
        accessControlMaxAge: 100
        addVaryHeader: true
```
