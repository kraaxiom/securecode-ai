"""
Exemple CORRIGE — Django settings.py : secrets chargés depuis l'environnement.

Fix pour CWE-798 (Use of Hard-coded Credentials) :
- Aucun secret n'est codé en dur ni versionné dans le dépôt.
- Toutes les valeurs sensibles sont chargées depuis des variables
  d'environnement (elles-mêmes injectées par un gestionnaire de secrets :
  AWS Secrets Manager, Azure Key Vault, GCP Secret Manager, ou Vault) au
  démarrage du processus, jamais stockées en clair dans un fichier
  versionné.
"""

import os

# settings.py

# CORRIGE : SECRET_KEY injectée via variable d'environnement, différente
# par environnement, jamais présente dans le code source.
SECRET_KEY = os.environ["DJANGO_SECRET_KEY"]

DEBUG = os.environ.get("DJANGO_DEBUG", "False") == "True"

# CORRIGE : identifiants de base de données chargés depuis l'environnement.
# En production, ces variables sont injectées par le gestionnaire de
# secrets de la plateforme (ex: AWS ECS task secrets, Kubernetes Secret).
DATABASES = {
    "default": {
        "ENGINE": "django.db.backends.postgresql",
        "NAME": os.environ["DB_NAME"],
        "USER": os.environ["DB_USER"],
        "PASSWORD": os.environ["DB_PASSWORD"],
        "HOST": os.environ["DB_HOST"],
        "PORT": os.environ.get("DB_PORT", "5432"),
    }
}

# CORRIGE : clés AWS jamais codées en dur. En production, préférer un rôle
# IAM (instance profile / IRSA) sans clé statique du tout ; à défaut,
# charger depuis l'environnement.
AWS_ACCESS_KEY_ID = os.environ.get("AWS_ACCESS_KEY_ID")
AWS_SECRET_ACCESS_KEY = os.environ.get("AWS_SECRET_ACCESS_KEY")

# CORRIGE : clé Stripe récupérée via le gestionnaire de secrets au
# démarrage, jamais stockée en clair dans un fichier versionné.
STRIPE_SECRET_KEY = os.environ["STRIPE_SECRET_KEY"]

# Rappel opérationnel (voir README) : si un secret a déjà été committé,
# le supprimer du code ne suffit pas — il doit être révoqué/régénéré
# auprès du fournisseur et l'historique Git purgé (git filter-repo/BFG).
