// Corrigé : Jailbreak de modèle (CWE-1427)
// Une couche de modération indépendante du modèle principal filtre à la
// fois l'entrée et la sortie, le nombre de tours suspects par session est
// surveillé et limité, et le system prompt n'est jamais considéré comme
// l'unique contrôle de sécurité.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/chat")]
public class ChatController : ControllerBase
{
    private const string SystemPrompt = "Tu es un assistant utile et respectueux des règles de l'entreprise.";
    private const int MaxSuspiciousTurns = 3;
    private const string RefusalMessage = "Je ne peux pas traiter cette demande.";

    private readonly ILlmClient _llm;
    private readonly ISafetyClassifier _safetyClassifier;
    private readonly ISessionStore _sessionStore;
    private readonly IAuditLog _auditLog;

    public ChatController(
        ILlmClient llm,
        ISafetyClassifier safetyClassifier,
        ISessionStore sessionStore,
        IAuditLog auditLog)
    {
        _llm = llm;
        _safetyClassifier = safetyClassifier;
        _sessionStore = sessionStore;
        _auditLog = auditLog;
    }

    [HttpPost]
    public async Task<IActionResult> Chat([FromBody] ChatRequest request)
    {
        var session = await _sessionStore.GetOrCreateAsync(request.SessionId);

        // Filtre d'entrée indépendant du modèle principal.
        if (await _safetyClassifier.FlagsAsync(request.Message))
        {
            session.RegisterSuspiciousTurn();
            await _auditLog.RecordAsync("blocked_input_flagged", session.Id);
            return Ok(new { response = RefusalMessage });
        }

        // Surveillance des tentatives répétées de contournement sur la session.
        if (session.SuspiciousTurnCount > MaxSuspiciousTurns)
        {
            await _auditLog.RecordAsync("session_flagged_repeated_attempts", session.Id);
            return Ok(new { response = RefusalMessage });
        }

        var messages = new List<ChatMessage>
        {
            new("system", SystemPrompt)
        };
        messages.AddRange(request.History);
        messages.Add(new("user", request.Message));

        var response = await _llm.ChatAsync(messages);

        // Filtre de sortie indépendant, appliqué avant de renvoyer la réponse.
        if (await _safetyClassifier.FlagsAsync(response.Content))
        {
            await _auditLog.RecordAsync("blocked_output_flagged", session.Id);
            return Ok(new { response = RefusalMessage });
        }

        await _sessionStore.SaveAsync(session);
        return Ok(new { response.Content });
    }
}
