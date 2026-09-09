# Blind SSRF — Python/Django

CWE-918 — Server-Side Request Forgery (SSRF), variante aveugle.

## Description

`vulnerable.py` enregistre une URL de webhook fournie par l'utilisateur et déclenche une vérification en tâche de fond (`register_webhook_task`). La requête sortante part bien vers l'URL contrôlée par l'attaquant, mais ni son statut, ni son contenu, ni ses erreurs ne sont jamais renvoyés au client : la réponse HTTP se limite à `{"status": "en cours de vérification"}`.

## Pourquoi le code vulnérable est dangereux

L'absence de retour visible ne supprime pas la SSRF, elle la rend seulement plus difficile à exploiter directement : l'attaquant doit s'appuyer sur des canaux indirects (délai de traitement, callback DNS/HTTP hors bande vers un serveur qu'il contrôle) pour confirmer qu'une requête interne a bien été émise. Le serveur reste un relais capable d'atteindre des ressources internes (services d'administration, API de métadonnées cloud, bases de données) sans qu'aucune trace exploitable ne soit générée côté application.

## Explication du correctif

`fixed.py` applique les mêmes contrôles qu'une SSRF classique (whitelist de schémas/domaines, résolution DNS unique puis validation de l'IP contre les plages privées/link-local, DNS pinning, redirections désactivées), car l'absence de retour visible ne dispense d'aucun de ces contrôles. Une différence clé est ajoutée : une **journalisation structurée systématique** de chaque tentative de requête sortante (acceptée, rejetée, ou en échec), indispensable puisque c'est le seul canal de détection disponible en l'absence de réponse exposée au client.

## Notes résiduelles

La journalisation seule ne bloque rien : elle doit être associée à une supervision active (alerte sur rejets répétés, tentatives vers des plages internes) pour détecter un scan en cours. La whitelist doit rester restrictive et revue régulièrement.
