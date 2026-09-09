# HTTP Parameter Pollution (CWE-235)

La version vulnérable accepte le paramètre `role` via le binding automatique `[FromQuery]`, sans vérifier s'il a été envoyé plusieurs fois, ce qui peut créer une divergence d'interprétation avec un WAF ou un proxy en amont et permettre un contournement de contrôle. La correction lit `Request.Query["role"]` brut et rejette explicitement toute requête où le paramètre apparaît plus d'une fois, garantissant un traitement cohérent et déterministe.
