# HTTP Response Splitting

La version vulnérable affecte `next_url` directement à l'en-tête `Location` sans validation, ce qui correspond à CWE-113 : une entrée contenant `\r\n` pourrait fractionner la réponse HTTP et permettre l'injection d'une seconde réponse. La version corrigée restreint `next_url` à une liste blanche de chemins internes (`ALLOWED_PATHS`) et rejette toute valeur contenant un schéma ou un hôte externe.
