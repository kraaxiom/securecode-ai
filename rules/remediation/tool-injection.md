# Remédiation — Tool Injection

## Principe
Valider strictement les arguments de chaque appel d'outil contre un schéma typé avant exécution, ne charger que des définitions d'outils depuis des sources de confiance vérifiées, et appliquer une liste blanche contextuelle d'outils selon le niveau de confiance du contenu traité.

## Python
```python
# Avant — vulnérable : arguments transmis directement du texte généré vers l'exécution
def execute_tool_call(call):
    tool = TOOL_REGISTRY[call.name]
    return tool.execute(**call.args)  # aucune validation de schéma

# Après — sécurisé : validation de schéma + allowlist contextuelle + outils de sources vérifiées
TRUSTED_TOOL_SOURCES = {"internal-registry"}

def load_tool_definitions(source):
    if source.id not in TRUSTED_TOOL_SOURCES:
        audit_log.record("rejected_untrusted_tool_source", source.id)
        raise UntrustedToolSource(source.id)
    return source.load()

def execute_tool_call(call, context_trust_level):
    if call.name not in ALLOWED_TOOLS_BY_TRUST[context_trust_level]:
        audit_log.record("blocked_tool_not_in_allowlist", call.name, context_trust_level)
        raise ToolNotAllowed(call.name)

    schema = TOOL_SCHEMAS[call.name]
    validated_args = schema.validate(call.args)  # rejette tout argument hors schéma
    tool = TOOL_REGISTRY[call.name]
    result = tool.execute(**validated_args)
    audit_log.record("tool_executed", call.name, validated_args)
    return result
```

## JavaScript / Node.js
```js
// Avant — vulnérable
async function executeToolCall(call) {
  const tool = toolRegistry[call.name];
  return tool.execute(call.args); // aucune validation
}

// Après — sécurisé
const TRUSTED_TOOL_SOURCES = new Set(["internal-registry"]);

function loadToolDefinitions(source) {
  if (!TRUSTED_TOOL_SOURCES.has(source.id)) {
    auditLog.record("rejected_untrusted_tool_source", source.id);
    throw new Error("Untrusted tool source");
  }
  return source.load();
}

async function executeToolCall(call, contextTrustLevel) {
  if (!allowedToolsByTrust[contextTrustLevel].has(call.name)) {
    auditLog.record("blocked_tool_not_in_allowlist", call.name, contextTrustLevel);
    throw new Error("Tool not allowed");
  }
  const validatedArgs = toolSchemas[call.name].validate(call.args);
  const result = await toolRegistry[call.name].execute(validatedArgs);
  auditLog.record("tool_executed", call.name, validatedArgs);
  return result;
}
```

## Checklist de vérification post-patch
- [ ] Chaque appel d'outil est validé contre un schéma typé strict avant exécution, indépendamment du texte généré par le modèle.
- [ ] Les définitions d'outils ne sont chargées que depuis des sources de confiance vérifiées et revues.
- [ ] Une liste blanche contextuelle limite les outils disponibles selon le niveau de confiance du contenu en cours de traitement.
- [ ] Chaque appel d'outil et ses arguments réels sont journalisés pour permettre l'audit après incident.
