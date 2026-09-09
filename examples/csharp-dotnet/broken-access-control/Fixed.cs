// Corrigé : Broken Access Control (CWE-284)
// La logique d'autorisation est centralisée dans un handler dédié (deny by default)
// appliqué systématiquement avant de servir la ressource sensible.
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/invoices")]
[Authorize]
public class InvoicesController : ControllerBase
{
    private readonly AppDbContext _db;
    private readonly IAuthorizationService _authorizationService;

    public InvoicesController(AppDbContext db, IAuthorizationService authorizationService)
    {
        _db = db;
        _authorizationService = authorizationService;
    }

    [HttpGet("{id:int}/download")]
    public async Task<IActionResult> Download(int id)
    {
        var invoice = _db.Invoices.FirstOrDefault(i => i.Id == id);
        if (invoice is null)
            return NotFound();

        // Vérification centralisée via une policy dédiée à la ressource : refus
        // par défaut sauf autorisation explicite (propriétaire ou administrateur).
        var authResult = await _authorizationService.AuthorizeAsync(User, invoice, "ViewInvoice");
        if (!authResult.Succeeded)
            return Forbid();

        return PhysicalFile(invoice.FilePath, "application/pdf");
    }
}

// InvoiceAuthorizationHandler.cs — règle centralisée, appliquée à chaque accès.
// protected override Task HandleRequirementAsync(AuthorizationHandlerContext context,
//     ViewInvoiceRequirement requirement, Invoice invoice)
// {
//     if (invoice.OwnerId == context.User.GetUserId() || context.User.IsInRole("Admin"))
//         context.Succeed(requirement);
//     return Task.CompletedTask;
// }
