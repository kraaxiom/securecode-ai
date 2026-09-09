# Fixture — python-django-demo

Mini-projet Python/Django volontairement vulnérable, utilisé pour valider le pipeline complet du skill (scan → détection → explication → patch → test → score → rapport) décrit dans `docs/Workflow.md`.

**Ne jamais déployer ce code — il contient des vulnérabilités intentionnelles.**

## Contenu

- `app/views.py` — vues exposant trois vulnérabilités intentionnelles :
  - SSTI — `knowledge/injections/ssti.md`
  - Path Traversal — `knowledge/file-inclusion/path-traversal.md`
  - Mass Assignment — `knowledge/authorization/mass-assignment.md`

## Utilisation attendue

1. Scanner `app/views.py` avec `rules/sast/python/ssti.yaml`, `rules/sast/python/path-traversal.yaml`, `rules/sast/python/mass-assignment.yaml`.
2. Comparer les findings obtenus à `tests/eval-set.json` (cas `python-django-demo-*`).
3. Appliquer les patchs via les fichiers `rules/remediation/` correspondants.
4. Vérifier que le fichier corrigé ne déclenche plus les règles correspondantes.
