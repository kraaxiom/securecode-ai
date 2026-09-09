# Model Extraction (CWE-200)

La version vulnérable expose un endpoint d'inférence sans aucune limite de débit ni quota par client, et renvoie les logits/probabilités bruts du modèle en plus du texte généré, ce qui facilite l'interrogation massive et systématique nécessaire à la reconstruction d'un modèle équivalent. La correction applique une limitation de débit stricte par clé API, ne renvoie que la sortie textuelle nécessaire au cas d'usage métier, et journalise les volumes de requêtes pour détecter un pattern d'extraction anormal.
