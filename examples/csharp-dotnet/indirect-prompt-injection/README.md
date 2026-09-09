# Indirect Prompt Injection (CWE-1427)

La version vulnérable récupère le contenu d'une page web ou d'un document externe et l'injecte tel quel dans le contexte du modèle, aux côtés des outils actifs de l'agent, sans marquage de provenance ni séparation entre « contenu à analyser » et « instructions à exécuter ». La correction délimite explicitement le contenu externe comme non fiable (balises de provenance), désactive les outils sensibles pendant la phase d'analyse de ce contenu, et exige une revue humaine avant toute action à fort impact déclenchée à la suite de cette analyse.
