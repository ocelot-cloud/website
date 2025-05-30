---
title: "Vaultwarden"
---

For production setups, after installation, it is crucial to log in as admin and change the admin security token.

### Setup Steps

- Install and start Vaultwarden
- Visit admin panel: `https://vaultwarden.<your-host>/admin`
- Enter admin token: `ocelot`
- Go to General Settings → Admin token / Argon2 PHC → set a secure password → Save

### Manual User Creation

- Admin: Admin panel → Users → enter the users email → Invite
- User: visit Vaultwarden home page → Create Account → enter the same user email and set a password → Create Account

### Self-registration (Optional)

To enable or disable self-registration: Admin panel → General Settings → toggle 'Allow new signups'
