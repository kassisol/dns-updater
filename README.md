# DNS Updater

DNS Updater runs as a Docker service on a Docker Swarm manager host and listens for service create events then add cname to DNS provider based on label provided.

Only one dns backend is supported which is `infoblox`.

## Getting started

### Run it

First add the config file in Docker Secret:

```bash
# docker secret create infoblox_config -
---
ipam_host: '<hostname>'
api_version: '2.4'
username: '<user>'
password: '<password>'
```

Then start dns-updater service on any of the Docker Swarm manager using that secret

```bash
# docker service create -d --constraint "node.role==manager" -v /var/run/docker.sock:/var/run/docker.sock --secret source=infoblox_config,target=/config.yml --name dns-updater kasissol/dns-updater:x.x.x -c /config.yml
```

### Use it

Labels are used to configure your services.

| Label     | Description                                  |
|-----------|----------------------------------------------|
| dns.host  | Give host name the CNAME will be pointed to. |
| dns.cname | Give the CNAME name.                         |

Example:

```bash
# docker service create -d --label "dns.host=foo.example.com" --label "dns.cname=bar.example.com" --name web1 nginx
```

## User Feedback

### Issues

If you have any problems with or questions about this application, please contact us through a [GitHub](https://github.com/kassisol/dns-updater/issues) issue.
