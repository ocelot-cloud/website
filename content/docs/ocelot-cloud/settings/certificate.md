---
title: "Certificate"
weight: 20
---

There are multiple ways to operate certificates in Ocelot-Cloud.

### 1) Default Certificate

By default, Ocelot-Cloud is available on port 80 via HTTP and on port 443 via HTTPS using a universal self-signed wildcard certificate. This default certificate works for all domains or IP addresses through which you access Ocelot-Cloud and any host configured in the Settings. This option may be sufficient for test deployments or private LAN networks. However, for most real-world use cases, a valid, properly signed wildcard certificate is recommended.

### 2) Generate Wildcard Certificate

This option generates a certificate via LetsEncrypt DNS-01 challenge. You must set **host** to a domain that you own for the setup to work, for example `sample.com`. The email address is optional, but if set, it will be used for Let's Encrypt notifications to let you know when your certificate is about to expire. Simply start the challenge, create the TXT DNS record as instructed, and wait a few minutes. Ocelot-Cloud will handle the rest for you. If you would like to, you can monitor the progress in the container logs using the following command:

```bash
docker logs -f ocelotcloud
```

In the background, Ocelot-Cloud creates a `fullchain.pem` file that contains keys and certificates for the domain `sample.com` and all its subdomains `*.sample.com`, such as `gitea.sample.com`, `mattermost.sample.com`, etc. The wildcard certificate is valid for 90 days, so this process must be repeated manually from time to time.

### 3) Upload Wildcard Certificates

Another option is to upload your own wildcard certificates. The uploaded **fullchain.pem** file must contain at least the key and certificate key for **\*.sample.com**.