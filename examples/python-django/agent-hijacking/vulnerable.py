# Vue Django vulnérable — CWE-1427 (Agent Hijacking / OWASP LLM01:2025)
# L'agent LangChain planifie ET exécute directement des outils à fort
# impact (paiement, suppression de compte) sans validation métier
# indépendante ni confirmation humaine. Un contenu externe non fiable
# (message utilisateur ou document ingéré par un outil de lecture) peut
# ainsi détourner l'agent pour déclencher des actions irréversibles.

import json

from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
from langchain.agents import AgentExecutor, Tool
from langchain_openai import ChatOpenAI

from billing.services import send_payment
from accounts.services import delete_account
from notifications.services import send_email

llm = ChatOpenAI(model="gpt-4o")

# Vulnérable : tous les outils, y compris les plus sensibles, sont
# exposés à l'agent avec les mêmes privilèges, sans distinction.
TOOLS = [
    Tool(name="send_payment", func=send_payment, description="Envoie un paiement à un bénéficiaire"),
    Tool(name="delete_account", func=delete_account, description="Supprime un compte utilisateur"),
    Tool(name="send_email", func=send_email, description="Envoie un e-mail"),
]

agent_executor = AgentExecutor.from_agent_and_tools(agent=llm, tools=TOOLS)


@csrf_exempt
def agent_chat(request):
    """Endpoint agent conversationnel : exécute directement le plan produit
    par le LLM, sans aucune couche de validation applicative."""
    payload = json.loads(request.body)
    user_input = payload.get("message", "")

    # Vulnérable : le plan de l'agent (y compris les appels à des outils
    # à fort impact) est exécuté tel quel, sans vérification côté serveur,
    # sans confirmation humaine, et sans journalisation détaillée.
    result = agent_executor.invoke({"input": user_input})

    return JsonResponse({"response": result["output"]})
