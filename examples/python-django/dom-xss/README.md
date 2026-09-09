# DOM-based XSS (CWE-79)

Le code vulnérable transmettait le paramètre de recherche `q` au template avec le filtre `|safe`, ce qui permettait d'injecter du contenu directement dans un bloc `<script>` destiné à écrire dans le DOM via `innerHTML`, un sink dangereux. Le correctif retire `|safe` (Django échappe désormais la valeur automatiquement) et remplace le sink côté client par `textContent`, une API texte qui n'interprète jamais de HTML, supprimant la classe de vulnérabilité indépendamment de la source de la donnée.
