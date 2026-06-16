---
name: cryptography-skill-selector
description: >-
  A cryptography skill-selection catalog for choosing the correct cryptography
  skill from the 26 official security domains. Use this domain catalog before
  opening a specific cryptography skill.
domain: cybersecurity
subdomain: cryptography
tags:
  - cryptography
  - skill-selection
  - cybersecurity-skills
  - domain-catalog
version: "1.0"
author: generated
license: Apache-2.0
---

# Cryptography Skill Catalog

Use this file to choose a cryptography skill. The `When use` text is copied
from each skill's own `description` metadata.

Source subdomains included: `cryptography`.

| No. | When use | Skill |
| ---: | --- | --- |
| 1 | A Certificate Authority (CA) is the trust anchor in a PKI hierarchy, responsible for issuing, signing, and revoking digital certificates. This skill covers building a two-tier CA hierarchy (Root CA + | [configuring-certificate-authority-with-openssl](skills/cyber/configuring-certificate-authority-with-openssl/SKILL.md) |
| 2 | Hardware Security Modules (HSMs) are tamper-resistant physical devices that safeguard cryptographic keys and perform cryptographic operations in a hardened environment. Keys stored in an HSM never lea | [configuring-hsm-for-key-storage](skills/cyber/configuring-hsm-for-key-storage/SKILL.md) |
| 3 | TLS 1.3 (RFC 8446) is the latest version of the Transport Layer Security protocol, providing significant improvements over TLS 1.2 in both security and performance. It reduces handshake latency to 1-R | [configuring-tls-1-3-for-secure-communications](skills/cyber/configuring-tls-1-3-for-secure-communications/SKILL.md) |
| 4 | AES (Advanced Encryption Standard) is a symmetric block cipher standardized by NIST (FIPS 197) used to protect classified and sensitive data. This skill covers implementing AES-256 encryption in GCM m | [implementing-aes-encryption-for-data-at-rest](skills/cyber/implementing-aes-encryption-for-data-at-rest/SKILL.md) |
| 5 | Ed25519 is a high-performance digital signature algorithm using the Edwards curve Curve25519. It provides 128-bit security with 64-byte signatures and 32-byte keys, offering significant advantages ove | [implementing-digital-signatures-with-ed25519](skills/cyber/implementing-digital-signatures-with-ed25519/SKILL.md) |
| 6 | End-to-end encryption (E2EE) ensures that only the communicating parties can read messages, with no intermediary (including the server) able to decrypt them. This skill implements a simplified version | [implementing-end-to-end-encryption-for-messaging](skills/cyber/implementing-end-to-end-encryption-for-messaging/SKILL.md) |
| 7 | Envelope encryption is a strategy where data is encrypted with a data encryption key (DEK), and the DEK itself is encrypted with a master key (KEK) managed by AWS KMS. This approach allows encrypting | [implementing-envelope-encryption-with-aws-kms](skills/cyber/implementing-envelope-encryption-with-aws-kms/SKILL.md) |
| 8 | JSON Web Tokens (JWT) defined in RFC 7519 are compact, URL-safe tokens used for authentication and authorization in web applications. This skill covers implementing secure JWT signing with HMAC-SHA256 | [implementing-jwt-signing-and-verification](skills/cyber/implementing-jwt-signing-and-verification/SKILL.md) |
| 9 | RSA (Rivest-Shamir-Adleman) is the most widely deployed asymmetric cryptographic algorithm, used for digital signatures, key exchange, and encryption. This skill covers generating, storing, rotating, | [implementing-rsa-key-pair-management](skills/cyber/implementing-rsa-key-pair-management/SKILL.md) |
| 10 | Zero-Knowledge Proofs (ZKPs) allow a prover to demonstrate knowledge of a secret (such as a password or private key) without revealing the secret itself. This skill implements the Schnorr identificati | [implementing-zero-knowledge-proof-for-authentication](skills/cyber/implementing-zero-knowledge-proof-for-authentication/SKILL.md) |
| 11 | A cryptographic audit systematically reviews an application's use of cryptographic primitives, protocols, and key management to identify vulnerabilities such as weak algorithms, insecure modes, hardco | [performing-cryptographic-audit-of-application](skills/cyber/performing-cryptographic-audit-of-application/SKILL.md) |
| 12 | Integrate Hardware Security Modules (HSMs) using PKCS#11 interface for cryptographic key management, signing operations, and secure key storage with python-pkcs11, AWS CloudHSM, and YubiHSM2. | [performing-hardware-security-module-integration](skills/cyber/performing-hardware-security-module-integration/SKILL.md) |
| 13 | Hash cracking is an essential skill for penetration testers and security auditors to evaluate password strength. Hashcat is the world's fastest password recovery tool, supporting over 300 hash types w | [performing-hash-cracking-with-hashcat](skills/cyber/performing-hash-cracking-with-hashcat/SKILL.md) |
| 14 | Assesses organizational readiness for post-quantum cryptography migration per NIST FIPS 203/204/205 standards. Performs cryptographic inventory scanning to identify quantum-vulnerable algorithms (RSA, ECDH, ECDSA), evaluates hybrid TLS configurations with X25519MLKEM768, and validates CRYSTALS-Kyber (ML-KEM) and CRYSTALS-Dilithium (ML-DSA) readiness. Implements crypto-agility assessment using oqs-provider for OpenSSL. Use when planning or executing the transition from classical to post-quantum cryptographic algorithms across enterprise infrastructure. | [performing-post-quantum-cryptography-migration](skills/cyber/performing-post-quantum-cryptography-migration/SKILL.md) |
| 15 | SSL/TLS certificate lifecycle management encompasses the full process of requesting, issuing, deploying, monitoring, renewing, and revoking X.509 certificates. Poor certificate management is a leading | [performing-ssl-certificate-lifecycle-management](skills/cyber/performing-ssl-certificate-lifecycle-management/SKILL.md) |
