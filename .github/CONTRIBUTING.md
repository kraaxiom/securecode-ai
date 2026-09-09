## Flux de contribution

Ce dépôt n'accepte aucune modification directe sur `main`. Toute
contribution passe par une Pull Request depuis un fork, et n'est
fusionnée qu'après revue explicite d'un mainteneur — jamais
automatiquement, quel que soit l'état des vérifications automatiques.

1. Fork du dépôt, branche dédiée (`feat/xxe-rust-example`,
   `fix/cwe-mapping-idor`, ...).
2. Une PR par sujet — pas de PR qui mélange dix vulnérabilités
   différentes, ça ralentit la revue et retarde ta contribution.
3. Respecte le gabarit de la vulnérabilité concernée (voir plus bas) :
   CWE réel et vérifié, portée strictement défensive, pas de payload
   d'exploitation fonctionnel.
4. Un mainteneur revoit, demande des ajustements si besoin, fusionne.
   Aucune contribution n'est appliquée sans cette étape.

