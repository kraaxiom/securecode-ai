# Fixture — node-express-demo

Mini-projet Node.js/Express volontairement vulnérable, utilisé pour valider le pipeline complet du skill (scan → détection → explication → patch → test → score → rapport) décrit dans `docs/Workflow.md`.

**Ne jamais déployer ce code — il contient des vulnérabilités intentionnelles.**

## Contenu

- `routes/users.js` — routes exposant trois vulnérabilités intentionnelles :
  - Command Injection — `knowledge/injections/command-injection.md`
  - Broken Access Control — `knowledge/authorization/broken-access-control.md`
  - Weak JWT secret — `knowledge/crypto/weak-jwt-secret.md`

## Utilisation attendue

1. Scanner `routes/users.js` avec `rules/sast/js/command-injection.yaml`, `rules/sast/js/broken-access-control.yaml`, `rules/sast/js/weak-jwt-secret.yaml`.
2. Comparer les findings obtenus à `tests/eval-set.json` (cas `node-express-demo-*`).
3. Appliquer les patchs via les fichiers `rules/remediation/` correspondants.
4. Vérifier que le fichier corrigé ne déclenche plus les règles correspondantes.
