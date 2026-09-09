# Stored XSS

Le handler `view_profile` affichait la biographie stockée en base directement dans le HTML sans encodage, exposant tous les visiteurs de la page à une charge persistée (CWE-79), considérée plus critique que le XSS réfléchi car ne nécessitant aucune interaction ciblée. La version corrigée échappe systématiquement la donnée persistée avant affichage, quel que soit son origine en base.
