# JWT Weak Secret — CWE-326

La version vulnérable signe les JWT en HS256 avec un secret court et prévisible (`"secret123"`), codé en dur dans le code source. Un tel secret est retrouvable par recherche exhaustive hors-ligne en quelques secondes à quelques minutes avec des outils standards (ex: hashcat), ce qui permet à un attaquant de forger des tokens valides pour n'importe quel utilisateur, y compris des comptes administrateurs.

La version corrigée utilise une clé aléatoire d'au moins 256 bits générée par un CSPRNG (`Keys.hmacShaKeyFor` / `Keys.secretKeyFor`), chargée depuis un gestionnaire de secrets externe (variables d'environnement, Vault) et jamais codée en dur dans le dépôt. La longueur et l'entropie de la clé rendent toute attaque par recherche exhaustive computationnellement infaisable.

**Référence** : CWE-326 (Inadequate Encryption Strength).
