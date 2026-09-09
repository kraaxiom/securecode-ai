---
id: coupon-abuse
category: business-logic
cwe: CWE-840
owasp: A04:2021-Insecure-Design
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Coupon Abuse (abus de codes promotionnels)

## Description
L'abus de coupons regroupe les failles de logique permettant à un utilisateur d'appliquer un code promotionnel plus de fois que prévu, de le combiner avec d'autres offres non cumulables, ou d'en générer/deviner de nouveaux non destinés à son compte. Ces failles proviennent généralement d'une validation insuffisante côté serveur sur l'unicité d'utilisation, le périmètre d'éligibilité, ou la prévisibilité du format des codes.

## Où ça apparaît typiquement
- Vérification d'usage unique d'un coupon effectuée côté client ou reposant uniquement sur un cookie/localStorage.
- Génération de codes promotionnels suivant un format prévisible (séquentiel, dérivé d'un identifiant utilisateur) permettant l'énumération.
- Absence de vérification que le coupon est bien lié au compte, à la commande ou à la période pour laquelle il a été émis.
- Cumul non contrôlé de plusieurs coupons ou codes de parrainage sur une même commande.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Marquage "coupon utilisé" effectué côté client (cookie, état local) sans enregistrement serveur associé au compte utilisateur.
- Codes promotionnels générés de façon prévisible ou séquentielle sans entropie suffisante pour empêcher l'énumération.
- Absence de contrainte d'unicité en base de données liant un coupon à un compte/une commande précis lors de son application.
- Endpoint d'application de coupon sans vérification de la période de validité ou du périmètre d'éligibilité (produit, montant minimum).

## Remédiation
- Enregistrer et vérifier l'usage des coupons côté serveur, avec contrainte d'unicité en base liée au compte utilisateur.
- Générer les codes promotionnels avec une entropie suffisante pour empêcher toute énumération ou prédiction.
- Vérifier systématiquement côté serveur le périmètre d'éligibilité (dates, montant, produits, cumul autorisé) à chaque application.
- Limiter le nombre de tentatives d'application de coupon par session/compte pour freiner les attaques par force brute de codes.
- Voir `rules/remediation/coupon-abuse.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/coupon-abuse/`.

## Références
- OWASP Top 10: A04:2021 – Insecure Design
- CWE-840: Business Logic Errors
