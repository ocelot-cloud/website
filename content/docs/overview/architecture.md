---
title: "Architecture"
weight: 25
---

Ocelot-Cloud is a monolithic application that runs on a single server. Each HTTP request from users is either processed by Ocelot-Cloud directly or passes through Ocelot-Cloud and is proxied to an app. The apps are "hidden" behind Ocelot-Cloud. This way, Ocelot-Cloud serves as an extra layer of security, as only authenticated users have access to the apps by default.

<img src="/images/ocelot-architecture.drawio.svg" style="width:60%;" alt="Architecture Diagram">

For security reasons, each app runs in its own isolated space with its unique Docker network and volumes. So apps are unaware of each other.