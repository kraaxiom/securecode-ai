# FAQ

## Le skill corrige-t-il automatiquement mon code ?

Cela dépend strictement du mode dans lequel l'agent se trouve (voir `SKILL.md` et `docs/Workflow.md`) :

- **En Mode Audit** (par défaut hors session de développement, ou sur demande explicite d'audit) : **non, jamais automatiquement.** Le skill produit un rapport de findings et de patchs proposés, puis demande explicitement : *"Je peux corriger les vulnérabilités détectées dans le code source du projet — dois-je procéder ?"*. Sans confirmation positive de l'utilisateur, aucun fichier n'est modifié.
- **En Mode Développement assisté** (actif automatiquement pendant qu'un agent écrit ou modifie du code) : le skill corrige **immédiatement et sans redemander confirmation** uniquement dans un cas précis — un finding de confiance `high`, sur du code que l'agent vient lui-même d'écrire dans la session en cours. Pour du code préexistant du projet, ou pour une confiance `low`/`medium`, le principe de confirmation du Mode Audit s'applique à nouveau.

Dans les deux modes, un patch appliqué est toujours accompagné d'un test de non-régression généré, et l'utilisateur est informé de ce qui a été corrigé (fichier, ligne, type de vulnérabilité).

## Ce skill peut-il générer des exploits ?

**Non.** La portée du skill est strictement défensive et cette limite est structurelle, pas seulement une consigne de comportement :

- Le contenu de `knowledge/` documente des **patterns de reconnaissance** (comment identifier une vulnérabilité dans du code source) et de **remédiation**, jamais de payload d'exploitation prêt à l'emploi ni de séquence d'attaque contre un système réel.
- Les prompts (`prompts/detect.md`, etc.) contraignent explicitement l'agent à ne produire que des findings basés sur des règles chargées, jamais de conseil d'exploitation.
- Le skill ne génère pas de scripts de scan de masse contre des cibles tierces, ni de techniques de contournement de protections (WAF bypass, EDR evasion, anti-forensic).
- Toute demande qui sortirait de ce cadre (ex: "génère un exploit pour cette CVE contre ce serveur en production") doit être redirigée vers l'usage prévu : détection et correction sur le code source de l'utilisateur, jamais vers une cible tierce.

Cette limite protège à la fois la responsabilité légale de l'auteur du skill et son acceptabilité sur des marketplaces professionnelles (`SPEC.md` section 10) — ce n'est pas une restriction arbitraire, c'est la définition même du produit : un auditeur/correcteur, pas un outil offensif.

## Ce skill remplace-t-il un pentest humain ?

Non. `SKILL.md` le rappelle explicitement : ce skill "ne remplace pas un pentest humain, mais réduit fortement l'introduction et la persistance de vulnérabilités connues". Un pentest humain couvre des dimensions que l'analyse statique/patterns ne couvre pas nativement : logique métier complexe spécifique au produit, chaînes d'exploitation combinant plusieurs faiblesses mineures, ingénierie sociale, tests dynamiques en conditions réelles avec autorisation. Le skill est un outil d'aide continue, pas un audit de sécurité complet et certifié.

## Quel est le taux de faux positifs / faux négatifs ?

Il n'existe pas de chiffre universel — cela dépend du langage, du type de règle (`ast-pattern` est généralement plus précis que `regex`, voir `docs/Scanner.md`), et de la spécificité du code du projet. Le skill limite structurellement les faux positifs via les clauses `exclude_if` de chaque règle et l'attribution d'une `confidence` explicite (`low|medium|high`) plutôt qu'une détection binaire — un finding `confidence: low` doit être lu comme une piste à vérifier, pas une certitude. La précision réelle sur un projet donné se mesure et s'améliore dans le temps via le suivi des statuts (`docs/Database.md`) : un `false_positive` marqué explicitement alimente la fiabilité perçue des scans suivants sur ce projet, même si cela n'entraîne pas d'apprentissage automatique global des règles elles-mêmes.

## Que se passe-t-il si je refuse la correction proposée en Mode Audit ?

Le rapport reste la seule sortie produite. Aucun fichier n'est modifié. Les findings restent `status: open` dans le store (`docs/Database.md`) et seront à nouveau signalés lors d'un scan ultérieur, jusqu'à ce qu'ils soient corrigés, ou explicitement marqués `false_positive`/`accepted_risk`.

## Le skill couvre-t-il tous les langages et frameworks possibles ?

Non — la couverture initiale prioritaire est PHP/Laravel/Symfony, Node.js, Python, Java, C#, Go, Rust (`SKILL.md`). De nouveaux langages/frameworks peuvent être ajoutés sans modifier le moteur (voir `docs/Plugins.md`) ; l'absence de règle pour un langage donné signifie une absence de détection sur ce langage, pas une garantie d'absence de vulnérabilité.

## Le skill détecte-t-il des failles "zero-day" ?

Il propose un **mode de revue heuristique complémentaire** (`docs/ZeroDay.md`), distinct du scan par signature qui constitue le cœur du skill. La différence est structurelle : `rules/sast/` détecte des patterns déjà catalogués (signature connue), tandis que le mode zero-day fait raisonner l'agent sur le code à la recherche de patterns dangereux **sans règle existante** — logique métier défaillante, combinaison dangereuse de primitives, déviation par rapport aux conventions du projet.

Ce mode n'est jamais automatique : il s'active uniquement sur demande explicite, après un scan par signature déjà réalisé. Chaque finding doit être **reproduit localement** (jamais contre un système tiers ou en production) avant d'être remonté, et consigné dans `<projet>/.securecode/zeroday-registry.json` avec un statut initial `unconfirmed` — jamais `confirmed`/`false_positive`/`accepted_risk`, qui restent des décisions humaines exclusives (`prompts/zeroday-verify.md`).

Le terme "zero-day" est utilisé ici au sens "pattern non catalogué découvert avant qu'une règle formelle n'existe", **pas** au sens "vulnérabilité activement exploitée dans la nature sur un logiciel tiers" — le skill ne fait aucune recherche de CVE non divulguées sur des produits tiers, ce qui sortirait de son cadre strictement défensif.

## Le skill détecte-t-il des failles "zero-day" ?

Il dispose d'un **Mode Recherche zero-day** distinct (`docs/ZeroDay.md`, `prompts/zeroday-detect.md`) qui effectue une revue heuristique pour repérer des patterns dangereux **sans règle `rules/sast/` existante** — logique métier défaillante, combinaisons dangereuses de primitives, déviations par rapport aux conventions du projet. Ce n'est ni une recherche de CVE non divulguées sur des logiciels tiers, ni un outil de fuzzing/symbolic execution : c'est un raisonnement assisté sur le code source du projet de l'utilisateur, avec deux garanties structurelles :

- **Reproductibilité obligatoire** : un finding zero-day n'est remonté qu'après une tentative de reproduction locale réussie (ou explicitement documentée comme non réussie, avec `confidence: low`) — jamais sur la seule base d'un raisonnement théorique.
- **Statuts de confirmation/rejet réservés à l'humain** : l'agent peut poser `unconfirmed`, jamais `confirmed`/`false_positive`/`accepted_risk` — ces décisions appartiennent à un expert cybersécurité nommé, via `prompts/zeroday-verify.md`.

Ce mode ne s'active jamais automatiquement, seulement sur demande explicite, en complément d'un audit standard — voir `SKILL.md` section "Mode Recherche zero-day".

## Le skill a-t-il une portée légale/éthique particulière ?

Le skill est conçu pour être utilisé sur du code dont l'utilisateur a la responsabilité ou l'autorisation légitime (son propre projet, un projet client dans le cadre d'une mission autorisée). Il ne doit pas être détourné pour scanner ou tenter de corriger du code appartenant à un tiers sans autorisation — ce cas de figure sort du cadre "auditeur + correcteur de sécurité applicative" défini par `SKILL.md` et relève de la responsabilité de l'utilisateur, pas d'une fonctionnalité du skill.
