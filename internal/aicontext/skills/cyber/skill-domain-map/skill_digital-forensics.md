---
name: digital-forensics-skill-selector
description: >-
  A digital forensics skill-selection catalog for choosing the correct digital
  forensics skill from the 26 official security domains. Use this domain catalog
  before opening a specific digital forensics skill.
domain: cybersecurity
subdomain: digital-forensics
tags:
  - digital-forensics
  - skill-selection
  - cybersecurity-skills
  - domain-catalog
version: "1.0"
author: generated
license: Apache-2.0
---

# Digital Forensics Skill Catalog

Use this file to choose a digital forensics skill. The `When use` text is copied
from each skill's own `description` metadata.

Source subdomains included: `digital-forensics`.

| No. | When use | Skill |
| ---: | --- | --- |
| 1 | Create forensically sound bit-for-bit disk images using dd and dcfldd while preserving evidence integrity through hash verification. | [acquiring-disk-image-with-dd-and-dcfldd](skills/cyber/acquiring-disk-image-with-dd-and-dcfldd/SKILL.md) |
| 2 | Analyze Chromium-based browser artifacts using Hindsight to extract browsing history, downloads, cookies, cached content, autofill data, saved passwords, and browser extensions from Chrome, Edge, Brave, and Opera for forensic investigation. | [analyzing-browser-forensics-with-hindsight](skills/cyber/analyzing-browser-forensics-with-hindsight/SKILL.md) |
| 3 | Perform comprehensive forensic analysis of disk images using Autopsy to recover files, examine artifacts, and build investigation timelines. | [analyzing-disk-image-with-autopsy](skills/cyber/analyzing-disk-image-with-autopsy/SKILL.md) |
| 4 | Investigate compromised Docker containers by analyzing images, layers, volumes, logs, and runtime artifacts to identify malicious activity and evidence. | [analyzing-docker-container-forensics](skills/cyber/analyzing-docker-container-forensics/SKILL.md) |
| 5 | Parse and analyze email headers to trace the origin of phishing emails, verify sender authenticity, and identify spoofing through SPF, DKIM, and DMARC validation. | [analyzing-email-headers-for-phishing-investigation](skills/cyber/analyzing-email-headers-for-phishing-investigation/SKILL.md) |
| 6 | Detect kernel-level rootkits in Linux memory dumps using Volatility3 linux plugins (check_syscall, lsmod, hidden_modules), rkhunter system scanning, and /proc vs /sys discrepancy analysis to identify hooked syscalls, hidden kernel modules, and tampered system structures. | [analyzing-linux-kernel-rootkits](skills/cyber/analyzing-linux-kernel-rootkits/SKILL.md) |
| 7 | Examine Linux system artifacts including auth logs, cron jobs, shell history, and system configuration to uncover evidence of compromise or unauthorized activity. | [analyzing-linux-system-artifacts](skills/cyber/analyzing-linux-system-artifacts/SKILL.md) |
| 8 | Analyze Windows LNK shortcut files and Jump List artifacts to establish evidence of file access, program execution, and user activity using LECmd, JLECmd, and manual binary parsing of the Shell Link Binary format. | [analyzing-lnk-file-and-jump-list-artifacts](skills/cyber/analyzing-lnk-file-and-jump-list-artifacts/SKILL.md) |
| 9 | Analyze the NTFS Master File Table ($MFT) to recover metadata and content of deleted files by examining MFT record entries, $LogFile, $UsnJrnl, and MFT slack space using MFTECmd, analyzeMFT, and X-Ways Forensics. | [analyzing-mft-for-deleted-file-recovery](skills/cyber/analyzing-mft-for-deleted-file-recovery/SKILL.md) |
| 10 | Analyze Microsoft Outlook PST and OST files for email forensic evidence including message content, headers, attachments, deleted items, and metadata using libpff, pst-utils, and forensic email analysis tools for legal investigations and incident response. | [analyzing-outlook-pst-for-email-forensics](skills/cyber/analyzing-outlook-pst-for-email-forensics/SKILL.md) |
| 11 | Parse Windows Prefetch files to determine program execution history including run counts, timestamps, and referenced files for forensic investigation. | [analyzing-prefetch-files-for-execution-history](skills/cyber/analyzing-prefetch-files-for-execution-history/SKILL.md) |
| 12 | Examine file system slack space, MFT entries, USN journal, and alternate data streams to recover hidden data and reconstruct file activity on NTFS volumes. | [analyzing-slack-space-and-file-system-artifacts](skills/cyber/analyzing-slack-space-and-file-system-artifacts/SKILL.md) |
| 13 | Investigate USB device connection history from Windows registry, event logs, and setupapi logs to track removable media usage and potential data exfiltration. | [analyzing-usb-device-connection-history](skills/cyber/analyzing-usb-device-connection-history/SKILL.md) |
| 14 | Parses and analyzes the Windows Amcache.hve registry hive to extract evidence of program execution, application installation, and driver loading for digital forensics investigations. Uses Eric Zimmerman's AmcacheParser and Timeline Explorer for artifact extraction, SHA-1 hash correlation with threat intel, and timeline reconstruction. Activates for requests involving Amcache forensics, program execution evidence, Windows artifact analysis, or application compatibility cache investigation. | [analyzing-windows-amcache-artifacts](skills/cyber/analyzing-windows-amcache-artifacts/SKILL.md) |
| 15 | Parse Windows LNK shortcut files to extract target paths, timestamps, volume information, and machine identifiers for forensic timeline reconstruction. | [analyzing-windows-lnk-files-for-artifacts](skills/cyber/analyzing-windows-lnk-files-for-artifacts/SKILL.md) |
| 16 | Parse Windows Prefetch files using the windowsprefetch Python library to reconstruct application execution history, detect renamed or masquerading binaries, and identify suspicious program execution patterns. | [analyzing-windows-prefetch-with-python](skills/cyber/analyzing-windows-prefetch-with-python/SKILL.md) |
| 17 | Extract and analyze Windows Registry hives to uncover user activity, installed software, autostart entries, and evidence of system compromise. | [analyzing-windows-registry-for-artifacts](skills/cyber/analyzing-windows-registry-for-artifacts/SKILL.md) |
| 18 | Analyze Windows Shellbag registry artifacts to reconstruct folder browsing activity, detect access to removable media and network shares, and establish user interaction with directories even after deletion using SBECmd and ShellBags Explorer. | [analyzing-windows-shellbag-artifacts](skills/cyber/analyzing-windows-shellbag-artifacts/SKILL.md) |
| 19 | Extract and analyze browser history, cookies, cache, downloads, and bookmarks from Chrome, Firefox, and Edge for forensic evidence of user web activity. | [extracting-browser-history-artifacts](skills/cyber/extracting-browser-history-artifacts/SKILL.md) |
| 20 | Extract cached credentials, password hashes, Kerberos tickets, and authentication tokens from memory dumps using Volatility and Mimikatz for forensic investigation. | [extracting-credentials-from-memory-dump](skills/cyber/extracting-credentials-from-memory-dump/SKILL.md) |
| 21 | Extract, parse, and analyze Windows Event Logs (EVTX) using Chainsaw, Hayabusa, and EvtxECmd to detect lateral movement, persistence, and privilege escalation. | [extracting-windows-event-logs-artifacts](skills/cyber/extracting-windows-event-logs-artifacts/SKILL.md) |
| 22 | Identify, collect, and analyze ransomware attack artifacts to determine the variant, initial access vector, encryption scope, and recovery options. | [investigating-ransomware-attack-artifacts](skills/cyber/investigating-ransomware-attack-artifacts/SKILL.md) |
| 23 | Conduct forensic investigations in cloud environments by collecting and analyzing logs, snapshots, and metadata from AWS, Azure, and GCP services. | [performing-cloud-forensics-investigation](skills/cyber/performing-cloud-forensics-investigation/SKILL.md) |
| 24 | Perform forensic acquisition and analysis of cloud storage services including Google Drive, OneDrive, Dropbox, and Box by collecting both API-based remote data and local sync client artifacts from endpoint devices. | [performing-cloud-storage-forensic-acquisition](skills/cyber/performing-cloud-storage-forensic-acquisition/SKILL.md) |
| 25 | Recover files from disk images and unallocated space using Foremost's header-footer signature carving to extract evidence regardless of file system state. | [performing-file-carving-with-foremost](skills/cyber/performing-file-carving-with-foremost/SKILL.md) |
| 26 | Perform forensic investigation of Linux system logs including syslog, auth.log, systemd journal, kern.log, and application logs to reconstruct user activity, detect unauthorized access, and establish event timelines on compromised Linux systems. | [performing-linux-log-forensics-investigation](skills/cyber/performing-linux-log-forensics-investigation/SKILL.md) |
| 27 | Collect, parse, and correlate system, application, and security logs to reconstruct events and establish timelines during forensic investigations. | [performing-log-analysis-for-forensic-investigation](skills/cyber/performing-log-analysis-for-forensic-investigation/SKILL.md) |
| 28 | Systematically investigate all persistence mechanisms on Windows and Linux systems to identify how malware survives reboots and maintains access. | [performing-malware-persistence-investigation](skills/cyber/performing-malware-persistence-investigation/SKILL.md) |
| 29 | Analyze volatile memory dumps using Volatility 3 to extract running processes, network connections, loaded modules, and evidence of malicious activity. | [performing-memory-forensics-with-volatility3](skills/cyber/performing-memory-forensics-with-volatility3/SKILL.md) |
| 30 | Acquire and analyze mobile device data using Cellebrite UFED and open-source tools to extract communications, location data, and application artifacts. | [performing-mobile-device-forensics-with-cellebrite](skills/cyber/performing-mobile-device-forensics-with-cellebrite/SKILL.md) |
| 31 | Capture and analyze network traffic using Wireshark and tshark to reconstruct network events, extract artifacts, and identify malicious communications. | [performing-network-forensics-with-wireshark](skills/cyber/performing-network-forensics-with-wireshark/SKILL.md) |
| 32 | Perform forensic analysis of network packet captures (PCAP/PCAPNG) using Wireshark, tshark, and tcpdump to reconstruct network communications, extract transferred files, identify malicious traffic, and establish evidence of data exfiltration or command-and-control activity. | [performing-network-packet-capture-analysis](skills/cyber/performing-network-packet-capture-analysis/SKILL.md) |
| 33 | Perform forensic analysis of SQLite databases to recover deleted records from freelists and WAL files, decode encoded timestamps, and extract evidence from browser history, messaging apps, and mobile device databases. | [performing-sqlite-database-forensics](skills/cyber/performing-sqlite-database-forensics/SKILL.md) |
| 34 | Detect and extract hidden data embedded in images, audio, and other media files using steganalysis tools to uncover covert communication channels. | [performing-steganography-detection](skills/cyber/performing-steganography-detection/SKILL.md) |
| 35 | Build comprehensive forensic super-timelines using Plaso (log2timeline) to correlate events across file systems, logs, and artifacts into a unified chronological view. | [performing-timeline-reconstruction-with-plaso](skills/cyber/performing-timeline-reconstruction-with-plaso/SKILL.md) |
| 36 | Perform comprehensive Windows forensic artifact analysis using Eric Zimmerman's open-source EZ Tools suite including KAPE, MFTECmd, PECmd, LECmd, JLECmd, and Timeline Explorer for parsing registry hives, prefetch files, event logs, and file system metadata. | [performing-windows-artifact-analysis-with-eric-zimmerman-tools](skills/cyber/performing-windows-artifact-analysis-with-eric-zimmerman-tools/SKILL.md) |
| 37 | Recover deleted files from disk images and storage media using PhotoRec's file signature-based carving engine regardless of file system damage. | [recovering-deleted-files-with-photorec](skills/cyber/recovering-deleted-files-with-photorec/SKILL.md) |
