// Vulnérable : Server-Side Includes (SSI) Injection (CWE-97)
// Le commentaire utilisateur est écrit tel quel dans un fichier .shtml
// servi par un serveur web interprétant les directives SSI, permettant à
// un attaquant d'injecter une directive (<!--#exec cmd="..." -->).
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/comments")]
public class CommentController : ControllerBase
{
    private const string CommentsFile = "wwwroot/comments.shtml";

    [HttpPost]
    public IActionResult AddComment([FromForm] string comment)
    {
        // Écriture directe, sans neutralisation des séquences SSI.
        System.IO.File.AppendAllText(CommentsFile, $"<p>{comment}</p>\n");
        return Ok();
    }
}
