// Vulnérable : Command Injection (CWE-78)
// L'hôte fourni par l'utilisateur est concaténé dans les arguments passés
// à un shell via cmd.exe, permettant l'injection de métacaractères shell.
using Microsoft.AspNetCore.Mvc;
using System.Diagnostics;

[ApiController]
[Route("api/network")]
public class PingController : ControllerBase
{
    [HttpGet("ping")]
    public IActionResult Ping([FromQuery] string host)
    {
        var psi = new ProcessStartInfo
        {
            FileName = "cmd.exe",
            Arguments = $"/c ping -n 3 {host}",
            RedirectStandardOutput = true,
            UseShellExecute = false
        };

        using var process = Process.Start(psi);
        var output = process!.StandardOutput.ReadToEnd();
        process.WaitForExit();

        return Ok(output);
    }
}
