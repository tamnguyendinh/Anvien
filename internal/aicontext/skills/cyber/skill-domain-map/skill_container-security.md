---
name: container-security-skill-selector
description: >-
  A container security skill-selection catalog for choosing the correct
  container security skill from the 26 official security domains. Use this
  domain catalog before opening a specific container security skill.
domain: cybersecurity
subdomain: container-security
tags:
  - container-security
  - skill-selection
  - cybersecurity-skills
  - domain-catalog
version: "1.0"
author: generated
license: Apache-2.0
---

# Container Security Skill Catalog

Use this file to choose a container security skill. The `When use` text is copied
from each skill's own `description` metadata.

Source subdomains included: `container-security`.

| No. | When use | Skill |
| ---: | --- | --- |
| 1 | Parses Kubernetes API server audit logs (JSON lines) to detect exec-into-pod, secret access, RBAC modifications, privileged pod creation, and anonymous API access. Builds threat detection rules from audit event patterns. Use when investigating Kubernetes cluster compromise or building k8s-specific SIEM detection rules. | [analyzing-kubernetes-audit-logs](skills/analyzing-kubernetes-audit-logs/SKILL.md) |
| 2 | Detect unauthorized modifications to running containers by monitoring for binary execution drift, file system changes, and configuration deviations from the original container image. | [detecting-container-drift-at-runtime](skills/detecting-container-drift-at-runtime/SKILL.md) |
| 3 | Container escape is a critical attack technique where an adversary breaks out of container isolation to access the host system or other containers. Detection involves monitoring for escape indicators | [detecting-container-escape-attempts](skills/detecting-container-escape-attempts/SKILL.md) |
| 4 | Detect container escape attempts in real-time using Falco runtime security rules that monitor syscalls, file access, and privilege escalation. | [detecting-container-escape-with-falco-rules](skills/detecting-container-escape-with-falco-rules/SKILL.md) |
| 5 | Detect and prevent privilege escalation in Kubernetes pods by monitoring security contexts, capabilities, and syscall patterns with Falco and OPA policies. | [detecting-privilege-escalation-in-kubernetes-pods](skills/detecting-privilege-escalation-in-kubernetes-pods/SKILL.md) |
| 6 | Hardening Docker containers for production involves applying security best practices aligned with CIS Docker Benchmark v1.8.0 to minimize attack surface, prevent privilege escalation, and enforce leas | [hardening-docker-containers-for-production](skills/hardening-docker-containers-for-production/SKILL.md) |
| 7 | Harden the Docker daemon by configuring daemon.json with user namespace remapping, TLS authentication, rootless mode, and CIS benchmark controls. | [hardening-docker-daemon-configuration](skills/hardening-docker-daemon-configuration/SKILL.md) |
| 8 | Reduce container attack surface by building application images on Google distroless base images that contain only the application runtime with no shell, package manager, or unnecessary OS utilities. | [implementing-container-image-minimal-base-with-distroless](skills/implementing-container-image-minimal-base-with-distroless/SKILL.md) |
| 9 | Enforce Kubernetes network segmentation using Calico CNI network policies and global network policies to control pod-to-pod traffic, restrict egress, and implement zero-trust microsegmentation. | [implementing-container-network-policies-with-calico](skills/implementing-container-network-policies-with-calico/SKILL.md) |
| 10 | Sign and verify container image provenance using Sigstore Cosign with keyless OIDC-based signing, attestations, and Kubernetes admission enforcement. | [implementing-image-provenance-verification-with-cosign](skills/implementing-image-provenance-verification-with-cosign/SKILL.md) |
| 11 | Implement Kubernetes network segmentation using Calico NetworkPolicy and GlobalNetworkPolicy for zero-trust pod-to-pod communication. | [implementing-kubernetes-network-policy-with-calico](skills/implementing-kubernetes-network-policy-with-calico/SKILL.md) |
| 12 | Pod Security Standards (PSS) define three levels of security policies -- Privileged, Baseline, and Restricted -- enforced by the Pod Security Admission (PSA) controller built into Kubernetes 1.25+. PS | [implementing-kubernetes-pod-security-standards](skills/implementing-kubernetes-pod-security-standards/SKILL.md) |
| 13 | Kubernetes NetworkPolicies provide pod-level network segmentation by defining ingress and egress rules that control traffic flow between pods, namespaces, and external endpoints. Combined with CNI plu | [implementing-network-policies-for-kubernetes](skills/implementing-network-policies-for-kubernetes/SKILL.md) |
| 14 | Enforce Kubernetes admission policies using OPA Gatekeeper with ConstraintTemplates, Rego rules, and the Gatekeeper policy library. | [implementing-opa-gatekeeper-for-policy-enforcement](skills/implementing-opa-gatekeeper-for-policy-enforcement/SKILL.md) |
| 15 | Implement Kubernetes Pod Security Admission to enforce baseline and restricted security profiles at namespace level using built-in admission controller. | [implementing-pod-security-admission-controller](skills/implementing-pod-security-admission-controller/SKILL.md) |
| 16 | Harden Kubernetes Role-Based Access Control by implementing least-privilege policies, auditing role bindings, eliminating cluster-admin sprawl, and integrating external identity providers. | [implementing-rbac-hardening-for-kubernetes](skills/implementing-rbac-hardening-for-kubernetes/SKILL.md) |
| 17 | Implement eBPF-based runtime security observability and enforcement in Kubernetes clusters using Cilium Tetragon for kernel-level threat detection and policy enforcement. | [implementing-runtime-security-with-tetragon](skills/implementing-runtime-security-with-tetragon/SKILL.md) |
| 18 | Implement software supply chain integrity verification for container builds using the in-toto framework to create cryptographically signed attestations across CI/CD pipeline steps. | [implementing-supply-chain-security-with-in-toto](skills/implementing-supply-chain-security-with-in-toto/SKILL.md) |
| 19 | Detects container escape attempts by analyzing namespace configurations, privileged container checks, dangerous capability assignments, and host path mounts using the kubernetes Python client. Identifies CVE-2022-0492 style escapes via cgroup abuse. Use when auditing container security posture or investigating escape attempts. | [performing-container-escape-detection](skills/performing-container-escape-detection/SKILL.md) |
| 20 | Scan container images, filesystems, and Kubernetes manifests for vulnerabilities, misconfigurations, exposed secrets, and license compliance issues using Aqua Security Trivy with SBOM generation and CI/CD integration. | [performing-container-security-scanning-with-trivy](skills/performing-container-security-scanning-with-trivy/SKILL.md) |
| 21 | Docker Bench for Security is an open-source script that checks dozens of common best practices around deploying Docker containers in production. Based on the CIS Docker Benchmark, it audits host confi | [performing-docker-bench-security-assessment](skills/performing-docker-bench-security-assessment/SKILL.md) |
| 22 | Audit Kubernetes cluster security posture against CIS benchmarks using kube-bench with automated checks for control plane, worker nodes, and RBAC. | [performing-kubernetes-cis-benchmark-with-kube-bench](skills/performing-kubernetes-cis-benchmark-with-kube-bench/SKILL.md) |
| 23 | Assess the security posture of Kubernetes etcd clusters by evaluating encryption at rest, TLS configuration, access controls, backup encryption, and network isolation. | [performing-kubernetes-etcd-security-assessment](skills/performing-kubernetes-etcd-security-assessment/SKILL.md) |
| 24 | Kubernetes penetration testing systematically evaluates cluster security by simulating attacker techniques against the API server, kubelet, etcd, pods, RBAC, network policies, and secrets. Using tools | [performing-kubernetes-penetration-testing](skills/performing-kubernetes-penetration-testing/SKILL.md) |
| 25 | Scan container images for known vulnerabilities using Anchore Grype with SBOM-based matching and configurable severity thresholds. | [scanning-container-images-with-grype](skills/scanning-container-images-with-grype/SKILL.md) |
| 26 | Trivy is a comprehensive open-source vulnerability scanner by Aqua Security that detects vulnerabilities in OS packages, language-specific dependencies, misconfigurations, secrets, and license violati | [scanning-docker-images-with-trivy](skills/scanning-docker-images-with-trivy/SKILL.md) |
| 27 | Perform security risk analysis on Kubernetes resource manifests using Kubesec to identify misconfigurations, privilege escalation risks, and deviations from security best practices. | [scanning-kubernetes-manifests-with-kubesec](skills/scanning-kubernetes-manifests-with-kubesec/SKILL.md) |
| 28 | Harbor is an open-source container registry that provides security features including vulnerability scanning (integrated Trivy), image signing (Notary/Cosign), RBAC, content trust policies, replicatio | [securing-container-registry-with-harbor](skills/securing-container-registry-with-harbor/SKILL.md) |
| 29 | Secure Helm chart deployments by validating chart integrity, scanning templates for misconfigurations, and enforcing security contexts in Kubernetes releases. | [securing-helm-chart-deployments](skills/securing-helm-chart-deployments/SKILL.md) |
