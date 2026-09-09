// Vulnérable : Header Injection (CWE-113)
// Le nom de fichier fourni par l'utilisateur est concaténé directement dans
// l'en-tête Content-Disposition, permettant l'injection de \r\n pour
// ajouter des en-têtes arbitraires.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/files")]
public class DownloadController : ControllerBase
{
    [HttpGet("download")]
    public IActionResult Download([FromQuery] string filename, [FromServices] IFileStore store)
    {
        Response.Headers["Content-Disposition"] = $"attachment; filename={filename}";
        var bytes = store.Read(filename);
        return File(bytes, "application/octet-stream");
    }
}
