---
title: "Licensing Model"
---

Our licensing strategy balances openness with a sustainable business model.

### 1. 0BSD Components

Parts of the codebase that are purely infrastructural are released under the open source 0BSD license, a public-domain-equivalent license that imposes no attribution or redistribution requirements. You may use, copy, modify, and redistribute them without restriction. The 0BSD license is recognized and accepted by major organizations. For instance, Google permits its employees to contribute to and use 0BSD-licensed projects internally.

### 2. Open-Core Model

Ocelot-Cloud is offered in two editions:

| Edition | Price | License     | Purpose                          |
|---------|-------|-------------|----------------------------------|
| Community Edition (CE) | Free | AGPL-3.0    | Fully open-source core           |
| Premium Edition (PE)   | Paid | Proprietary | Advanced, non-essential features |

The PE generates revenue to fund development.

#### AGPL + CLA (Temporary)

The AGPL-3.0 copyleft license and our Contributor License Agreement (CLA) are primarily safeguards against proprietary capture, situations where a well-funded company could fork the project, close the source, and commercialize it without contributing back.
* AGPL-3.0 forces any publicly deployed derivative to remain open-source.
* The CLA assigns contribution ownership to us, allowing us, not outside actors, to dual-license the same code for the PE to fund our work.

Once we have developed the PE to a point where the risk of proprietary capture is minimized, we plan to remove that protection and relicense the CE under 0BSD, so that developers no longer have to read or sign legal documents in order to contribute. Then, under the terms of the 0BSD license, all prior claims of ownership will be extinguished as if the CLA never existed.

#### What this means for *users*

- Running CE unmodified has no obligations, use it as you wish.

#### What this means for *developers*

- When modifying CE source code:
  - Your code modifications are automatically licensed under AGPL-3.0.
  - You must make the modified source code available to any user who has access to your modified version.
  - You may, but are not required to, contribute the modified code to our GitHub repository via pull request for others to benefit.

When contributing to Ocelot-Cloud:

1. Fork the repo and implement your changes.
2. Sign the `legal/cla.md` once by copying the `legal/cla-signing-template.md` to `legal/contributors/<your-github-name>.md` and fill it out.
3. Open a pull request.

Your signature covers all future contributions.
