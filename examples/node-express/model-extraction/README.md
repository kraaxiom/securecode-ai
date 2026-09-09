## Model Extraction (CWE-200)

Le code vulnérable expose un endpoint d'inférence sans limitation de débit qui renvoie les logits complets du modèle en plus du texte de sortie, permettant à un attaquant d'interroger massivement et systématiquement l'API pour reconstituer un modèle de substitution fonctionnellement équivalent. La correction applique un quota strict par clé API via une limitation de débit, retire les logits bruts de la réponse pour ne renvoyer que le texte nécessaire au cas d'usage, et ajoute une détection des patterns de requêtes évoquant une extraction systématique.
