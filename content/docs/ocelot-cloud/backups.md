---
title: "Backups"
weight: 40
---

Ocelot-Cloud can create backups of each app on the Installed Apps page. On the "Backups" page you can view and restore backups of apps.

{{< alert title="Note" color="warning" >}}
Restoring an app backup will erase any previously existing data that was not backed up. It is recommended that you create a current backup before restoring an old backup.
{{< /alert >}}

By default, only "Local" backups are created on the same device as Ocelot-Cloud and can be viewed here, so you can only select "Local" as backend repository type. On the settings page you can enable remote backups, which extends the selectable backend repository type by "Remote", which displays backups created on the configured remote repository.

Note that restoring **ocelotcloud / ocelotdb** only restores the metadata of the apps, not the persistent storage of the apps. This means, for example, that if you want to do a full restore to a new device using a remote backup repository, you must first restore the backup of the **ocelotcloud / ocelotdb** app, and then restore the backups of each app that you need.