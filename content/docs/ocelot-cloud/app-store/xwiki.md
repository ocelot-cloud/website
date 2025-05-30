---
title: "XWiki"
---

The installation process is straightforward with no additional considerations.

TODO: Disable: in Administration → Users & Rights → Users, remove the Register right from Unregistered Users (or grant it only to XWikiAdminGroup)

TODO explain how to enable self-registration for an open community:
* Gitea — Site Administration → Configuration → Service → Disable self-registration toggle.
* Mattermost — System Console → Authentication → Signup → Enable account creation / Enable open server switches.
* XWiki — Administration → Users & Rights → remove the Register right from Unregistered Users.
* OCIS — No public signup feature; only the built-in admin can add users via the Users page or CLI.
* Vaultwarden — /admin panel → uncheck Allow signups.