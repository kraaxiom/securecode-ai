# Conteneur Azure Blob Storage public (CWE-284)

Le code vulnérable crée un conteneur avec `PublicAccess::Blob` et génère un jeton SAS aux permissions larges (lecture, écriture, suppression) expirant dans dix ans, ce qui équivaut à un accès public permanent et non authentifié. La correction crée le conteneur avec `PublicAccess::None` (accès privé par défaut) et limite tout partage ponctuel à un SAS en lecture seule expirant après 15 minutes. Élimine la classe de vulnérabilité CWE-284 (Improper Access Control).
