---
title: "Gitea"
---

For production Gitea installations, the correct HOST value must first be set on the Settings page. This value defines internal Gitea parameters that cannot be changed later via the Gitea web interface. Incorrect values may result in operational issues.

If your Gitea instance was installed with the incorrect HOST value, there are two options:

* Update the HOST value and reinstall the app.
* Manually fix the configuration inside the running container.

To edit the configuration manually:

```bash
docker exec -it ocelotcloud_gitea_gitea bash
apk add nano
nano /data/gitea/conf/app.ini
```

Update the following values:

```text
DOMAIN = gitea.my-domain.com
SSH_DOMAIN = gitea.my-domain.com
ROOT_URL = https://gitea.my-domain.com/
```

Save the file, exit the container, and restart the app.