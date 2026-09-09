# Jailbreak de modèle (CWE-1427)

La version vulnérable expose un endpoint de chat qui transmet directement les messages de l'utilisateur au modèle avec pour seule barrière le system prompt, sans couche de modération indépendante en entrée ou en sortie, ni surveillance du nombre de tours suspects dans une session. La correction ajoute un classifieur de sécurité indépendant du modèle principal qui filtre à la fois l'entrée et la sortie, limite le nombre de tentatives suspectes par session, et ne considère jamais le system prompt comme l'unique contrôle de sécurité.
