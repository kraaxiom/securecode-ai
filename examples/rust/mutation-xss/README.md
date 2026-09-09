# Mutation XSS (mXSS)

Le handler `save_rich_content` sanitisait le HTML riche via un simple remplacement de chaînes ciblant `<script>`, une approche ignorant les quirks de reparsing du navigateur qui peuvent faire réapparaître du markup actif après vérification (CWE-79). La version corrigée utilise `ammonia`, bibliothèque de sanitisation HTML maintenue s'appuyant sur un vrai parseur conforme aux règles du navigateur, réduisant le risque de contournement par mutation lors des allers-retours DOM/chaîne.
