# Weak Password Policy

Le handler `register` n'exigeait qu'une longueur minimale de 6 caractères sans aucune vérification contre les mots de passe compromis connus, facilitant le brute force et le credential stuffing (CWE-521). La version corrigée impose une longueur minimale de 12 caractères conforme aux recommandations NIST 800-63B et ajoute une vérification `is_password_compromised` contre une liste de fuites connues, sans imposer de règles de composition artificielles.
