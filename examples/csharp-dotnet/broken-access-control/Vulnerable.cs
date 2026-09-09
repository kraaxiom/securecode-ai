// Vulnérable : Broken Access Control (CWE-284)
// Le téléchargement d'une facture n'applique aucune vérification d'autorisation
// propre à la ressource : seule l'authentification générale est contrôlée.
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/invoices")]
[Authorize]
public class InvoicesController : ControllerBase
{
    private readonly AppDbContext _db;

    public InvoicesController(AppDbContext db)
    {
        _db = db;
    }

    [HttpGet("{id:int}/download")]
    public IActionResult Download(int id)
    {
        var invoice = _db.Invoices.FirstOrDefault(i => i.Id == id);
        if (invoice is null)
            return NotFound();

        // Aucune vérification que l'utilisateur est autorisé à voir cette facture
        // précise : le contrôle d'accès repose uniquement sur l'authentification.
        return PhysicalFile(invoice.FilePath, "application/pdf");
    }
}
