# Model Inversion (CWE-200)

La version vulnérable renvoie, sur un modèle fine-tuné sur des données internes sensibles, des scores de confiance bruts et des sorties très détaillées sans aucune limite sur le volume de requêtes, ce qui permet une reconstruction progressive de données mémorisées lors de l'entraînement. La correction restreint la granularité des sorties, ajoute du bruit (confidentialité différentielle) et n'expose que des résultats agrégés, tout en limitant le nombre de requêtes par client.
