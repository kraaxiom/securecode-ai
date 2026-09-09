// Vulnérable : Insecure Direct Object Reference (CWE-639)
// Le document est récupéré par son identifiant transmis par le client, sans
// vérifier qu'il appartient bien à l'utilisateur authentifié.
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

    [HttpGet("{id:int}")]
    public IActionResult GetDocument(int id)
    {
        // L'ID séquentiel est utilisé tel quel : en incrémentant la valeur,
        // un utilisateur peut accéder aux documents d'un autre compte.
        var document = _db.Documents.FirstOrDefault(d => d.Id == id);
        if (document is null)
            return NotFound();

        return Ok(document);
    }
}
