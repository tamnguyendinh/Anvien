---
name: red-teaming-skill-selector
description: >-
  A red teaming skill-selection catalog for choosing the correct red teaming
  skill from the 26 official security domains. Use this domain catalog before
  opening a specific red teaming skill.
domain: cybersecurity
subdomain: red-teaming
tags:
  - red-teaming
  - skill-selection
  - cybersecurity-skills
  - domain-catalog
version: "1.0"
author: generated
license: Apache-2.0
---

# Red Teaming Skill Catalog

Use this file to choose a red teaming skill. The `When use` text is copied
from each skill's own `description` metadata.

Source subdomains included: `purple-team`, `red-team`, `red-teaming`.

| No. | When use | Skill |
| ---: | --- | --- |
| 1 | Build and configure a resilient command-and-control infrastructure using BishopFox's Sliver C2 framework with redirectors, HTTPS listeners, and multi-operator support for authorized red team engagements. | [building-c2-infrastructure-with-sliver-framework](skills/cyber/building-c2-infrastructure-with-sliver-framework/SKILL.md) |
| 2 | Deploy and configure the Havoc C2 framework with teamserver, HTTPS listeners, redirectors, and Demon agents for authorized red team operations. | [building-red-team-c2-infrastructure-with-havoc](skills/cyber/building-red-team-c2-infrastructure-with-havoc/SKILL.md) |
| 3 | Perform DCSync attacks to replicate Active Directory credentials and establish domain persistence by extracting KRBTGT, Domain Admin, and service account hashes for Golden Ticket creation. | [conducting-domain-persistence-with-dcsync](skills/cyber/conducting-domain-persistence-with-dcsync/SKILL.md) |
| 4 | Plan and execute a comprehensive red team engagement covering reconnaissance through post-exploitation using MITRE ATT&CK-aligned TTPs to evaluate an organization's detection and response capabilities. | [conducting-full-scope-red-team-engagement](skills/cyber/conducting-full-scope-red-team-engagement/SKILL.md) |
| 5 | Conduct internal Active Directory reconnaissance using BloodHound Community Edition to map attack paths, identify privilege escalation chains, and discover misconfigurations in domain environments. | [conducting-internal-reconnaissance-with-bloodhound-ce](skills/cyber/conducting-internal-reconnaissance-with-bloodhound-ce/SKILL.md) |
| 6 | Pass-the-Ticket (PtT) is a lateral movement technique that uses stolen Kerberos tickets (TGT or TGS) to authenticate to services without knowing the user's password. By extracting Kerberos tickets fro | [conducting-pass-the-ticket-attack](skills/cyber/conducting-pass-the-ticket-attack/SKILL.md) |
| 7 | Plan and execute authorized vishing (voice phishing) pretext calls to assess employee susceptibility to social engineering and evaluate security awareness controls. | [conducting-social-engineering-pretext-call](skills/cyber/conducting-social-engineering-pretext-call/SKILL.md) |
| 8 | Spearphishing simulation is a targeted social engineering attack vector used by red teams to gain initial access. Unlike broad phishing campaigns, spearphishing uses OSINT-derived intelligence to craf | [conducting-spearphishing-simulation-campaign](skills/cyber/conducting-spearphishing-simulation-campaign/SKILL.md) |
| 9 | Red team engagement planning is the foundational phase that defines scope, objectives, rules of engagement (ROE), threat model selection, and operational timelines before any offensive testing begins. | [executing-red-team-engagement-planning](skills/cyber/executing-red-team-engagement-planning/SKILL.md) |
| 10 | Exploit misconfigured Active Directory Certificate Services (AD CS) ESC1 vulnerability to request certificates as high-privileged users and escalate domain privileges during authorized red team assessments. | [exploiting-active-directory-certificate-services-esc1](skills/cyber/exploiting-active-directory-certificate-services-esc1/SKILL.md) |
| 11 | BloodHound is a graph-based Active Directory reconnaissance tool that uses graph theory to reveal hidden and unintended relationships within AD environments. Red teams use BloodHound to identify attac | [exploiting-active-directory-with-bloodhound](skills/cyber/exploiting-active-directory-with-bloodhound/SKILL.md) |
| 12 | Exploit Kerberos Constrained Delegation misconfigurations in Active Directory to impersonate privileged users via S4U2self and S4U2proxy extensions for lateral movement and privilege escalation. | [exploiting-constrained-delegation-abuse](skills/cyber/exploiting-constrained-delegation-abuse/SKILL.md) |
| 13 | Perform Kerberoasting attacks using Impacket's GetUserSPNs to extract and crack Kerberos TGS tickets for Active Directory service accounts. | [exploiting-kerberoasting-with-impacket](skills/cyber/exploiting-kerberoasting-with-impacket/SKILL.md) |
| 14 | MS17-010 (EternalBlue) is a critical vulnerability in Microsoft's SMBv1 implementation that allows remote code execution. Originally discovered by the NSA and leaked by the Shadow Brokers in 2017, it | [exploiting-ms17-010-eternalblue-vulnerability](skills/cyber/exploiting-ms17-010-eternalblue-vulnerability/SKILL.md) |
| 15 | Exploit the noPac vulnerability chain (CVE-2021-42278 sAMAccountName spoofing and CVE-2021-42287 KDC PAC confusion) to escalate from standard domain user to Domain Admin in Active Directory environments. | [exploiting-nopac-cve-2021-42278-42287](skills/cyber/exploiting-nopac-cve-2021-42278-42287/SKILL.md) |
| 16 | Exploit the Zerologon vulnerability (CVE-2020-1472) in the Netlogon Remote Protocol to achieve domain controller compromise by resetting the machine account password to empty. | [exploiting-zerologon-vulnerability-cve-2020-1472](skills/cyber/exploiting-zerologon-vulnerability-cve-2020-1472/SKILL.md) |
| 17 | Use BloodHound and SharpHound to enumerate Active Directory relationships and identify attack paths from compromised users to Domain Admin. | [performing-active-directory-bloodhound-analysis](skills/cyber/performing-active-directory-bloodhound-analysis/SKILL.md) |
| 18 | Enumerate and audit Active Directory forest trust relationships using impacket for SID filtering analysis, trust key extraction, cross-forest SID history abuse detection, and inter-realm Kerberos ticket assessment. | [performing-active-directory-forest-trust-attack](skills/cyber/performing-active-directory-forest-trust-attack/SKILL.md) |
| 19 | Extract stored credentials from compromised endpoints using the LaZagne post-exploitation tool to recover passwords from browsers, databases, system vaults, and applications during authorized red team operations. | [performing-credential-access-with-lazagne](skills/cyber/performing-credential-access-with-lazagne/SKILL.md) |
| 20 | Perform authorized initial access using EvilGinx3 adversary-in-the-middle phishing framework to capture session tokens and bypass multi-factor authentication during red team engagements. | [performing-initial-access-with-evilginx3](skills/cyber/performing-initial-access-with-evilginx3/SKILL.md) |
| 21 | Kerberoasting is a post-exploitation technique that targets service accounts in Active Directory by requesting Kerberos TGS (Ticket Granting Service) tickets for accounts with Service Principal Names | [performing-kerberoasting-attack](skills/cyber/performing-kerberoasting-attack/SKILL.md) |
| 22 | Perform lateral movement across Windows networks using WMI-based remote execution techniques including Impacket wmiexec.py, CrackMapExec, and native WMI commands for stealthy post-exploitation during red team engagements. | [performing-lateral-movement-with-wmiexec](skills/cyber/performing-lateral-movement-with-wmiexec/SKILL.md) |
| 23 | Open Source Intelligence (OSINT) gathering is the first active phase of a red team engagement, where operators collect publicly available information about the target organization to identify attack s | [performing-open-source-intelligence-gathering](skills/cyber/performing-open-source-intelligence-gathering/SKILL.md) |
| 24 | Conduct authorized physical penetration testing using tailgating, badge cloning, lock bypassing, and rogue device deployment to evaluate facility security controls. | [performing-physical-intrusion-assessment](skills/cyber/performing-physical-intrusion-assessment/SKILL.md) |
| 25 | Linux privilege escalation involves elevating from a low-privilege user account to root access on a compromised system. Red teams exploit misconfigurations, vulnerable services, kernel exploits, and w | [performing-privilege-escalation-on-linux](skills/cyber/performing-privilege-escalation-on-linux/SKILL.md) |
| 26 | Executes Atomic Red Team tests mapped to MITRE ATT&CK techniques, performs coverage gap analysis across the ATT&CK matrix, and runs detection validation loops to measure blue team visibility. Covers Invoke-AtomicRedTeam PowerShell execution, ATT&CK Navigator layer generation for heatmaps, Sigma rule correlation, and continuous atomic testing pipelines. Activates for requests involving purple team exercises, atomic test execution, ATT&CK coverage assessment, detection engineering validation, or adversary emulation testing. | [performing-purple-team-atomic-testing](skills/cyber/performing-purple-team-atomic-testing/SKILL.md) |
| 27 | Conduct red team operations using the Covenant C2 framework for authorized adversary simulation, including listener setup, grunt deployment, task execution, and lateral movement tracking. | [performing-red-team-with-covenant](skills/cyber/performing-red-team-with-covenant/SKILL.md) |
