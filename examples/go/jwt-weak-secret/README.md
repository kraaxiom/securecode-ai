# jwt-weak-secret (CWE-326)

La version vulnerable signe les JWT avec un secret HMAC court et code en dur, cassable par force brute ou dictionnaire hors ligne. La version corrigee charge un secret fort (au moins 256 bits d'entropie) depuis un coffre-fort de secrets / variable d'environnement, et refuse explicitement de signer si le secret est absent ou insuffisamment robuste. C'est la remediation standard recommandee contre CWE-326.
