---
title: "Installed Apps"
weight: 10
---

On the Installed Apps page, you can manage your apps. You can perform the following operations on an app:

* Start
* Stop
* Backup
* Update, which will create a backup and then install the latest version of the app.
* Delete, which removes the app and its artifacts; only its backups are preserved.

### Public Apps

By default, all apps are private, meaning only authenticated Ocelot-Cloud users can access them. However, you can make an app public, allowing anyone to access it without authentication.

Typical use cases for public apps include:

* Public content without self-registration: Hosting public resources, such as blogs or wikis, whose content is created by users who have been invited by an admin.
* Open communities with self-registration: For servers where anyone can register and participate without needing admin approval. 
* Client integration: Some desktop or mobile clients cannot connect to private Ocelot-Cloud apps since they cannot pass Ocelot-Cloud’s authentication layer. Making the app public removes this barrier.

### Miscellaneous

* It is not possible to start two apps with the same name.
* Apps have a state which can be either **Available** or **Uninitialized**. At the moment, the implementation is quite simple, and an app that has been started immediately becomes **Available**. However, in the background, each app can take a moment, sometimes even minutes, to download and launch the containers. So it can take a while before the app is actually available. We will improve this in the future.

### Special Case: ocelotcloud / ocelotdb

Ocelot-Cloud uses a postgres database called **ocelotdb** to store all its data. It is technically efficient to treat the **ocelotdb** as an app, since it works practically the same way. This is why **ocelotcloud / ocelotdb** is listed in the Installed Apps page by default. Since the database must always be running and available, your operations on the database app are limited to creating backups. Conversely, restoring backups on the Backups page means restoring Ocelot-Cloud data, such as users, settings, installed apps meta data, etc.