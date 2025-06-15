---
title: "Self-Hosting Example"
---

This sample guide walks you through setting up a hardened Ubuntu Server host for Ocelot-Cloud. It is intended for users with more technical experience. Adapt it to your own environment as needed.

### Restrict Physical Access

Place the device in a locked rack or drawer.

### OS Installation

* Download Ubuntu Server: Grab the latest ISO from {{<external_link "https://ubuntu.com/download/server" "Canonical’s download page" >}}.
* Flash it to a USB drive (Rufus, balenaEtcher, etc.). You will need to connect a screen and a keyboard in order to complete the initial setup.
* During setup:
  * Pick minimal installation, no GUI.
  * Enable SSH server.
  * Optional: Choose LVM with full-disk encryption (FDE) if you want encrypted storage (see below).

### Internet Connection

Connect the server to a wired LAN to avoid having to configure a Wi-Fi connection on boot. With FDE, Wi-Fi setup would even be required twice: once in the encrypted-mode initramfs and again in the decrypted-mode main operating system. Please note that the LAN IP address of the server may differ between initramfs mode and the main operating system, since the router may recognise them as two distinct devices.

### Basic hardening

After the installation, log in via SSH and run the following commands to harden your server:

```bash
sudo apt update
sudo apt install -y ufw fail2ban unattended-upgrades nano
sudo systemctl enable --now fail2ban unattended-upgrades

# Firewall: only allow SSH + HTTPS (+ HTTP if needed)
sudo ufw default deny incoming
sudo ufw default allow outgoing
sudo ufw allow 22/tcp
# sudo ufw allow 80/tcp # optional, if you want to use HTTP, but not recommended
sudo ufw allow 443/tcp
sudo ufw enable
```

##### Additional Open Ports

Please note that Docker may override these firewall rules. For example, when you install an app, it may require Docker to open additional ports in order to function properly.

##### SSH Hardening

Add your public key to the file `~/.ssh/authorized_keys`, then try logging in via SSH. If it works, disable password authentication.

```bash
sudo sed -Ei 's/^#?PasswordAuthentication .*/PasswordAuthentication no/' /etc/ssh/sshd_config
sudo systemctl reload sshd
```

### Update Automation

To automate system updates, create a script and schedule it with cron.

##### Create the update script

```bash
nano /usr/local/sbin/update.sh
```

Insert the following content:

```bash
#!/bin/bash

set -e

export DEBIAN_FRONTEND=noninteractive
echo "Starting update job at $(TZ=Europe/Berlin date +%Y-%m-%d_%H:%M:%S)"
echo "Pruning idle docker images"
docker image prune -af
echo "Updating"
apt-get update
echo "Upgrading"
apt-get upgrade -y
echo "Removing no longer required packages"
apt-get autoremove --purge
echo "Rebooting"
shutdown -r now
echo ""
```

Make the script executable:

```bash
chmod 700 /usr/local/sbin/update.sh
```

##### Schedule the script with cron

Edit the root user's crontab:

```
crontab -e 
```

Add the following line to run the script daily at 4:00 AM and log output:

```
0 4 * * * /bin/bash /usr/local/sbin/update.sh >> /var/log/updates.log 2>&1
```


### Full Disk Encryption

FDE keeps your data unreadable for third parties even if the server is physically stolen.

To continue, you will need a private/public key pair on your work PC. Run the commands below on the server to install your public key in its initramfs. This will then allow you to access the server via SSH during boot and unlock the disk.

```bash
sudo apt install -y dropbear-initramfs
read -rp "Paste your SSH public key: " k; echo "$k" | sudo tee -a /etc/dropbear/initramfs/authorized_keys
sudo update-initramfs -u
```

Use the `remote-unlock.sh` script below on your work PC to unlock the server disk. This script is designed for Linux, e.g. Ubuntu Desktop. It can also be used on Windows via WSL or Git Bash.

```bash
#!/usr/bin/env bash
read -rsp "LUKS passphrase: " p; echo
printf %s "$p" | ssh -i ~/.ssh/id_ed25519 root@SERVER-IP cryptroot-unlock
```

##### Reboot Workflow

* reboot the server
* on user PC: 
  * Run the remote‑unlock script.
  * SSH in normally once the host comes back.

### Server Reboots

To ensure that the latest kernel and security updates are applied, you should reboot the server periodically, for example once a month. Reboots can be automated using a cron job. However, if you have FDE, you need to unlock the disk during boot, which requires you to unlock the encryption manually. You should be aware of this additional effort when choosing FDE.

### Install Docker

Follow the official {{<external_link "https://docs.docker.com/engine/install/ubuntu/" "Docker installation guide" >}}. This Docker installation prepares the installation of Ocelot-Cloud.