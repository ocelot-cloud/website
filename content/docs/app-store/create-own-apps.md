---
title: "Creating Own Apps"
---

### App Store Requirements

Any software with a web interface can become an app. You simply need to have a publicly available Docker image containing the software and follow the steps below to create an app in the App Store. When the configuration is done right and the app is installed in Ocelot-Cloud, it will handle networking, certificate handling, backups and updates in the background.

### Apps and Versions

An app is a conceptual label for a software package, such as the software name "gitea" for a source code management software. The detailed configuration required to install an app is defined in its versions. Each version ships with a docker-compose.yml and optionally an app.yml for Ocelot-Cloud integration. When an Ocelot-Cloud instance downloads and installs a package from the store, it essentially runs:

```sh
docker compose up -d
```

### Where to create my own versions?

Go to **My Apps**, create an app, select it, click **Edit**. This is the menu to upload a version for an app.

### Minimal Example

Here is a minimal example `docker-compose.yml` file suitable for being packaged and uploaded:

```yaml
services:
  gitea:
    image: gitea/gitea:1.20.2
    container_name: samplemaintainer_gitea_gitea
```

The `docker-compose.yml` has specific conventions which are explained in detail below. There must be a "main service", in this case "gitea", which contains a web server to which Ocelot-Cloud can redirect HTTP requests. When a version in the app store is downloaded and started by an Ocelot-Cloud, it will automatically add some default settings in the background, as shown below:

```yaml
services:
  gitea:
    image: gitea/gitea:1.20.2
    container_name: samplemaintainer_gitea_gitea

    restart: unless-stopped
    networks:
      - samplemaintainer_gitea
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
  samplemaintainer_gitea:
    external: true
```

This means that restarting, networking, and reducing the default capabilities are automatically taken care of in the background, keeping the upload lean.

### Full Example

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
      - GITEA__server__DOMAIN=gitea.${HOST}
      - GITEA__server__ROOT_URL=https://gitea.${HOST}/
      - GITEA__server__SSH_DOMAIN=gitea.${HOST}
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

##### ${HOST}

Some apps require the domain they are accessible at for security or compatibility reasons. As this value varies, you can use `${HOST}` as a placeholder. When an app is downloaded and started, this placeholder is automatically replaced by the HOST value defined on the Settings page of the Ocelot-Cloud.

If you need a full URL with protocol, use HTTPS and not HTTP, since Ocelot-Cloud assumes that production environments use TLS.

### Configuration for Ocelot-Cloud

You may also need to define an `app.yml` that will automatically integrate the container correctly into Ocelot-Cloud. If missing, Ocelot-Cloud tries to redirect HTTP requests to port 80 to the path "/" of the main container "gitea". Here is an example:

```yaml
port: 3000
url_path: "/"
```

### Packaging and Upload

This `app.yml` causes HTTP requests with target `http(s)://gitea.'your-host-domain'` received by Ocelot-Cloud to be proxied to port 3000 of the `gitea` container, where the Gitea web server is listening.

Once you have prepared these two files according to all the conventions, zip them together and drag them to the upload area of the Version management page mentioned at the beginning. For zipping, you can use a graphical tool or a CLI command like this one:

```sh
zip 1.20.2.zip app.yml docker-compose.yml
```

**Convention:** Versions are named after the version number of the main service they are meant to deploy. It is recommended to use version names that correspond to the version of the main service, in this case "gitea", like this `1.20.2.zip`.

### Conventions and Restrictions

This section lists the rules to follow when uploading a version to the App Store. The conventions are chosen in such a way that you still have a fully functional `docker-compose.yml` which can be deployed on your local PC via `docker compose up -d`, making the development and debugging process of apps easier.

Don't be afraid to upload a zip as the server will automatically check the files for compliance. If a rule is broken, the upload will fail and the server will provide helpful feedback on what needs to be fixed. Successfully uploaded versions can be easily deleted if you made a mistake.

The version zip file:

- Must contain a `docker-compose.yml`.
- May contain an optional `app.yml`.
- Must not contain other files.

The `docker-compose.yml`:

- May only have the root keywords: `volumes`, `services`.
- May only use the following service keywords: `image`, `container_name`, `ports`, `volumes`, `depends_on`, `environment`, `deploy:resources`, `tmpfs`, `tty`, `user`, `command`.
- May only contain services that use at least the keywords `image` and `container_name`.
- May only contain `volumes` and `container_name` with the prefix `[maintainer]_[appname]_*`. For example, `ocelotcloud_gitea_mariadb`.
- Must contain one main service with the name of the app, e.g. `gitea`. The value of its `container_name` must follow the pattern `[maintainername]_[appname]_[appname]`. For example, the container name `ocelotcloud_gitea_gitea`. This service is where Ocelot-Cloud expects a web server to be running so that it can proxy HTTP requests to the port and URL path defined in the `app.yml`.
- May only contain values for the `image` keyword with a tag of a deterministic version, e.g., `gitea/gitea:1.20.2` but not `gitea/gitea` or `gitea/gitea:latest`. This ensures reproducibility.

The `app.yml`:

- May contain the keywords: `port` and `url_path`. If not set, they default to port 80 and `/`.

### App Design Recommendations

You have complete freedom in designing your app. The following recommendations are not mandatory, but can help improve the user experience and integration to Ocelot-Cloud. If you are developing an app or can influence its behavior, we recommend the following:

* All dynamic configuration should be done via the app’s web interface.
* On first startup, prompt the first user to create an admin account to manage the instance.
* Self-registration should be disabled by default, with user creation handled either manually or via integration with a central authentication system (e.g. OIDC). Self-registration can then be offered as an optional feature.
* Implement OIDC authentication to enable Ocelot-Cloud’s single sign-on (SSO). Ideally, configure OIDC via the Docker Compose environment variables, using the `${HOST}` placeholder.
* For installation, a simple `docker compose up -d` command should suffice, without any additional manual steps being required.
* For browser-based user interaction, there should only be one container running a web server and exposing one HTTP port. Additional containers can, of course, be used for app-internal tasks such as databases.

### Reserved Exposed Ports

Some official apps expose specific ports to the public via the docker-compose.yml file. For example:

```yaml
ports:
  - 2222:22
```

Please note that the HTTP port should never be exposed in this way, as it is easier and more secure to let Ocelot-Cloud proxy the traffic to this port internally.

If you provide a third-party app, avoid port conflicts. Use the same port only if it’s the same app. For different apps, choose a non-reserved port. 

| Port | App | Service |
| ------------- |--------------| ------- |
| 2222          | backup       | SSH     |
| 2223          | gitea        | SSH     |

Additional ports will be added in the future as new official apps are added to the App Store.