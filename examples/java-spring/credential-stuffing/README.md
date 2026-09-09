# Credential Stuffing — CWE-307

La version vulnérable accepte un volume illimité de tentatives de connexion, sans MFA et sans aucune détection de vélocité. Un attaquant peut donc rejouer automatiquement, contre ce service, des paires identifiant/mot de passe issues de fuites d'autres sites (le "credential stuffing" exploite la réutilisation de mots de passe entre services), avec un taux de succès non négligeable et sans alerte.

La version corrigée ajoute une détection de vélocité qui repère un nombre anormal de tentatives sur des comptes **différents** depuis une même source (signature distinctive du credential stuffing, à l'inverse du brute force ciblé sur un seul compte), rend le MFA obligatoire pour finaliser toute connexion, et vérifie — même en cas de mot de passe correct — si celui-ci figure dans une base de mots de passe compromis connus, pour forcer son renouvellement immédiat le cas échéant.

**Référence** : CWE-307 (Improper Restriction of Excessive Authentication Attempts).
