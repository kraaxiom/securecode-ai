// Vulnérable : XSS via SVG (CWE-79)
// Le fichier SVG uploadé par l'utilisateur (avatar) est stocké et servi tel
// quel depuis le même domaine que l'application, avec le Content-Type
// image/svg+xml et sans en-tête forçant le téléchargement. Le SVG peut
// contenir des éléments actifs (<script>, gestionnaires on*) qui s'exécutent
// dans le contexte d'origine de l'application lorsqu'il est ouvert directement.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("avatars")]
public class AvatarController : ControllerBase
{
    private readonly string _uploadDir = "wwwroot/uploads/avatars";

    [HttpPost]
    public async Task<IActionResult> Upload(IFormFile file)
    {
        var path = Path.Combine(_uploadDir, file.FileName);
        using var stream = new FileStream(path, FileMode.Create);
        await file.CopyToAsync(stream);
        return Ok(new { file.FileName });
    }

    [HttpGet("{fileName}")]
    public IActionResult Get(string fileName)
    {
        var path = Path.Combine(_uploadDir, fileName);
        var bytes = System.IO.File.ReadAllBytes(path);

        // Servi tel quel, inline, depuis l'origine de l'application principale.
        return File(bytes, "image/svg+xml");
    }
}
