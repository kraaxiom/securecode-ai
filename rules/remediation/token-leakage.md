# Remédiation — Token Leakage

## Principe
Ne jamais exposer une clé d'API de fournisseur LLM côté client ; toujours proxifier les appels via un backend qui détient le secret. Masquer les jetons dans les journaux et mettre en place une rotation régulière avec scoping par environnement.

## JavaScript (client — Avant/Après)
```js
// Avant — vulnérable : clé API du fournisseur LLM codée en dur côté client (bundle JS)
const client = new AnthropicClient({ apiKey: "sk-ant-api03-XXXXXXXXXXXXXXXX" });
async function askAssistant(question) {
  return client.messages.create({ model: "claude-3", messages: [{ role: "user", content: question }] });
}

// Après — sécurisé : appel proxifié via le backend, aucun secret côté client
async function askAssistant(question) {
  return fetch("/api/assistant", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ question }),
  }).then(r => r.json());
}
```

## Node.js (backend — proxy)
```js
// Après — sécurisé : le backend détient seul la clé, jamais transmise au client
const client = new AnthropicClient({ apiKey: process.env.LLM_API_KEY }); // via secret manager, jamais en dur

app.post("/api/assistant", requireAuth, async (req, res) => {
  const response = await client.messages.create({ model: "claude-3", messages: [{ role: "user", content: req.body.question }] });
  res.json(response);
});

// Masquage des jetons dans les journaux
logger.use(redactHeaders(["authorization", "x-api-key"]));
```

## PHP (backend — proxy)
```php
// Avant — vulnérable : clé exposée dans une réponse de debug
$response = $client->post('/v1/messages', ['headers' => ['Authorization' => "Bearer $apiKey"]]);
error_log(print_r($response->getHeaders(), true)); // journalise l'en-tête d'autorisation en clair

// Après — sécurisé : clé résolue depuis un secret manager, en-têtes masqués dans les logs
$apiKey = SecretManager::resolve('llm_api_key');
$response = $client->post('/v1/messages', ['headers' => ['Authorization' => "Bearer $apiKey"]]);
error_log(print_r(redact_sensitive_headers($response->getHeaders()), true));
```

## Checklist de vérification post-patch
- [ ] Aucune clé d'API de fournisseur LLM n'est présente dans le code exécuté côté client (bundle JS, app mobile).
- [ ] Tous les appels au LLM passent par un backend qui détient seul le secret, avec authentification de l'appelant.
- [ ] Les journaux applicatifs masquent systématiquement les en-têtes d'autorisation et jetons.
- [ ] Une rotation régulière des clés est en place avec scoping distinct par environnement (dev/staging/prod).
