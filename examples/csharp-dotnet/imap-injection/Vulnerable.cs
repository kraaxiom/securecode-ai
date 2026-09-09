// Vulnérable : IMAP Injection (CWE-93)
// Le critère de recherche fourni par l'utilisateur est concaténé directement
// dans la commande IMAP SEARCH, permettant l'injection de guillemets ou de
// mots-clés IMAP supplémentaires.
using Microsoft.AspNetCore.Mvc;
using MailKit.Net.Imap;
using MailKit.Search;

[ApiController]
[Route("api/mailbox")]
public class MailSearchController : ControllerBase
{
    [HttpGet("search")]
    public IActionResult Search([FromQuery] string q, [FromServices] ImapClient client)
    {
        // Construction manuelle d'une requête de recherche textuelle brute.
        var rawQuery = $"SUBJECT \"{q}\"";
        var results = client.Inbox.Search(SearchQuery.SubjectContains(rawQuery));
        return Ok(results.Select(uid => uid.Id));
    }
}
