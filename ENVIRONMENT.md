# Environment variables setup

All or almost all projects use a predefined set of environment variables. Those
variables have to be set either for the user running ``docker compse`` commands,
or for all users via for example ``/etc/environment`` file.

This means if you for example use Portainer to deploy Docker Compose projects,
you have to pass those environment variables to the agent container.

Below is a list of mentioned variables and their meaning.

| Variable name             | Example value      | Description                                   |
| ------------------------- | ------------------ | --------------------------------------------- |
| ``NETWORK_DOMAIN`` [^1]   | ``local``          | Domain used by machines in local network.     |
| ``MACHINE_DOMAIN``        | ``server1.local``  | FQDN of machine running the project.          |
| ``COMPOSE_DOMAIN``        | ``services.tld``   | FQDN under which subdomain will be created.   |
| ``SCALE_DOMAIN`` [^2]     | ``server1.remote`` | Machine's domain in Tailscale network.        |
| ``DOCKER_BRIDGE_GATEWAY`` | ``127.16.0.1``     | IP of Docker bridge gateway.                  |
| ``CERT_RESOLVER`` [^3]    | ``letsencrypt``    | Default certificate resolver in Traefik.      |
| ``AUTH_MIDDLEWARE`` [^3]  | ``authentik@file`` | Default authentication middleware in Traefik. |
| ``DEFAULT_PROXY`` [^2]    | ``traefik``        | Default reverse proxy software.               |
| ``DOCKER_MANAGER`` [^2]   | ``portainer``      | Default Docker Compose manager software.      |

[^1]: Currently obsolete, will be removed in the future.
[^2]: Introduced as an experiment, not used at the moment.
[^3]: Will become obsolete after introducing improved Traefik configuration.
