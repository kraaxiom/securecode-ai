// Vulnérable : Blind Cross-Site Scripting (CWE-79)
// Le message d'un ticket support soumis par un utilisateur externe non authentifié
// est réaffiché tel quel dans le tableau de bord d'administration, sans encodage.
// L'exécution se produit dans un contexte que l'attaquant ne peut pas observer
// directement (session de l'administrateur).
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("admin/tickets")]
public class AdminTicketsController : ControllerBase
{
    private readonly ITicketRepository _tickets;

    public AdminTicketsController(ITicketRepository tickets)
    {
        _tickets = tickets;
    }

    [HttpGet("{id}")]
    public ContentResult ViewTicket(int id)
    {
        var ticket = _tickets.GetById(id);

        var html = $@"
            <div class='ticket'>
                <h2>Ticket #{ticket.Id}</h2>
                <p class='ticket-message'>{ticket.Message}</p>
                <p class='ticket-useragent'>{ticket.UserAgent}</p>
            </div>";

        return new ContentResult { Content = html, ContentType = "text/html" };
    }
}
