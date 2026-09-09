// Corrigé : Blind Cross-Site Scripting (CWE-79)
// Toute donnée externe non fiable (message de ticket, User-Agent journalisé) est
// encodée contextuellement avant d'être insérée dans le HTML du tableau de bord
// admin, y compris pour les interfaces internes considérées à tort comme sûres.
using System.Text.Encodings.Web;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("admin/tickets")]
public class AdminTicketsController : ControllerBase
{
    private readonly ITicketRepository _tickets;
    private readonly HtmlEncoder _htmlEncoder;

    public AdminTicketsController(ITicketRepository tickets, HtmlEncoder htmlEncoder)
    {
        _tickets = tickets;
        _htmlEncoder = htmlEncoder;
    }

    [HttpGet("{id}")]
    public ContentResult ViewTicket(int id)
    {
        var ticket = _tickets.GetById(id);

        var safeMessage = _htmlEncoder.Encode(ticket.Message ?? string.Empty);
        var safeUserAgent = _htmlEncoder.Encode(ticket.UserAgent ?? string.Empty);

        var html = $@"
            <div class='ticket'>
                <h2>Ticket #{ticket.Id}</h2>
                <p class='ticket-message'>{safeMessage}</p>
                <p class='ticket-useragent'>{safeUserAgent}</p>
            </div>";

        return new ContentResult { Content = html, ContentType = "text/html" };
    }
}

// Recommandation complémentaire : dans une vue Razor, préférer l'encodage
// automatique natif (@ticket.Message) plutôt que Html.Raw() sur une donnée
// externe, et appliquer une Content Security Policy stricte sur /admin.
