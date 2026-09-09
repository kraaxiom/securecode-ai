# Remédiation — Secret Leakage via LLM

## Principe
Ne jamais inclure de secrets en clair dans un prompt système, filtrer la sortie du modèle pour détecter des motifs de secrets avant tout renvoi, et restreindre strictement l'accès des agents aux fichiers/variables sensibles.

## Python
```python
# Avant — vulnérable : clé API en clair dans le system prompt, accès fichier non restreint
SYSTEM_PROMPT = f"Tu es un assistant interne. Clé API interne : {INTERNAL_API_KEY}. Utilise-la pour appeler le service X."

def read_config_for_agent(path):
    return open(path).read()  # accès libre, y compris fichiers de config sensibles

# Après — sécurisé : référence indirecte résolue côté application, scope restreint, filtrage de sortie
SYSTEM_PROMPT = "Tu es un assistant interne. Pour appeler le service X, utilise l'outil call_service_x (le secret est géré par l'application, jamais exposé dans ce contexte)."

ALLOWED_CONFIG_DIRS = {"/app/public-config"}

def read_config_for_agent(path):
    if not is_within_allowed_dirs(path, ALLOWED_CONFIG_DIRS):
        audit_log.record("blocked_sensitive_file_access", path)
        raise PermissionError("Accès refusé à ce chemin")
    return open(path).read()

def call_service_x(**kwargs):
    return http_client.post(SERVICE_X_URL, headers={"Authorization": f"Bearer {resolve_secret('service_x')}"}, json=kwargs)

def get_llm_response(messages):
    response = llm.chat(messages=messages)
    if secret_scanner.contains_secret(response.content):  # filtrage de sortie
        audit_log.record("blocked_secret_in_output", response.id)
        return REDACTED_RESPONSE
    return response
```

## JavaScript / Node.js
```js
// Avant — vulnérable
const SYSTEM_PROMPT = `Tu es un assistant interne. Clé API interne : ${INTERNAL_API_KEY}.`;

// Après — sécurisé
const SYSTEM_PROMPT = "Tu es un assistant interne. Utilise l'outil callServiceX ; le secret est géré côté application.";

const ALLOWED_CONFIG_DIRS = ["/app/public-config"];

function readConfigForAgent(path) {
  if (!isWithinAllowedDirs(path, ALLOWED_CONFIG_DIRS)) {
    auditLog.record("blocked_sensitive_file_access", path);
    throw new Error("Accès refusé à ce chemin");
  }
  return fs.readFileSync(path, "utf8");
}

async function callServiceX(args) {
  return httpClient.post(SERVICE_X_URL, args, { headers: { Authorization: `Bearer ${resolveSecret("service_x")}` } });
}

async function getLlmResponse(messages) {
  const response = await llm.chat({ messages });
  if (secretScanner.containsSecret(response.content)) {
    auditLog.record("blocked_secret_in_output", response.id);
    return REDACTED_RESPONSE;
  }
  return response;
}
```

## Checklist de vérification post-patch
- [ ] Aucun secret en clair n'apparaît dans les templates de prompt système versionnés.
- [ ] Un filtrage de sortie (output scanning) détecte les motifs de secrets avant tout renvoi de réponse du modèle.
- [ ] L'accès de l'agent au système de fichiers/variables d'environnement est restreint par allowlist explicite.
- [ ] Les journaux de conversation masquent ou expurgent les secrets potentiellement présents.
