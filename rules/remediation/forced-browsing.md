# Remédiation — Forced Browsing

## Principe
Appliquer un contrôle d'authentification/autorisation explicite sur toute route ou ressource sensible, indépendamment de sa visibilité dans l'interface, et ne jamais stocker de fichiers sensibles dans un répertoire servi publiquement.

## PHP (Laravel)
```php
// Avant — vulnérable : page d'admin non liée dans l'UI mais accessible sans contrôle
Route::get('/internal-reports', [ReportController::class, 'index']);

// Après — sécurisé : middleware d'autorisation explicite, indépendant de la découvrabilité
Route::get('/internal-reports', [ReportController::class, 'index'])
    ->middleware(['auth', 'can:viewInternalReports,App\Models\User']);

// Fichiers sensibles : jamais dans public/, servis via un contrôleur qui vérifie l'autorisation
Route::get('/exports/{file}', [ExportController::class, 'download'])->middleware('auth');
```

## JS / Node.js (Express)
```js
// Avant — vulnérable : fichiers de sauvegarde exposés statiquement, sans lien direct mais accessibles
app.use('/backups', express.static('backups/')); // aucune protection

// Après — sécurisé : contrôle d'accès sur chaque fichier servi, jamais de répertoire sensible en static
app.get('/backups/:file', requireAuth, requireRole('admin'), async (req, res) => {
  const filePath = resolveSafePath('backups', req.params.file); // évite aussi la traversée de chemin
  res.sendFile(filePath);
});
```

## Python (Django)
```python
# Avant — vulnérable : étape intermédiaire d'un flux accessible en sautant les étapes précédentes
def payment_confirmation(request):
    return render(request, "confirmation.html")  # pas de vérification de l'état du flux

# Après — sécurisé : vérification côté serveur de l'état d'avancement réel
def payment_confirmation(request):
    checkout = request.session.get("checkout")
    if not checkout or checkout.get("status") != "payment_authorized":
        return redirect("checkout_start")  # refuse l'accès direct à l'étape finale
    return render(request, "confirmation.html")
```

## Checklist de vérification post-patch
- [ ] Toute route ou ressource sensible impose un contrôle d'authentification/autorisation, indépendamment de sa visibilité dans l'UI.
- [ ] Aucun fichier sensible (sauvegarde, log, export, config) n'est servi depuis un répertoire statique public.
- [ ] Les flux multi-étapes vérifient côté serveur que les étapes précédentes ont été réellement complétées.
- [ ] Un test confirme qu'une route non liée dans l'UI reste protégée par le même contrôle que les routes visibles.
- [ ] Un audit périodique recherche les fichiers/répertoires accidentellement exposés publiquement.
