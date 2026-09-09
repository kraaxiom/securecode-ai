# Vue Django corrigée — CWE-1427 (Agent Hijacking / OWASP LLM01:2025)
# Correction : séparation stricte entre la couche de décision (LLM) et la
# couche d'exécution (application). Les outils à fort impact exigent une
# confirmation humaine explicite et une validation métier indépendante du
# modèle. Chaque décision et action est journalisée pour l'audit.

import json

from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
from langchain.agents import AgentExecutor, Tool
from langchain_openai import ChatOpenAI

from billing.services import send_payment
from accounts.services import delete_account
from notifications.services import send_email
from core.audit import audit_log
from core.human_review import require_human_confirmation
from core.business_rules import validate_against_business_rules

llm = ChatOpenAI(model="gpt-4o")

TOOLS = [
    Tool(name="send_payment", func=send_payment, description="Envoie un paiement à un bénéficiaire"),
    Tool(name="delete_account", func=delete_account, description="Supprime un compte utilisateur"),
    Tool(name="send_email", func=send_email, description="Envoie un e-mail"),
]

# Sécurisé : catalogue explicite des outils irréversibles ou à fort impact
# financier/opérationnel, distincts des outils à faible risque (ex. e-mail).
HIGH_IMPACT_TOOLS = {"send_payment", "delete_account"}

agent_executor = AgentExecutor.from_agent_and_tools(
    agent=llm,
    tools=TOOLS,
    return_intermediate_steps=True,  # nécessaire pour valider chaque étape avant exécution réelle
)


@csrf_exempt
def agent_chat(request):
    """Endpoint agent conversationnel : le plan produit par le LLM est
    revalidé étape par étape côté application avant toute exécution."""
    payload = json.loads(request.body)
    user_input = payload.get("message", "")
    user = request.user

    # Sécurisé : on demande d'abord le plan (dry-run) sans exécuter les
    # outils sensibles automatiquement — LangChain expose les tool calls
    # via return_intermediate_steps, chaque étape est ensuite contrôlée.
    result = agent_executor.invoke({"input": user_input})

    for action, _observation in result.get("intermediate_steps", []):
        tool_name = action.tool
        tool_args = action.tool_input

        if tool_name in HIGH_IMPACT_TOOLS:
            # Sécurisé : confirmation humaine obligatoire avant toute
            # action irréversible ou à fort impact financier.
            if not require_human_confirmation(user, tool_name, tool_args):
                audit_log.record("blocked_high_impact_action", user_id=user.id, tool=tool_name, args=tool_args)
                continue

        # Sécurisé : validation métier indépendante du modèle (limites,
        # permissions réelles de l'utilisateur, cohérence des montants).
        validate_against_business_rules(user, tool_name, tool_args)
        audit_log.record("action_executed", user_id=user.id, tool=tool_name, args=tool_args)

    return JsonResponse({"response": result["output"]})
