// Corrigé : Server-Side Includes (SSI) Injection (CWE-97)
// Le contenu HTML est échappé, puis toute séquence de directive SSI
// résiduelle ("<!--#") est neutralisée avant écriture. Préférable encore :
// ne jamais écrire dans un fichier interprété par SSI (utiliser un moteur
// de templates applicatif à la place).
using Microsoft.AspNetCore.Mvc;
using System.Net;
using System.Text.RegularExpressions;

[ApiController]
[Route("api/comments")]
public class CommentController : ControllerBase
{
    private const string CommentsFile = "wwwroot/comments.shtml";

    [HttpPost]
    public IActionResult AddComment([FromForm] string comment)
    {
        if (string.IsNullOrWhiteSpace(comment))
            return BadRequest("Commentaire vide.");

        var escaped = WebUtility.HtmlEncode(comment);
        var safe = StripSsiDirectives(escaped);

        System.IO.File.AppendAllText(CommentsFile, $"<p>{safe}</p>\n");
        return Ok();
    }

    private static string StripSsiDirectives(string value) =>
        Regex.Replace(value, "<!--#", "&lt;!--#");
}
