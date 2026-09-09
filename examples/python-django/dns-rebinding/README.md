# DNS Rebinding (contournement de filtre SSRF) — Python/Django

CWE-918 — Server-Side Request Forgery (SSRF), contournement par DNS rebinding.

## Description

`vulnerable.py` valide une URL de webhook en deux temps : une première résolution DNS (`socket.gethostbyname`) vérifie que l'IP obtenue est publique, puis une requête HTTP séparée est envoyée via `requests.get()`, qui résout le nom de domaine une seconde fois de manière indépendante.

## Pourquoi le code vulnérable est dangereux

Un attaquant qui contrôle le serveur DNS de son domaine peut configurer un TTL très court et faire varier la réponse : au moment de la validation, le domaine résout vers une IP publique légitime ; au moment de la requête HTTP réelle (quelques millisecondes après), il résout vers une IP interne (`127.0.0.1`, `169.254.169.254`, etc.). Le filtre basé uniquement sur la validation "avant" la requête est ainsi contourné, car les deux résolutions DNS sont indépendantes (TOCTOU — Time-Of-Check to Time-Of-Use).

## Explication du correctif

`fixed.py` élimine la fenêtre de rebinding en ne résolvant le DNS **qu'une seule fois** (`resolve_pin_and_validate`) : l'IP obtenue est validée puis directement utilisée pour la connexion HTTP (DNS pinning), en conservant l'en-tête `Host` pour que le serveur cible et la vérification TLS restent cohérents avec le nom de domaine d'origine. Il n'existe alors plus de seconde résolution DNS que l'attaquant pourrait faire varier entre la validation et l'usage réel.

## Notes résiduelles

Le pinning protège contre le rebinding applicatif, mais pas contre un attaquant qui contrôlerait déjà l'infrastructure réseau elle-même. Pour une défense complète, combiner ce contrôle avec une isolation réseau du service effectuant les requêtes sortantes (segment sans accès aux ressources internes sensibles).
