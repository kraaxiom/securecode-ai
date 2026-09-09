# Remédiation — Model Extraction

## Principe
Appliquer des quotas et une limitation de débit stricts par client/clé API sur les endpoints d'inférence, ne renvoyer que les informations nécessaires au cas d'usage (éviter les logits bruts), et surveiller les patterns de requêtes automatisées.

## Python
```python
# Avant — vulnérable : API sans limite de débit, expose les logits complets
@app.post("/v1/infer")
def infer(request):
    result = model.predict(request.json["input"], return_logits=True)
    return jsonify({"output": result.text, "logits": result.logits.tolist()})

# Après — sécurisé : quota par clé API, sortie minimale, détection d'usage anormal
@app.post("/v1/infer")
@rate_limit(key_func=lambda r: r.headers["X-API-Key"], limit="100/hour")
def infer(request):
    api_key = request.headers["X-API-Key"]
    if usage_monitor.detects_extraction_pattern(api_key):  # volume/régularité anormale
        audit_log.record("suspected_model_extraction", api_key)
        abort(429, "Usage anormal détecté")

    result = model.predict(request.json["input"], return_logits=False)
    return jsonify({"output": result.text})  # pas de logits/probabilités bruts exposés
```

## JavaScript / Node.js
```js
// Avant — vulnérable
app.post("/v1/infer", async (req, res) => {
  const result = await model.predict(req.body.input, { returnLogits: true });
  res.json({ output: result.text, logits: result.logits });
});

// Après — sécurisé
app.post("/v1/infer", rateLimit({ keyGenerator: (req) => req.headers["x-api-key"], max: 100, windowMs: 3600_000 }), async (req, res) => {
  const apiKey = req.headers["x-api-key"];
  if (await usageMonitor.detectsExtractionPattern(apiKey)) {
    auditLog.record("suspected_model_extraction", apiKey);
    return res.status(429).json({ error: "Usage anormal détecté" });
  }
  const result = await model.predict(req.body.input, { returnLogits: false });
  res.json({ output: result.text });
});
```

## Checklist de vérification post-patch
- [ ] Une limitation de débit et un quota stricts sont appliqués par clé API/utilisateur sur les endpoints d'inférence.
- [ ] Les logits/probabilités bruts ne sont exposés que si strictement nécessaires au cas d'usage métier.
- [ ] Un système de surveillance détecte les volumes de requêtes systématiques ou les patterns non humains.
- [ ] L'usage autorisé de l'API est encadré contractuellement (CGU/CLUF) pour les modèles à forte valeur.
