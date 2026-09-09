// Vulnérable : OS Command Injection (CWE-78)
// Le nom de fichier fourni par l'utilisateur est concaténé dans les
// arguments passés à un shell, permettant l'injection de méta-caractères
// pour exécuter des commandes arbitraires.
using Microsoft.AspNetCore.Mvc;
using System.Diagnostics;

[ApiController]
[Route("api/files")]
public class ArchiveController : ControllerBase
{
    [HttpGet("compress")]
    public IActionResult Compress([FromQuery] string fileName)
    {
        var psi = new ProcessStartInfo
        {
            FileName = "cmd.exe",
            Arguments = $"/c 7z a archive.zip {fileName}",
            RedirectStandardOutput = true,
            UseShellExecute = false
        };

        using var process = Process.Start(psi);
        var output = process!.StandardOutput.ReadToEnd();
        process.WaitForExit();

        return Ok(output);
    }
}
