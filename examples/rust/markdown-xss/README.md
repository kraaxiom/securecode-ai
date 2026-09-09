# Markdown XSS

Le handler `render_comment` convertissait le Markdown utilisateur en HTML puis l'insérait directement dans la réponse, sans sanitiser d'éventuelles constructions actives (CWE-79). La version corrigée conserve un parseur Markdown sans extension HTML brut et fait systématiquement passer le HTML généré par `ammonia`, une bibliothèque de sanitisation à liste blanche restreignant balises, attributs et schémas d'URL autorisés (`http`, `https`, `mailto`).
