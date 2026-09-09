# mfa-bypass (CWE-287)

La version vulnerable delivre le jeton final des lors que le client envoie un flag `MFAVerified: true`, sans jamais verifier reellement le code OTP cote serveur, permettant un contournement trivial du second facteur. La version corrigee separe le login en deux etapes serveur : une session `pending_mfa` est creee apres le mot de passe, et le jeton n'est emis qu'apres verification serveur du code OTP lie a cette session. C'est la remediation standard recommandee contre CWE-287.
