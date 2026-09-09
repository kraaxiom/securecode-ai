# Remédiation — Memory Poisoning

## Principe
Exiger une confirmation explicite de l'utilisateur avant toute écriture durable en mémoire persistante, cloisonner strictement la mémoire par utilisateur/tenant, et traiter le contenu réinjecté comme une donnée à revalider plutôt qu'une instruction de confiance.

## Python
```python
# Avant — vulnérable : écriture automatique en mémoire persistante sans validation
def process_turn(user_id, message, agent_response):
    extracted_facts = llm.extract_facts(agent_response)
    memory_store.append(user_id, extracted_facts)  # écrit sans confirmation ni cloisonnement vérifié

def next_session(user_id):
    memory = memory_store.get(user_id)
    return llm.chat(messages=[{"role": "system", "content": f"Contexte connu : {memory}"}])

# Après — sécurisé : confirmation utilisateur + cloisonnement + revalidation
def process_turn(user_id, message, agent_response):
    extracted_facts = llm.extract_facts(agent_response)
    for fact in extracted_facts:
        if user_confirms(user_id, fact):  # confirmation explicite avant écriture durable
            memory_store.append(tenant_scoped_key(user_id), fact, status="user_confirmed")
        else:
            audit_log.record("memory_write_rejected", user_id, fact)

def next_session(user_id):
    memory = memory_store.get(tenant_scoped_key(user_id))
    validated_memory = revalidate(memory)  # traité comme donnée, pas comme instruction de confiance
    return llm.chat(messages=[{"role": "system", "content": f"Contexte à vérifier, non prescriptif : {validated_memory}"}])
```

## JavaScript / Node.js
```js
// Avant — vulnérable
async function processTurn(userId, message, agentResponse) {
  const facts = await llm.extractFacts(agentResponse);
  await memoryStore.append(userId, facts);
}

// Après — sécurisé
async function processTurn(userId, message, agentResponse) {
  const facts = await llm.extractFacts(agentResponse);
  for (const fact of facts) {
    if (await userConfirms(userId, fact)) {
      await memoryStore.append(tenantScopedKey(userId), fact, { status: "user_confirmed" });
    } else {
      auditLog.record("memory_write_rejected", userId, fact);
    }
  }
}

async function nextSession(userId) {
  const memory = await memoryStore.get(tenantScopedKey(userId));
  const validatedMemory = revalidate(memory);
  return llm.chat({ messages: [{ role: "system", content: `Contexte à vérifier, non prescriptif : ${validatedMemory}` }] });
}
```

## Checklist de vérification post-patch
- [ ] Aucune écriture en mémoire persistante n'a lieu sans confirmation explicite de l'utilisateur concerné.
- [ ] La mémoire est strictement cloisonnée par utilisateur/tenant (clé scopée, pas de partage implicite).
- [ ] Le contenu mémoire réinjecté en contexte est marqué comme non prescriptif et revalidé, pas traité comme une instruction de confiance.
- [ ] L'utilisateur dispose d'une interface pour consulter, corriger et supprimer les entrées mémoire associées à son compte.
