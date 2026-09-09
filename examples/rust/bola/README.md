# BOLA — Broken Object Level Authorization

`vulnerable.rs` récupère une commande uniquement par son identifiant, sans vérifier qu'elle appartient à l'utilisateur authentifié, illustrant CWE-639 (Authorization Bypass Through User-Controlled Key). `fixed.rs` ajoute la clause `owner_id = $2` directement dans la requête SQL, de sorte qu'un utilisateur ne peut jamais récupérer les données d'une commande qui ne lui appartient pas, conformément à `rules/remediation/bola.md`. Le filtrage d'appartenance est appliqué au niveau de la requête de données elle-même, et non en post-traitement applicatif.
