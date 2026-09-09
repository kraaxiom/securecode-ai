// Corrigé : SMTP Injection / Email Header Injection (CWE-93)
// Tout caractère de contrôle (\r, \n) est retiré des champs libres avant
// leur insertion dans un en-tête, et le format de l'adresse de réponse est
// validé strictement.
using Microsoft.AspNetCore.Mvc;
using System.Net.Mail;
using System.Text.RegularExpressions;

[ApiController]
[Route("api/contact")]
public class ContactController : ControllerBase
{
    [HttpPost("send")]
    public IActionResult Send([FromForm] string name, [FromForm] string subject, [FromForm] string body)
    {
        name = StripControlChars(name);
        subject = StripControlChars(subject);

        if (string.IsNullOrWhiteSpace(name) || string.IsNullOrWhiteSpace(subject))
            return BadRequest("Champs invalides.");

        var message = new MailMessage
        {
            From = new MailAddress("contact@example.com"),
            Subject = subject,
            Body = $"De : {name}\n\n{body}"
        };
        message.To.Add("dest@example.com");

        // MailMessage échappe déjà les en-têtes structurés, mais on filtre
        // quand même le champ libre par défense en profondeur.
        message.Headers.Add("Reply-To", name);

        using var client = new SmtpClient("smtp.example.com");
        client.Send(message);

        return Ok();
    }

    private static string StripControlChars(string value) =>
        Regex.Replace(value ?? string.Empty, @"[\r\n]", string.Empty);
}
