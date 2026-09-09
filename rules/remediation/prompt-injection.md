# Remédiation — Prompt Injection

## Principe défensif
Utiliser les API supportant une séparation structurée des rôles (system/user) plutôt que la concaténation de chaînes, traiter toute sortie du modèle comme une donnée non fiable avant action, et appliquer le principe du moindre privilège aux outils exposés à l'agent. Ce document décrit uniquement des défenses, aucun payload d'injection.

## Python
```python
# Avant — vulnérable : concaténation brute, aucune validation de sortie avant action
def chat(user_input):
    prompt = SYSTEM_PROMPT + "\n" + user_input  # concaténation de texte brut
    response = llm.complete(prompt, tools=[delete_file, send_email])
    for call in response.tool_calls:
        call.tool.execute(**call.args)  # exécution directe sans validation
    return response

# Après — sécurisé : rôles structurés + validation de sortie + moindre privilège
def chat(user_input):
    response = llm.chat(
        messages=[
            {"role": "system", "content": SYSTEM_PROMPT},
            {"role": "user", "content": user_input},
        ],
        tools=[send_email],  # outils à faible impact uniquement, par défaut
    )
    for call in response.tool_calls:
        validate_against_schema(call.tool.name, call.args)  # validation applicative indépendante
        if call.tool.name in HIGH_IMPACT_TOOLS:
            require_human_confirmation(call)
        call.tool.execute(**call.args)
    audit_log.record("chat_turn", user_input, response)
    return response
```

## JavaScript / Node.js
```js
// Avant — vulnérable
async function chat(userInput) {
  const prompt = SYSTEM_PROMPT + "\n" + userInput;
  const response = await llm.complete(prompt, { tools: [deleteFile, sendEmail] });
  for (const call of response.toolCalls) {
    await call.tool.execute(call.args); // aucune validation
  }
  return response;
}

// Après — sécurisé
async function chat(userInput) {
  const response = await llm.chat({
    messages: [
      { role: "system", content: SYSTEM_PROMPT },
      { role: "user", content: userInput },
    ],
    tools: [sendEmail],
  });
  for (const call of response.toolCalls) {
    validateAgainstSchema(call.tool.name, call.args);
    if (HIGH_IMPACT_TOOLS.has(call.tool.name)) {
      await requireHumanConfirmation(call);
    }
    await call.tool.execute(call.args);
  }
  auditLog.record("chat_turn", userInput, response);
  return response;
}
```

## Checklist de vérification post-patch
- [ ] Le prompt est construit via les rôles structurés de l'API (system/user/assistant), jamais par concaténation de chaîne brute.
- [ ] Chaque appel d'outil déclenché par le modèle est validé contre un schéma strict avant exécution.
- [ ] Le principe du moindre privilège est appliqué aux outils exposés à l'agent, avec confirmation pour les actions sensibles.
- [ ] Les prompts système et les sorties utilisées pour des décisions automatisées sont journalisés.
