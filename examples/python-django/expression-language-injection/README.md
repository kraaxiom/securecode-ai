# Expression Language Injection

La version vulnérable concatène l'entrée utilisateur `nom` directement dans la source du template Django avant compilation via `Template()`, ce qui correspond à CWE-917 : l'attaquant peut injecter des tags ou variables de template (`{% ... %}`, `{{ ... }}`) évalués côté serveur. La version corrigée fixe le template comme chaîne statique et transmet `nom` uniquement en tant que variable du contexte de rendu, jamais comme fragment de code de template.
