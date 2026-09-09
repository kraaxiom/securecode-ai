// Vulnérable : Mauvaise configuration Firebase (CWE-284)
// Le backend fait confiance à l'userId fourni par le client sans vérifier
// le jeton d'authentification Firebase ni la propriété du document,
// reflétant côté serveur des règles Firestore en mode test ("if true").
using Google.Cloud.Firestore;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/documents")]
public class DocumentsController : ControllerBase
{
    private readonly FirestoreDb _firestore;

    public DocumentsController(FirestoreDb firestore) => _firestore = firestore;

    // VULNÉRABLE : userId et docId viennent directement de la requête client,
    // aucune vérification du token Firebase ni de la propriété du document.
    [HttpGet("{userId}/{docId}")]
    public async Task<IActionResult> GetDocument(string userId, string docId)
    {
        var doc = await _firestore
            .Collection("users").Document(userId)
            .Collection("documents").Document(docId)
            .GetSnapshotAsync();

        if (!doc.Exists) return NotFound();
        return Ok(doc.ToDictionary());
    }

    // VULNÉRABLE : écriture acceptée sans authentification, n'importe qui
    // peut modifier les documents de n'importe quel utilisateur.
    [HttpPut("{userId}/{docId}")]
    public async Task<IActionResult> UpdateDocument(string userId, string docId, [FromBody] Dictionary<string, object> data)
    {
        await _firestore
            .Collection("users").Document(userId)
            .Collection("documents").Document(docId)
            .SetAsync(data);

        return Ok();
    }
}
