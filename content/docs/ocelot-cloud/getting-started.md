---
title: "Getting Started"
weight: 5
---

To run Ocelot-Cloud, you need to have Docker installed on your system. For installation instructions, see the {{< external_link "https://docs.docker.com/get-docker/" "Docker Installation Guide" >}}. Installation of Ocelot-Cloud is done with the single command below. You may want to change the **INITIAL_ADMIN_NAME**, and if this is a production deployment, you should also use a more secure **INITIAL_ADMIN_PASSWORD**.

```bash
docker run \
  -d \
  --name ocelotcloud \
  -p 80:8080 \
  -p 443:8443 \
  -e INITIAL_ADMIN_NAME=admin \
  -e INITIAL_ADMIN_PASSWORD=password \
  -v /tmp:/tmp \
  -v /var/run/docker.sock:/var/run/docker.sock \
  --restart unless-stopped \
  ocelotcloud/ocelotcloud:beta
```

The setup may take a few seconds or minutes to finish. Below are some options for continuing.

### Option 1: Local Setup

This setup is intended for testing and exploring Ocelot-Cloud on your local PC.

* Visit [http://localhost](http://localhost) in your browser.
* Log in
* Settings Page > set **Host** to **localhost** and save
* App Store Page > download an app
* Installed Apps Page > start and open app

Feel free to explore the other features from here.

### Option 2: Casual LAN Setup

This is for a setup in a private, trusted LAN network to perform minimal work to get Ocelot-Cloud working.

* Visit **https://[device-lan-ip-address]**. 
  * There is a self-signed certificate, so your browser will throw a warning that this certificate is not safe. Assuming there is no malicious hacker in your network, you can simply ignore this. Click on **Advanced**. Depending on the browser, you may see options like **Proceed to [device-lan-ip-address] (unsafe)** or **Accept the Risk and Continue**, which must be clicked.
* Log in
* Settings Page > set **Host** to e.g. **my-domain.local** and save
  * Since we are not working with real certificates you can choose any domain you want.

Now we need a mechanism that allows your computer to resolve the domain name to the IP address of the Ocelot-Cloud server. The easiest way to start is to simply add the following entry to your hosts file: Every system has such a file. For example, on Linux, it is located at **/etc/hosts** and you could add entries like this:

```
<ocelot-cloud-lan-ip> my-domain.local
<ocelot-cloud-lan-ip> gitea.my-domain.local
<ocelot-cloud-lan-ip> mattermost.my-domain.local
(etc...)
```

On Windows, it is located at **C:\Windows\System32\drivers\etc\hosts**.

The disadvantage of the hosts file approach is that you have to enter such entries on every device, and you need to add all subdomains, which creates extra work when adding new apps. A more complex, yet scalable, solution is to set up a local DNS server. In any case, you can now download, start, and open an app.

### Option 3: Production Setup

This setup is intended for production use and assumes that you own a domain and have a DNS server that can resolve domains to the IP address of the Ocelot-Cloud server. The simplest setup is to host Ocelot-Cloud on a public server and purchase DNS services from a provider to create DNS records. However, it is also possible to host Ocelot-Cloud in a corporate LAN and use an internal DNS server.

* Create two A DNS records pointing to the IP address of the Ocelot-Cloud server for your domain. For example the 
  * base domain **production-domain.com** and the
  * wildcard domain **\*.production-domain.com**
* Visit [https://production-domain.com](https://production-domain.com)
  * As explained above, ignore the certificate warning and continue.
* Log in
* Settings Page > set **Host** to **production-domain.com** and save
* Now you can either 
  * upload your own wildcard certificate covering the domains above
  * or generate a wildcard certificate

### Additions

When testing Ocelot-Cloud, you should also enable remote backups and consider applying automatic updates. By default, Ocelot-Cloud automatically updates its apps, but not itself. Therefore, we recommend enabling daily automatic updates of Ocelot-Cloud by running Watchtower alongside it.

```bash
docker run -d \
  --name watchtower \
  -v /var/run/docker.sock:/var/run/docker.sock \
  --restart unless-stopped \
  containrrr/watchtower:1.7.1 ocelotcloud
```
