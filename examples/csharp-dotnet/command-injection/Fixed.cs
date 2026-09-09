// Corrigé : Command Injection (CWE-78)
// L'entrée est validée strictement comme une adresse IP avant utilisation,
// et l'exécutable est appelé directement avec des arguments distincts,
// sans passer par un interpréteur shell.
using Microsoft.AspNetCore.Mvc;
using System.Diagnostics;
using System.Net;

[ApiController]
[Route("api/network")]
public class PingController : ControllerBase
{
    [HttpGet("ping")]
    public IActionResult Ping([FromQuery] string host)
    {
        if (!IPAddress.TryParse(host, out _))
            return BadRequest("Hôte invalide.");

        var psi = new ProcessStartInfo
        {
            FileName = "ping",
            UseShellExecute = false,
            RedirectStandardOutput = true
        };
        psi.ArgumentList.Add("-n");
        psi.ArgumentList.Add("3");
        psi.ArgumentList.Add(host);

        using var process = Process.Start(psi);
        var output = process!.StandardOutput.ReadToEnd();
        process.WaitForExit();

        return Ok(output);
    }
}
