// Vulnérable : Broken Object Level Authorization (CWE-639)
// L'objet est récupéré uniquement par l'ID fourni par le client, sans vérifier
// que la commande appartient bien à l'utilisateur authentifié.
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/orders")]
[Authorize]
public class OrdersController : ControllerBase
{
    private readonly AppDbContext _db;

    public OrdersController(AppDbContext db)
    {
        _db = db;
    }

    [HttpGet("{id:int}")]
    public IActionResult GetOrder(int id)
    {
        // Aucune clause de filtrage sur le propriétaire : n'importe quel
        // utilisateur authentifié peut lire la commande d'un autre client
        // en changeant simplement l'ID dans l'URL.
        var order = _db.Orders.FirstOrDefault(o => o.Id == id);
        if (order is null)
            return NotFound();

        return Ok(order);
    }
}
