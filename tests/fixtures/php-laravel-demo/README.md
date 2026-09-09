# Fixture — php-laravel-demo

Mini-projet PHP/Laravel volontairement vulnérable, utilisé pour valider le pipeline complet du skill (scan → détection → explication → patch → test → score → rapport) décrit dans `docs/Workflow.md`.

**Ne jamais déployer ce code — il contient des vulnérabilités intentionnelles.**

## Contenu

- `app/Http/Controllers/UserController.php` — contrôleur exposant trois vulnérabilités intentionnelles :
  - SQL Injection (UNION-based) — `knowledge/injections/sqli-union.md`
  - IDOR — `knowledge/authorization/idor.md`
  - Reflected XSS — `knowledge/xss/reflected-xss.md`

## Utilisation attendue

1. Scanner `app/Http/Controllers/UserController.php` avec les règles `rules/sast/php/sqli-union.yaml`, `rules/sast/php/idor.yaml`, `rules/sast/php/reflected-xss.yaml`.
2. Comparer les findings obtenus à `tests/eval-set.json` (cas `php-laravel-demo-*`).
3. Appliquer les patchs via `rules/remediation/sqli-union.md`, `rules/remediation/idor.md`, `rules/remediation/reflected-xss.md`.
4. Vérifier que le fichier corrigé ne déclenche plus les règles correspondantes.
