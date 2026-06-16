---
name: deception-technology-skill-selector
description: >-
  A deception technology skill-selection catalog for choosing the correct
  deception technology skill from the 26 official security domains. Use this
  domain catalog before opening a specific deception technology skill.
domain: cybersecurity
subdomain: deception-technology
tags:
  - deception-technology
  - skill-selection
  - cybersecurity-skills
  - domain-catalog
version: "1.0"
author: generated
license: Apache-2.0
---

# Deception Technology Skill Catalog

Use this file to choose a deception technology skill. The `When use` text is copied
from each skill's own `description` metadata.

Source subdomains included: `deception-technology`.

| No. | When use | Skill |
| ---: | --- | --- |
| 1 | Deploys deception-based honeytokens in Active Directory including fake privileged accounts with AdminCount=1, fake SPNs for Kerberoasting detection (honeyroasting), decoy GPOs with cpassword traps, and fake BloodHound paths. Monitors Windows Security Event IDs 4769, 4625, 4662, 5136 for honeytoken interaction. Use when implementing AD deception defenses for detecting lateral movement, credential theft, and reconnaissance. | [deploying-active-directory-honeytokens](skills/deploying-active-directory-honeytokens/SKILL.md) |
| 2 | Deploy and monitor Canary Tokens via the Thinkst Canary API for deception-based breach detection using web bug tokens, DNS tokens, document tokens, and AWS key tokens. | [implementing-deception-based-detection-with-canarytoken](skills/implementing-deception-based-detection-with-canarytoken/SKILL.md) |
| 3 | Deploy and manage network honeypots using OpenCanary, T-Pot, or Cowrie to detect unauthorized access, lateral movement, and attacker reconnaissance. | [implementing-network-deception-with-honeypots](skills/implementing-network-deception-with-honeypots/SKILL.md) |
