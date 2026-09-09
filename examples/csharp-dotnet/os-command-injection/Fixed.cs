// Corrigé : OS Command Injection (CWE-78)
// Le nom de fichier est validé par liste blanche stricte (nom de fichier
// simple, sans séparateur de chemin), puis l'exécutable est invoqué
// directement avec des arguments distincts, sans passer par un shell.
using Microsoft.AspNetCore.Mvc;
using System.Diagnostics;
using System.Text.RegularExpressions;

[ApiController]
[Route("api/files")]
public class ArchiveController : ControllerBase
{
    [HttpGet("compress")]
    public IActionResult Compress([FromQuery] string fileName)
    {
        // Liste blanche stricte : nom de fichier simple uniquement.
        if (string.IsNullOrEmpty(fileName) || !Regex.IsMatch(fileName, @"^[a-zA-Z0-9_\-.]{1,128}$"))
            return BadRequest("Nom de fichier invalide.");

        var psi = new ProcessStartInfo
        {
            FileName = "7z",
            UseShellExecute = false,
            RedirectStandardOutput = true
        };
        psi.ArgumentList.Add("a");
        psi.ArgumentList.Add("archive.zip");
        psi.ArgumentList.Add(fileName);

        using var process = Process.Start(psi);
        var output = process!.StandardOutput.ReadToEnd();
        process.WaitForExit();

        return Ok(output);
    }
}
