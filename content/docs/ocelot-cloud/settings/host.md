---
title: "Host"
weight: 10
---

The host value is needed to access installed apps. For example, if you set the host to **company.com** and click Open on the gitea app on the **Installed Apps** page, **gitea.company.com** will open in a new tab. For local testing, you can simply use **localhost** as the host value.

Note: This proxying requires that the domain (e.g. gitea.company.com) resolves to the server where Ocelot-Cloud is installed, either via a DNS record or an entry in your system's hosts file.

Some apps use the HOST value for security or compatibility reasons. Changing the HOST value may cause those apps to stop working properly. You can resolve this issue by simply restarting these apps, which will make Ocelot-Cloud pass the new HOST value to the app.
