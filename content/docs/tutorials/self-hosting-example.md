---
title: "Self-Hosting Example"
---

This sample guide walks you through setting up a hardened Ubuntu Server host for Ocelot-Cloud. It is intended for users with more technical experience. Adapt it to your own environment as needed.

### Restrict Physical Access

Place the device in a locked rack or drawer and disable unused USB ports in BIOS/UEFI.

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
sudo apt install -y ufw fail2ban unattended-upgrades
sudo systemctl enable --now fail2ban unattended-upgrades

# Firewall: only allow SSH + HTTPS (+ HTTP if needed)
sudo ufw default deny incoming
sudo ufw default allow outgoing
sudo ufw allow 22/tcp
sudo ufw allow 80/tcp # optional, if you want to use HTTP
sudo ufw allow 443/tcp
sudo ufw enable
```

Add your public key to the file `~/.ssh/authorized_keys`, then try logging in via SSH. If it works, disable password authentication.
```bash
sudo sed -Ei 's/^#?PasswordAuthentication .*/PasswordAuthentication no/' /etc/ssh/sshd_config
sudo systemctl reload sshd
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

Reboot workflow 
* reboot the server
* on user PC: 
  * Run the remote‑unlock script.
  * SSH in normally once the host comes back.

### Server Reboots

To ensure that the latest kernel and security updates are applied, you should reboot the server periodically, for example once a month. However, if you have FDE, you need to unlock the disk during boot, which requires you to unlock the encryption manually. You should be aware of this additional effort when choosing FDE.

### Install Docker

This prepares the installation of Ocelot-Cloud. You can use either the Snap package:

```bash
sudo snap install docker
sudo systemctl enable --now docker
```

or the Docker.io package from the official Ubuntu repositories. 

```bash
sudo apt install -y docker.io
sudo systemctl enable --now docker
```