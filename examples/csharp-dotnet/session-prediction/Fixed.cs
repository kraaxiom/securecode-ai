// Corrigé : Session Prediction (CWE-330)
// L'identifiant de session est généré exclusivement via un générateur
// cryptographiquement sûr (RandomNumberGenerator), avec 256 bits d'entropie,
// et régénéré à chaque changement de niveau de privilège.
using System.Security.Cryptography;
using System.Collections.Concurrent;

public class SessionService
{
    private readonly ConcurrentDictionary<string, int> _sessions = new();

    public string CreateSession(int userId)
    {
        // Générateur cryptographique dédié : 256 bits d'entropie, aucune valeur prévisible.
        var sessionId = Convert.ToHexString(RandomNumberGenerator.GetBytes(32));
        _sessions[sessionId] = userId;
        return sessionId;
    }

    public string RegenerateSession(string oldSessionId)
    {
        // Renouvellement de l'identifiant à chaque élévation de privilège (ex. connexion).
        if (!_sessions.TryRemove(oldSessionId, out var userId))
        {
            throw new InvalidOperationException("Session inexistante.");
        }

        return CreateSession(userId);
    }

    public int? GetUserId(string sessionId) =>
        _sessions.TryGetValue(sessionId, out var userId) ? userId : null;
}
