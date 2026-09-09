// Corrigé : Forced Browsing (CWE-425)
// Les fichiers sensibles ne sont plus servis statiquement : un contrôleur dédié
// applique un contrôle d'authentification/autorisation explicite avant de
// renvoyer le fichier, indépendamment de sa découvrabilité côté client.
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/backups")]
[Authorize(Policy = "RequireAdminRole")]
public class BackupsController : ControllerBase
{
    private readonly IBackupStorage _storage;

    public BackupsController(IBackupStorage storage)
    {
        _storage = storage;
    }

    [HttpGet("{fileName}")]
    public IActionResult Download(string fileName)
    {
        // Résolution sécurisée du chemin (empêche aussi la traversée de répertoire)
        // et contrôle d'autorisation systématique, quel que soit le nom du fichier.
        var safePath = _storage.ResolveSafePath(fileName);
        if (safePath is null)
            return NotFound();

        return PhysicalFile(safePath, "application/octet-stream", fileName);
    }
}

// Program.cs — le répertoire de sauvegardes n'est plus jamais servi statiquement :
// app.UseStaticFiles(); // ne pointe plus vers un dossier contenant des sauvegardes/exports
