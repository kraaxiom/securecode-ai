# Remédiation — Model Inversion

## Principe
Appliquer des techniques de confidentialité différentielle ou d'anonymisation avant tout entraînement sur données sensibles, réaliser des tests de mémorisation avant mise en production, et limiter/surveiller les volumes de requêtes permettant une reconstruction progressive.

## Python
```python
# Avant — vulnérable : fine-tuning sur données sensibles brutes, aucune limite de requêtes
def fine_tune(model, sensitive_dataset):
    return model.fit(sensitive_dataset)  # pas d'anonymisation ni de confidentialité différentielle

@app.post("/v1/infer")
def infer(request):
    return model.predict(request.json["input"])  # aucune limite de volume de requêtes

# Après — sécurisé : anonymisation + DP-SGD + test de mémorisation + limitation de requêtes
def fine_tune(model, sensitive_dataset):
    anonymized = anonymize_pii(sensitive_dataset)
    trained = model.fit(anonymized, optimizer=dp_sgd_optimizer(noise_multiplier=1.1))
    memorization_score = run_memorization_test(trained, sensitive_dataset)
    if memorization_score > MAX_ACCEPTABLE_MEMORIZATION:
        raise ModelMemorizationTooHigh(memorization_score)
    return trained

@app.post("/v1/infer")
@rate_limit(key_func=lambda r: r.headers["X-API-Key"], limit="50/hour")
def infer(request):
    api_key = request.headers["X-API-Key"]
    if usage_monitor.detects_reconstruction_pattern(api_key):
        audit_log.record("suspected_model_inversion", api_key)
        abort(429, "Usage anormal détecté")
    return model.predict(request.json["input"])
```

## JavaScript / Node.js
```js
// Avant — vulnérable
app.post("/v1/infer", async (req, res) => {
  res.json(await model.predict(req.body.input)); // aucune limite de volume
});

// Après — sécurisé
app.post("/v1/infer", rateLimit({ keyGenerator: (req) => req.headers["x-api-key"], max: 50, windowMs: 3600_000 }), async (req, res) => {
  const apiKey = req.headers["x-api-key"];
  if (await usageMonitor.detectsReconstructionPattern(apiKey)) {
    auditLog.record("suspected_model_inversion", apiKey);
    return res.status(429).json({ error: "Usage anormal détecté" });
  }
  res.json(await model.predict(req.body.input));
});
```

## Checklist de vérification post-patch
- [ ] Les données sensibles utilisées en entraînement/fine-tuning sont anonymisées et/ou entraînées avec confidentialité différentielle.
- [ ] Un test de mémorisation est exécuté avant chaque mise en production d'un modèle entraîné sur données propriétaires.
- [ ] Le volume et la structure des requêtes d'inférence sont limités et surveillés pour détecter une reconstruction progressive.
- [ ] L'usage de données synthétiques est privilégié quand c'est possible, au lieu de données réelles sensibles.
