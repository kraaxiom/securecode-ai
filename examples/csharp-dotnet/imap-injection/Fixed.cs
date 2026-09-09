// Corrigé : IMAP Injection (CWE-93)
// Le terme de recherche est nettoyé des guillemets et caractères de contrôle
// et transmis via l'API structurée MailKit (SearchQuery), qui encode
// correctement le littéral au lieu de construire une commande brute.
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
        var safeTerm = (q ?? string.Empty)
            .Replace("\"", "")
            .Replace("\r", "")
            .Replace("\n", "");
        if (safeTerm.Length > 200)
            safeTerm = safeTerm[..200];

        var results = client.Inbox.Search(SearchQuery.SubjectContains(safeTerm));
        return Ok(results.Select(uid => uid.Id));
    }
}
