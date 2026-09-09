// Vulnérable : Session Prediction (CWE-330)
// L'identifiant de session est dérivé d'un hash MD5 de l'ID utilisateur et
// de l'horodatage courant, deux valeurs prévisibles ou devinables, permettant
// à un attaquant de reconstruire ou deviner des identifiants de session valides.
using System.Security.Cryptography;
using System.Text;

public class SessionService
{
    private readonly Dictionary<string, int> _sessions = new();

    public string CreateSession(int userId)
    {
        // Dérivé de valeurs connues/prévisibles : ID utilisateur + timestamp.
        var raw = $"{userId}-{DateTime.UtcNow.Ticks}";
        var hash = MD5.HashData(Encoding.UTF8.GetBytes(raw));
        var sessionId = Convert.ToHexString(hash);

        _sessions[sessionId] = userId;
        return sessionId;
    }

    public int? GetUserId(string sessionId) =>
        _sessions.TryGetValue(sessionId, out var userId) ? userId : null;
}
