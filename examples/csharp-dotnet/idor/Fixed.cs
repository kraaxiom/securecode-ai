// Corrigé : Insecure Direct Object Reference (CWE-639)
// La requête inclut une clause de vérification d'appartenance directement
// au niveau de la base de données, en complément d'un identifiant non séquentiel.
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/documents")]
[Authorize]
public class DocumentsController : ControllerBase
{
    private readonly AppDbContext _db;

    public DocumentsController(AppDbContext db)
    {
        _db = db;
    }

    [HttpGet("{id:guid}")] // identifiant non séquentiel (UUID) en complément du contrôle d'accès
    public IActionResult GetDocument(Guid id)
    {
        var currentUserId = User.GetUserId();

        // Clause d'appartenance appliquée directement dans la requête de données,
        // pas en post-traitement après récupération de l'objet.
        var document = _db.Documents.FirstOrDefault(d => d.Id == id && d.OwnerId == currentUserId);
        if (document is null)
            return NotFound();

        return Ok(document);
    }
}
