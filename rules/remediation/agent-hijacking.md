# Remédiation — Agent Hijacking

## Principe
Séparer strictement la couche de décision du LLM de la couche d'exécution applicative. Toute action à fort impact (paiement, suppression, envoi externe) doit passer par une validation métier indépendante du modèle, avec confirmation humaine pour les opérations irréversibles.

## Python
```python
# Avant — vulnérable : l'agent invoque directement un outil à fort impact
def run_agent(user_input):
    plan = llm.plan(user_input, tools=[send_payment, delete_account, send_email])
    for step in plan.steps:
        step.tool.execute(**step.args)  # exécution directe, aucune validation

# Après — sécurisé : validation métier + confirmation humaine avant action irréversible
HIGH_IMPACT_TOOLS = {"send_payment", "delete_account"}

def run_agent(user_input):
    plan = llm.plan(user_input, tools=[send_payment, delete_account, send_email])
    for step in plan.steps:
        if step.tool.name in HIGH_IMPACT_TOOLS:
            if not require_human_confirmation(step):
                audit_log.record("blocked_high_impact_action", step)
                continue
        validate_against_business_rules(step)  # contrôle indépendant du modèle
        step.tool.execute(**step.args)
        audit_log.record("action_executed", step)
```

## JavaScript / Node.js
```js
// Avant — vulnérable
async function runAgent(userInput) {
  const plan = await llm.plan(userInput, { tools: [sendPayment, deleteAccount, sendEmail] });
  for (const step of plan.steps) {
    await step.tool.execute(step.args); // aucune validation
  }
}

// Après — sécurisé
const HIGH_IMPACT_TOOLS = new Set(["sendPayment", "deleteAccount"]);

async function runAgent(userInput) {
  const plan = await llm.plan(userInput, { tools: [sendPayment, deleteAccount, sendEmail] });
  for (const step of plan.steps) {
    if (HIGH_IMPACT_TOOLS.has(step.tool.name)) {
      const confirmed = await requireHumanConfirmation(step);
      if (!confirmed) { auditLog.record("blocked_high_impact_action", step); continue; }
    }
    validateAgainstBusinessRules(step); // contrôle indépendant du modèle
    await step.tool.execute(step.args);
    auditLog.record("action_executed", step);
  }
}
```

## Checklist de vérification post-patch
- [ ] Toute action irréversible ou à fort impact financier/opérationnel exige une confirmation humaine explicite avant exécution.
- [ ] La validation des arguments d'outil est effectuée côté application, indépendamment de ce que le modèle a produit.
- [ ] Les permissions de l'agent sont réduites lors du traitement de contenu externe non fiable (moins de privilèges contextuels).
- [ ] Chaque décision et appel d'outil de l'agent est journalisé de façon exploitable pour l'audit post-incident.
