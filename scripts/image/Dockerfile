FROM scratch

ARG version

LABEL maintainer="hbm@kassisol.com"
LABEL version=$version
LABEL description="DNS Updater runs as a Docker service on a Docker Swarm manager host and listen for service create events then add cname to DNS provider based on label provided"

COPY build/dns-updater /dns-updater

ENTRYPOINT ["/dns-updater"]
