// Vulnérable : SMTP Injection / Email Header Injection (CWE-93)
// Le nom fourni par l'utilisateur est concaténé directement dans l'en-tête
// "Reply-To", sans filtrage des caractères CR/LF, permettant d'injecter des
// en-têtes arbitraires (Cc, Bcc) ou un second corps de message.
using Microsoft.AspNetCore.Mvc;
using System.Net.Mail;

[ApiController]
[Route("api/contact")]
public class ContactController : ControllerBase
{
    [HttpPost("send")]
    public IActionResult Send([FromForm] string name, [FromForm] string subject, [FromForm] string body)
    {
        var message = new MailMessage
        {
            From = new MailAddress("contact@example.com"),
            // Concaténation directe : name peut contenir "\r\nBcc: attaquant@evil.com"
            Subject = subject,
            Body = $"De : {name}\n\n{body}"
        };
        message.To.Add("dest@example.com");
        message.Headers.Add("Reply-To", name);

        using var client = new SmtpClient("smtp.example.com");
        client.Send(message);

        return Ok();
    }
}
