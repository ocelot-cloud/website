---
title: "FAQ"
weight: 40
---

{{< details "What is Ocelot-Cloud?" >}}
Ocelot-Cloud is a infrastructure management platform that allows you to set up and maintain your own secure digital infrastructure. It automates tasks like installing apps, backups, and updates, making self-hosting accessible without requiring extensive technical expertise.
{{< /details >}}

{{< details "Who is Ocelot-Cloud for?" >}}
Ocelot-Cloud is designed for individuals and small to medium-sized organizations that want control over their digital infrastructure without relying on external cloud providers. It's ideal for those who value privacy, customization, and long-term sustainability.
{{< /details >}}

{{< details "How difficult is it to install Ocelot-Cloud?" >}}
Ocelot-Cloud can be installed on any device that supports a widely used technology called 'Docker', making it largely hardware and operating system independent. Installation requires only a single Docker command on the command line. All server maintenance tasks are done automatically by Ocelot-Cloud in the background.
{{< /details >}}

{{< details "What does self-hosted mean?" >}}
Self-hosted means you run the software on your own hardware, giving you complete control over your data and infrastructure. This is in contrast to external hosting, where your data and apps are managed by a third party. This is usually more convenient because hardware maintenance is outsourced, but it can expose you to risks such as unauthorized data access and vendor lock-in.
{{< /details >}}

{{< details "Is Ocelot-Cloud secure?" >}}
Practically, yes. Security is a top priority for Ocelot-Cloud, and we've built in extensive protections, see "Threat Model" article. While we prioritize minimizing risk through a security-by-design approach, it's technically impossible to achieve 100% security, and no system is completely immune. We remain committed to ongoing efforts to keep our platform as secure as possible.
{{< /details >}}

{{< details "Will Ocelot-Cloud spy on my data?" >}}
No. The Community Edition of Ocelot-Cloud is fully open-source. Every line of code is publicly available, so it can easily be audited to detect any malicious changes. You can also run it on your own hardware, giving you full control over your data. While any closed-source software, including our Premium Edition, could theoretically hide harder to detect spyware, doing so would contradict our core values and destroy the trust we've built with our community.
{{< /details >}}

{{< details "How does Ocelot-Cloud compare to Nextcloud?" >}}
Nextcloud is an app focused on file storage and collaboration. Ocelot-Cloud, on the other hand, is a generic platform that allows you to install and manage multiple apps, including Nextcloud, but also many others. Ocelot-Cloud also handles administrative tasks such as updates, backups and more, making it much more versatile and powerful than a standalone Nextcloud.
{{< /details >}}

{{< details "How does Ocelot-Cloud compare to traditional cloud providers?" >}}
By hosting your own cloud with Ocelot-Cloud, you eliminate the need to trust external providers with your data. This reduces risks of unauthorized access, surveillance, and vendor lock-in. You have full control over your data and can customize your infrastructure to meet your specific needs without worrying about changing terms or excessive fees.
{{< /details >}}

{{< details "How does Ocelot-Cloud compare to Proton?" >}}
{{< external_link "https://proton.me/" "Proton" >}} offers services that allow users to upload data to its servers in an end-to-end encrypted form, ensuring that only the user, not even Proton, can read the data. Proton provides free client software that transparently manages encryption and decryption on the user's device, as well as uploading and downloading data to the external Proton servers. Although Proton offers better privacy than traditional cloud providers, it has certain limitations compared to self-hosting:
* **Metadata Collection**: Proton can still collect metadata about user activity, such as timestamps, IP addresses, and interaction patterns, which can provide insight into usage even when content remains private.
* **Legal Compliance**: Proton is subject to the laws of the country in which it operates (Switzerland). If a government requests collected data or imposes certain regulations, such as monitoring your activity, Proton must comply, even if they don't want to, which could compromise user privacy.
* **Internet Exposure**: Proton's servers are accessible via the Internet, creating potential attack vectors. In contrast, self-hosted infrastructure can be installed in a private network, reducing exposure to Internet-based threats.
* **Limited Features and Customization**: Because Proton manages its own servers, users cannot customize the infrastructure or integrate additional apps. Self-hosting provides full control over software, configurations, and custom integrations.
* **Vendor Lock-In**: Users are subject to Proton's pricing and policies, as the company maintains full control over its servers. A potentially costly data migration may be required if their terms become unfavorable.

Ocelot-Cloud can be used as a SaaS. You can install end-to-end encrypted apps like {{< external_link "https://www.vaultwarden.net/" "Vaultwarden" >}} or {{< external_link "https://cryptpad.org/" "Cryptpad" >}} bringing a similar mix of convenience and privacy as Proton. But with Ocelot-Cloud you have the option to switch to self-hosting with ease at any time.
{{< /details >}}

{{< details "How do I migrate from self-hosted to SaaS or vice versa?" >}}
Migration with Ocelot-Cloud is straightforward. The key feature is the creation of remote backups on external servers. In order to migrate e.g., from SaaS to self-hosted, you simply create backups of all apps in a remote repository. Then you set up your own server, install Ocelot-Cloud, configure access to the remote repository and restore the backups from the remote backup server.
{{< /details >}}

{{< details "What hardware do I need to run Ocelot-Cloud?" >}}
The processor must be based on the AMD64 architecture. Ocelot-Cloud is designed to be resource efficient, so it can run on low-end hardware. The actual hardware requirements depend on the apps you install. To determine the right hardware, check the CPU, memory, and disk space requirements of each app. If you're unsure, choose a system with more resources. For environments with limited hardware, you can explicitly choose resource-efficient apps.
{{< /details >}}
