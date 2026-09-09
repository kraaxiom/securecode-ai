# Remédiation — Jailbreak de modèle

## Principe défensif
Ne jamais considérer le system prompt comme seule barrière de sécurité. Ajouter une couche de modération indépendante du modèle (classifieur en entrée et en sortie), limiter le contexte conversationnel exploitable, et surveiller les schémas de conversation anormaux. Ce document ne documente aucune technique de contournement — uniquement des contrôles défensifs applicatifs.

## Python
```python
# Avant — vulnérable : le system prompt est la seule barrière, aucune modération indépendante
def chat(user_message, history):
    messages = [{"role": "system", "content": SYSTEM_PROMPT}] + history + [{"role": "user", "content": user_message}]
    return llm.chat(messages=messages)

# Après — sécurisé : modération d'entrée/sortie indépendante + surveillance conversationnelle
def chat(user_message, history, session):
    if safety_classifier.flags(user_message):  # classifieur indépendant du LLM principal
        audit_log.record("blocked_input_flagged", session.id)
        return REFUSAL_MESSAGE

    if session.suspicious_turn_count() > MAX_SUSPICIOUS_TURNS:
        audit_log.record("session_flagged_repeated_attempts", session.id)
        return REFUSAL_MESSAGE

    messages = [{"role": "system", "content": SYSTEM_PROMPT}] + history + [{"role": "user", "content": user_message}]
    response = llm.chat(messages=messages)

    if safety_classifier.flags(response.content):  # filtre de sortie indépendant
        audit_log.record("blocked_output_flagged", session.id)
        return REFUSAL_MESSAGE

    return response
```

## JavaScript / Node.js
```js
// Avant — vulnérable
async function chat(userMessage, history) {
  const messages = [{ role: "system", content: SYSTEM_PROMPT }, ...history, { role: "user", content: userMessage }];
  return llm.chat({ messages });
}

// Après — sécurisé
async function chat(userMessage, history, session) {
  if (await safetyClassifier.flags(userMessage)) {
    auditLog.record("blocked_input_flagged", session.id);
    return REFUSAL_MESSAGE;
  }
  if (session.suspiciousTurnCount() > MAX_SUSPICIOUS_TURNS) {
    auditLog.record("session_flagged_repeated_attempts", session.id);
    return REFUSAL_MESSAGE;
  }

  const messages = [{ role: "system", content: SYSTEM_PROMPT }, ...history, { role: "user", content: userMessage }];
  const response = await llm.chat({ messages });

  if (await safetyClassifier.flags(response.content)) {
    auditLog.record("blocked_output_flagged", session.id);
    return REFUSAL_MESSAGE;
  }
  return response;
}
```

## Checklist de vérification post-patch
- [ ] Une couche de modération indépendante du modèle principal filtre à la fois l'entrée et la sortie.
- [ ] Le system prompt n'est jamais le seul contrôle de sécurité — des contrôles applicatifs existent en aval de la réponse.
- [ ] Le nombre de tours de conversation suspects par session est surveillé et limite l'accès en cas de dépassement.
- [ ] Des campagnes de red teaming régulières sont documentées, avec suivi des régressions après mise à jour du modèle.
