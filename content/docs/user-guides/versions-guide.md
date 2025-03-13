---
title: "App Store Version Creation Guide"
tags: ["store", "app"]
---

This article explains how to create new apps and versions in the <a href="https://store.ocelot-cloud.org/" target="_blank" rel="noopener noreferrer">App Store</a>.

## What is a version?

Each app, like the "gitea" source repository app, can have multiple versions, just like any software itself has versions. An app version is a zip file containing a `docker-compose.yml` and an optional `app.yml` for integration with the Ocelot-Cloud. When you download and run an app in the Ocelot-Cloud, it is executing:

```sh
docker compose up -d
```

## Where to create my own versions?

Go to **"My Apps"**, create an app, select it, click **"Edit"**. This is the menu to upload a version for an app.

## Minimal Example

Here is a minimal example `docker-compose.yml` file suitable for upload:

```yaml
services:
  gitea:
    image: gitea/gitea:1.20.2
    container_name: samplemaintainer_gitea_gitea
```

The `docker-compose.yml` has specific conventions which are explained in detail below. There must be a "main service," in this case "gitea," which contains a web server to which Ocelot-Cloud can redirect HTTP requests. When a version is successfully uploaded, the server will automatically add some default settings in the background, as shown below:

```yaml
services:
  gitea:
    image: gitea/gitea:1.20.2
    container_name: samplemaintainer_gitea_gitea

    restart: unless-stopped
    networks:
      - gitea-net
    cap_drop:
      - ALL
    cap_add:
      - CAP_NET_BIND_SERVICE
      - CAP_CHOWN
      - CAP_FOWNER
      - CAP_SETGID
      - CAP_SETUID
      - CAP_DAC_OVERRIDE

networks:
  gitea-net:
    external: true
```

This means that restarting, networking, and reducing the default capabilities are automatically taken care of in the background, keeping the `docker-compose.yml` upload lean.

## Full Example

Here is an example of a `docker-compose.yml` that can be uploaded in a zip using multiple keywords:

```yaml
volumes:
  samplemaintainer_gitea_data:

services:
  gitea:
    image: gitea/gitea:1.20.2
    container_name: samplemaintainer_gitea_gitea
    environment:
      - USER_UID=1000
      - USER_GID=1000
      - DISABLE_REGISTRATION=true
    ports:
      - 2222:22
    volumes:
      - samplemaintainer_gitea_data:/data
    deploy:
      resources:
        limits:
          cpus: 0.5
          memory: 512M
        reservations:
          cpus: 0.2
          memory: 256M
```

## Configuration for Ocelot-Cloud

You may also need to define an `app.yml` that will automatically integrate the container correctly into Ocelot-Cloud. If missing, Ocelot-Cloud tries to redirect HTTP requests to port 80 to the path "/" of the main container "gitea". Here is an example:

```yaml
port: 3000
url_path: "/"
```

## Packaging and Upload

This `app.yml` causes HTTP requests with target `http(s)://gitea.'your-host-domain'` received by Ocelot-Cloud to be proxied to port 3000 of the `gitea` container, where the Gitea web server is listening.

Once you have prepared these two files according to all the conventions, zip them together and drag them to the upload area of the Version management page mentioned at the beginning. For zipping, you can use a graphical tool or a CLI command like this one:

```sh
zip 1.20.2.zip app.yml docker-compose.yml
```

**Tip:** Versions are usually named after the version number of the main service they are meant to deploy. It is recommended to use version names that correspond to the version of the main service, in this case "gitea", like this `1.20.2.zip`.

## Conventions and Restrictions

### The version zip file:

- Must contain a `docker-compose.yml`.
- May contain an optional `app.yml`.
- Must not contain other files.

### The `docker-compose.yml`:

- May only have the root keywords: `volumes`, `services`.
- May only use the following service keywords: `image`, `container_name`, `ports`, `volumes`, `depends_on`, `environment`, `deploy:resources`, `tmpfs`.
- May only contain services that use at least the keywords `image` and `container_name`.
- May only contain `volumes` and `container_name` with the prefix `[maintainer]_[appname]_*`. For example, `ocelotcloud_gitea_mariadb`.
- Must contain one main service by setting the value of its `container_name` to the following pattern `[maintainername]_[appname]_[appname]`. For example, `ocelotcloud_gitea_gitea`. This service is where Ocelot-Cloud expects a web server to be running so that it can proxy HTTP requests to the port and URL path defined in the `app.yml`.
- May only contain values for the `image` keyword with a tag of a deterministic version, e.g., `gitea/gitea:1.20.2` but not `gitea/gitea` or `gitea/gitea:latest`. This ensures reproducibility.

### The `app.yml`:

- May contain the keywords: `port` and `url_path`. If not set, they default to port 80 and `/`.
