---
title: "Maintenance"
weight: 30
---

The features **Automatic Updates** and **Automatic Backups** are enabled by default and are handled by the Maintenance Agent in the background. The Maintenance Agent runs once a day.

If both options are enabled, the Maintenance Agent compares the version of each locally installed app with the latest version in the App Store:
* If no update is available, a backup is created.
* If an update is available, a backup is created, then the update is applied. If the update fails, you can restore the previous version.

In **Preferred time for maintenance (UTC)** you can set the hour at which the update should be performed. It is recommended to select a time in which the Ocelot-Cloud is unlikely to be used, like 4:00 in your respective time zone. The time for this option is set in UTC (London time).

If remote backups are enabled on the **Settings** page, the Maintenance Agent will also create backups on the remote backup server.

### Retention Policy

To free up storage space, the Maintenance Agent applies a retention policy. It keeps only the:
* last 7 daily backups
* last 4 weekly backups
* last 12 monthly backups

This means that backups are kept for up to a year, and older and intermediate backups are automatically deleted. On the **Backups** page, you can still manually delete all backups, including automatically created ones. There, you can distinguish the backup type by the backup descriptions, which are either "auto backup" or "manual backup". Manually created backups are not deleted by the retention policy and are kept until you manually delete them.

Backups use disk space efficiently because subsequent backups store only the differences from the previous backup.