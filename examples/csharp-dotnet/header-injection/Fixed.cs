// Corrigé : Header Injection (CWE-113)
// Le nom de fichier est nettoyé des caractères de contrôle et réduit à son
// nom de base avant d'être utilisé, et l'en-tête est défini via l'API sûre
// FileContentResult plutôt qu'une écriture manuelle.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/files")]
public class DownloadController : ControllerBase
{
    [HttpGet("download")]
    public IActionResult Download([FromQuery] string filename, [FromServices] IFileStore store)
    {
        var safeFilename = Path.GetFileName(filename.Replace("\r", "").Replace("\n", ""));
        var bytes = store.Read(safeFilename);
        return File(bytes, "application/octet-stream", safeFilename);
    }
}
