// Corrigé : Mauvaise configuration Firebase (CWE-284)
// Le backend vérifie le jeton d'identité Firebase à chaque requête,
// s'assure que l'UID authentifié correspond à la ressource demandée,
// et utilise un compte de service dédié à privilèges minimaux.
using FirebaseAdmin.Auth;
using Google.Cloud.Firestore;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/documents")]
public class DocumentsController : ControllerBase
{
    private readonly FirestoreDb _firestore;

    public DocumentsController(FirestoreDb firestore) => _firestore = firestore;

    [HttpGet("{userId}/{docId}")]
    public async Task<IActionResult> GetDocument(string userId, string docId)
    {
        var authenticatedUid = await VerifyAndGetUidAsync();
        if (authenticatedUid is null) return Unauthorized();

        // Un utilisateur ne peut lire que ses propres documents.
        if (authenticatedUid != userId) return Forbid();

        var doc = await _firestore
            .Collection("users").Document(userId)
            .Collection("documents").Document(docId)
            .GetSnapshotAsync();

        if (!doc.Exists) return NotFound();
        return Ok(doc.ToDictionary());
    }

    [HttpPut("{userId}/{docId}")]
    public async Task<IActionResult> UpdateDocument(string userId, string docId, [FromBody] Dictionary<string, object> data)
    {
        var authenticatedUid = await VerifyAndGetUidAsync();
        if (authenticatedUid is null) return Unauthorized();
        if (authenticatedUid != userId) return Forbid();

        // Validation de schéma minimale : seuls les champs autorisés sont acceptés.
        var allowedFields = new HashSet<string> { "title", "content", "updatedAt" };
        if (data.Keys.Any(k => !allowedFields.Contains(k)))
            return BadRequest("Champs non autorisés.");

        await _firestore
            .Collection("users").Document(userId)
            .Collection("documents").Document(docId)
            .SetAsync(data, SetOptions.MergeAll);

        return Ok();
    }

    // Vérifie le jeton d'identité Firebase transmis dans l'en-tête Authorization
    // et retourne l'UID authentifié, ou null si le jeton est absent/invalide.
    private async Task<string?> VerifyAndGetUidAsync()
    {
        var authHeader = Request.Headers.Authorization.ToString();
        if (!authHeader.StartsWith("Bearer ")) return null;

        var idToken = authHeader["Bearer ".Length..];
        try
        {
            // Compte de service dédié, à privilèges minimaux (rôle limité
            // à la vérification de jetons), configuré au démarrage de l'app.
            var decoded = await FirebaseAuth.DefaultInstance.VerifyIdTokenAsync(idToken);
            return decoded.Uid;
        }
        catch (FirebaseAuthException)
        {
            return null;
        }
    }
}
