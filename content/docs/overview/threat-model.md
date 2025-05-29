---
title: "Threat Model"
weight: 35
type: "docs"
---

A threat model outlines the security assumptions, potential risks, and defense strategies of a system. This document describes how Ocelot-Cloud mitigates threats to protect its digital infrastructure.

## Security Measures

#### Cryptography
We only use modern, well-tested cryptographic algorithms, and administrators can supply their own wildcard certificates to secure any HTTPS endpoint.

#### Container Isolation
Every app runs in its own Docker container with no shared volumes, networks, or Docker socket access, and with restricted Linux capabilities.

#### Defensive Coding
* Strict input validation blocks SQL injection, XSS, RCE, and similar attacks.
* Protection against CSRF attacks is built into the security framework.
* App releases are zip files. When an app is uploaded to the App Store or downloaded by Ocelot-Cloud, it is checked for zip bombs and its consistency is validated to ensure that apps are always executed inside a sandbox.

#### Secure Supply Chain
* Static code analysis tools run frequently.
* Automated update tools ensure releases use the latest libraries and Docker images.

#### Access Control
* All components enforce role-based authentication and authorization for data.
* Only authenticated users in the local Ocelot-Cloud instance can access apps.

#### Backups and Updates 
* Ocelot-Cloud automatically updates apps and creates backups.
* Ocelot allows backups to be stored on remote servers using end-to-end encryption. This enables disaster recovery in case of hardware failure or data loss, but must be set up by the operator.

## Limitations and Assumptions

The following areas remain under the operator’s control and are out of scope for Ocelot-Cloud’s built-in safeguards.

#### Infrastructure

* When an operator hosts Ocelot-Cloud on his own hardware, he is responsible for the physical security of the servers and networks.
* DNS configuration is managed by the operator; Ocelot-Cloud does not operate its own DNS server.
* Monitoring and intrusion detection are not yet included. Operators are responsible for implementing their own solutions, if needed.

#### OS

* Host OS and Docker Engine updates must be maintained by the operator to ensure security. We recommend using a simple, automatically updating system like Ubuntu, along with a cron job to reboot regularly.
* OS-level network security such as firewall configuration must be enforced by the operator to protect the deployment.

#### Apps and Data

* App Store content is not moderated, except for official apps. Third-party apps may contain malicious code, so we recommend that operators only use trusted third-party sources or evaluate third-party apps before installation.
* App updates depend on each app's maintainer. Outdated apps may introduce vulnerabilities.
* A potential risk is the compromise of the App Store, which would allow malicious app versions to be distributed via automatic updates. To mitigate this risk, we plan to implement signed app releases by maintainers.

### Miscellaneous

* The owner of the Ocelot-Cloud installation is responsible for legal and regulatory compliance, such as privacy laws.
* Ocelot-Cloud comes with a default self-signed wildcard certificate. Although this might be suitable for trusted LAN environments, it is vulnerable to man-in-the-middle (MITM) attacks. To ensure security, operators must upload or generate a trusted wildcard certificate and keep it up to date.