---
id: weak-random
category: crypto
cwe: CWE-338
owasp: A02:2021-Cryptographic Failures
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Générateur de nombres aléatoires non cryptographique

## Description
L'utilisation d'un générateur de nombres pseudo-aléatoires (PRNG) standard, non conçu pour un usage cryptographique, pour produire des tokens, clés, identifiants de session ou secrets rend ces valeurs prévisibles. Un attaquant capable d'observer une partie de la séquence ou de connaître l'état interne du générateur peut en déduire les valeurs futures ou passées.

## Où ça apparaît typiquement
- Génération de tokens de réinitialisation de mot de passe, clés d'API, ou identifiants de session.
- Génération de valeurs "aléatoires" pour du chiffrement (IV, sel, clé) via `Math.random()`, `rand()`, `random.random()`.
- Génération de codes OTP ou de nonces de sécurité.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à `Math.random()`, `rand()`/`mt_rand()`, `random.random()`/`random.randint()`, `System.Random`, `math/rand` (Go) dans un contexte de génération de secret.
- Absence d'appel à une API CSPRNG dédiée (`crypto.randomBytes`, `random_bytes`, `secrets` module, `SecureRandom`, `crypto/rand`).
- Graine (seed) fixée manuellement pour un générateur utilisé dans un contexte sécuritaire.

## Remédiation
- Utiliser systématiquement un CSPRNG fourni par la plateforme (`crypto.randomBytes` en Node.js, `secrets`/`os.urandom` en Python, `random_bytes` en PHP, `SecureRandom` en Java, `crypto/rand` en Go).
- Ne jamais fixer de graine manuelle pour des valeurs à usage sécuritaire.
- Générer les tokens avec une entropie suffisante (au moins 128 bits).
- Voir `rules/remediation/weak-random.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/weak-random/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cryptographic Storage Cheat Sheet
- CWE-338: Use of Cryptographically Weak Pseudo-Random Number Generator (PRNG)
