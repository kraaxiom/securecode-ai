"""
Exemple VULNERABLE — Django settings.py : secrets codés en dur.

CWE-798: Use of Hard-coded Credentials
Les clés API, la SECRET_KEY Django et les identifiants de base de données
sont écrits en clair dans le code source, versionnés dans le dépôt Git et
donc visibles par quiconque a accès au repository (y compris dans
l'historique, même après suppression ultérieure).
"""

# settings.py

# VULNERABLE : clé secrète Django codée en dur, identique en dev et en prod.
SECRET_KEY = "django-insecure-8x!k2p9qz3w7e1r4t6y8u0i2o4p6a8s0d2f4g6h8j0k2l4"

DEBUG = True

# VULNERABLE : identifiants de base de données en clair dans le code versionné.
DATABASES = {
    "default": {
        "ENGINE": "django.db.backends.postgresql",
        "NAME": "prod_db",
        "USER": "admin",
        "PASSWORD": "Sup3rS3cret!2024",
        "HOST": "db.example.com",
        "PORT": "5432",
    }
}

# VULNERABLE : clés cloud AWS codées en dur, utilisées pour boto3 ailleurs
# dans le projet. Si ce fichier fuite (repo public, capture d'écran,
# fork accidentel), les clés permettent un accès complet au compte AWS.
AWS_ACCESS_KEY_ID = "AKIAIOSFODNN7EXAMPLEKEY"
AWS_SECRET_ACCESS_KEY = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"

# VULNERABLE : clé API Stripe de production codée en dur.
STRIPE_SECRET_KEY = "sk_live_EXAMPLE_NOT_A_REAL_KEY"
