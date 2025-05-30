---
title: "Gitea"
---

A general rule in Gitea is that many settings cannot be changed via the web GUI once the installation wizard is completed. However, you can still modify them by directly editing the `app.ini` file inside the container as described below. As this is a somewhat hacky solution, we recommend deciding whether to allow self-registration before deploying a production instance.

### Installation Wizard

When visiting the Gitea web UI for the first time, do the following:

- Disable self-registration if you want to prevent strangers from registering: Server and Third-Party Service Settings → uncheck 'Disable Self-Registration'
- Admin Account: Administrator Account Settings → fill in the required fields

### Create Users

If self-registration is disabled, you can create users like this:
- Log in as Admin → Profile and Settings (top right) → Site Administration → Identity & Access → User Accounts

### Adapting Configuration After Installation

Open the `app.ini` file in the Gitea container:

```bash
docker exec -it ocelotcloud_gitea_gitea bash
apk add nano
nano /data/gitea/conf/app.ini
```

Adapt the parameters as needed. For example, adjust the self-registration policy by setting `DISABLE_REGISTRATION` to `true` or `false`.

Save the `app.ini` file and restart the app.