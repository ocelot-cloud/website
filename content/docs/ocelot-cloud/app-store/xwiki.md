---
title: "XWiki"
---

When you first access the XWiki web UI, follow the installation wizard to create the admin account.

### Disable Self-Registration

To prevent users from self-registering:

* Drawer menu (top left) → Administration → Users & Rights → Rights → Users
  * In the 'XWikiAdminGroup' row, click the 'Register' column until a green checkmark appears (explicit allow)

### Make Wiki Private

To restrict access to authenticated users only:

- Drawer menu (top left) → Administration → Users & Rights → Rights → Users
- Enable:
  - Prevent unregistered users from viewing pages, regardless of the page rights
  - Prevent unregistered users from editing pages, regardless of the page rights

### Create Users

With self-registration disabled, administrators can create users manually:

- Drawer menu (top left) → Administration → Users & Rights → Users → Add User
