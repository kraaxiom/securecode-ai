# Moteur de règles (Engine)

Ce document couvre ce qui se passe **après** que le scanner (`docs/Scanner.md`) a produit une liste brute de findings : chargement/priorité des règles, gestion des faux positifs, corrélation et dédoublonnage. C'est l'étape 3-4 du pipeline (`docs/Workflow.md`).

## 1. Chargement et priorité des règles

- Les règles sont chargées par langage, puis filtrées par pertinence au fichier examiné (`docs/AI.md`). Il n'y a pas de notion de priorité entre catégories différentes (une règle `injections/` n'est pas "plus prioritaire" qu'une règle `auth/`) : toutes les règles applicables à un fichier sont évaluées.
- En revanche, **au sein d'une même zone de code, une règle `ast-pattern` est toujours préférée à une règle `regex` couvrant le même slug** si les deux existent pour le langage, car l'AST réduit le risque de faux positif. Si une règle `regex` de secours existe pour un langage sans support AST, elle s'applique seule.
- Si deux versions d'une même règle existent (mise à jour de `rules/sast/<lang>/<slug>.yaml`), la version la plus récente du fichier fait foi — le moteur ne conserve pas d'historique de versions de règles au-delà de ce que gère le contrôle de version du projet lui-même.

## 2. Gestion des faux positifs

Deux mécanismes distincts, à ne pas confondre :

1. **`exclude_if` (défini dans la règle)** — condition qui, vérifiée, empêche l'émission du finding ou le déclasse en `info`. Voir `docs/Scanner.md` section 5.
2. **`status: false_positive` (défini a posteriori sur un finding existant)** — un finding déjà émis, revu par un humain ou par l'agent en Mode Audit, et jugé non exploitable dans le contexte réel du projet. Ce statut est persisté dans le findings store (`docs/Database.md`) et doit être respecté par les scans suivants : un finding déjà marqué `false_positive` pour un couple `rule_id + file + line` ne doit pas être ré-remonté comme `open` à l'identique lors d'un scan ultérieur, sauf si le code à cet endroit a changé.

Un finding ne doit jamais être supprimé silencieusement du rapport pour "nettoyer" — le passage à `false_positive` ou `accepted_risk` doit rester traçable dans le store, contrairement à une simple omission.

## 3. Corrélation et dédoublonnage

Avant l'étape d'explication, les findings bruts issus du scan sont traités comme suit :

- **Dédoublonnage strict** : deux findings identiques sur `rule_id + file + line_start` (ex: deux passes de scan qui se chevauchent) sont fusionnés en un seul.
- **Regroupement par fichier** : les findings d'un même fichier sont groupés pour permettre à l'étape d'explication de charger le contexte du fichier une seule fois plutôt que par finding.
- **Chevauchement de règles** : quand deux règles différentes matchent la même zone de code pour des raisons proches (ex: une règle générique "injection SQL" et une règle spécifique "UNION-based SQLi" matchent la même ligne), le moteur conserve les deux findings s'ils portent des `cwe`/`rule_id` distincts et informatifs, mais les présente groupés dans le rapport plutôt que comme deux alertes indépendantes et redondantes pour le lecteur. Si les deux règles sont strictement redondantes (même `cwe`, même portée), ne garder que la règle la plus spécifique.

## 4. Clé de corrélation dans le temps

Pour permettre le suivi d'un finding entre deux scans successifs du même projet (nécessaire à `docs/Database.md`), la clé de corrélation est `rule_id + file + line_start` (approximative : un léger décalage de ligne dû à une modification en amont dans le fichier doit être toléré par une correspondance approchée, pas seulement une égalité stricte de numéro de ligne). Le champ `id` du finding, lui, est un identifiant unique par scan et ne doit pas être utilisé comme clé de suivi dans le temps.

## 5. Interaction avec les modes du skill

Le moteur ne décide jamais lui-même d'appliquer un patch : il ne fait que produire une liste de findings triés, dédoublonnés, avec un statut de faux positif éventuel déjà appliqué. La décision d'agir sur ces findings (correction immédiate ou attente de confirmation) appartient exclusivement à la couche pipeline (`docs/Workflow.md`), selon le mode opératoire actif.
