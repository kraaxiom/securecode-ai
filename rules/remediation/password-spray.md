# Remédiation — Password Spraying

## Principe
Compléter la limitation par compte par une agrégation globale des échecs (par IP/source, par fenêtre de temps, tous comptes confondus) et déployer la MFA, qui neutralise l'essentiel de l'impact d'un spray réussi.

## PHP (Laravel)
```php
// Avant — vulnérable : limitation uniquement par compte
$key = 'login:' . strtolower($request->input('email'));
if (RateLimiter::tooManyAttempts($key, 5)) { /* ... */ }

// Après — sécurisé : agrégation globale par IP en plus de la limite par compte
$accountKey = 'login:account:' . strtolower($request->input('email'));
$ipKey = 'login:ip:' . $request->ip();

if (RateLimiter::tooManyAttempts($accountKey, 5) || RateLimiter::tooManyAttempts($ipKey, 50)) {
    SecurityAlert::dispatchIf(
        RateLimiter::attempts($ipKey) > 50,
        'password_spray_suspected',
        $request->ip()
    );
    return back()->withErrors(['email' => 'Trop de tentatives.']);
}
RateLimiter::hit($accountKey, 900);
RateLimiter::hit($ipKey, 900);
```

## JS / Node.js (Express)
```js
// Avant — vulnérable : compteur uniquement par compte, un spray distribué passe inaperçu
const attempts = await getFailedAttempts(req.body.email);
if (attempts > 5) return res.status(429).json({ error: 'Trop de tentatives' });

// Après — sécurisé : agrégation globale par IP + politique de mot de passe robuste
const [accountAttempts, ipAttempts] = await Promise.all([
  getFailedAttempts(req.body.email),
  getFailedAttemptsByIp(req.ip),
]);
if (accountAttempts > 5 || ipAttempts > 50) {
  if (ipAttempts > 50) await flagSuspiciousSource(req.ip, 'password_spray');
  return res.status(429).json({ error: 'Trop de tentatives.' });
}
```

## Python (Django)
```python
# Avant — vulnérable : agrégation par compte uniquement
if get_failed_attempts(email) > 5:
    return JsonResponse({"error": "Trop de tentatives"}, status=429)

# Après — sécurisé : agrégation globale par IP en complément
account_attempts = get_failed_attempts(email)
ip_attempts = get_failed_attempts_by_ip(get_client_ip(request))

if account_attempts > 5 or ip_attempts > 50:
    if ip_attempts > 50:
        flag_suspicious_source(get_client_ip(request), reason="password_spray")
    return JsonResponse({"error": "Trop de tentatives."}, status=429)
```

## Checklist de vérification post-patch
- [ ] Une agrégation globale (par IP/source, tous comptes confondus) complète la limitation par compte.
- [ ] Une politique de mot de passe interdit les valeurs les plus courantes ciblées par le spray.
- [ ] La MFA est déployée sur les comptes à privilèges.
- [ ] Un test confirme qu'un volume élevé d'échecs répartis sur des comptes distincts depuis une même source déclenche une alerte.
- [ ] Les alertes de sécurité liées au password spraying sont routées vers la supervision.
