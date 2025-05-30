---
title: "App Store"
weight: 20
no_list: true
---

This article explains the **App Store** page in the Ocelot-Cloud GUI. The actual {{< external_link "https://store.ocelot-cloud.org/" "App Store" >}} is a separate online service where anyone can upload their own apps.

The App Store lets you search and install apps on your Ocelot-Cloud instance with just a few clicks. By default, the latest version is suggested for installation, but you have the option to install an older version.

### Official Apps

When installing an official app, you should always read the according articles below as they contain setup instructions , security guidelines , information about self-registration

{{< children >}}

By default, the self-registration feature is disabled for all official apps for security reasons. However, if you intend to found an open community which allows self-registration, the articles provides instructions on how to enable it.

### Unofficial Apps

By default, only official apps provided by us are shown when searching for apps. This means that your search results will only contain apps of use to the official maintainer with the name "ocelotcloud". By enabling "Show Unofficial Apps", you have the option to search for unofficial apps from third-party maintainers. The search term will then match maintainers or app names.

{{< alert title="Note" color="warning" >}}
Unofficial apps are not moderated by us and may contain malware or other harmful content. We do not recommend installing them unless you trust the maintainer. Installing apps from unknown sources may compromise your server and data.
{{< /alert >}}

### Migrating to Another Maintainer

If an app is no longer maintained and not receiving updates, you can migrate to a new maintainer:

* A new maintainer in the App Store must create an app of the same name and version, using identical Docker volumes and paths in the `docker-compose.yml`.
* Ocelot-Cloud users must rename Docker volumes on their local server, e.g., `oldmaintainer_gitea_data` to `newmaintainer_gitea_data`.
* After renaming, users can download and start the according version of the app.