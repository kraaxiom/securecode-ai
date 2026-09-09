// Vulnérable : Universal Cross-Site Scripting / UXSS (CWE-79 / CWE-1021)
// L'application intègre un widget tiers (chat support) via une iframe sans
// attribut sandbox, et ne définit aucune Content Security Policy ni
// Permissions-Policy. Si le widget tiers est compromis, ou si une faille de
// navigateur/extension permet de contourner la Same-Origin Policy, l'absence
// de restrictions applicatives maximise la surface d'exposition de l'application.
using Microsoft.AspNetCore.Mvc;

public class SupportChatController : ControllerBase
{
    [HttpGet("/support")]
    public ContentResult Page()
    {
        var html = @"
            <!DOCTYPE html>
            <html>
            <head>
                <!-- Aucune Content-Security-Policy définie -->
            </head>
            <body>
                <h1>Support client</h1>

                <!-- Iframe tierce sans attribut sandbox : capacités par défaut
                     du navigateur (scripts, formulaires, popups...) autorisées. -->
                <iframe src='https://widget-tiers.example.com/chat'></iframe>

                <!-- Script tiers chargé sans Subresource Integrity : toute
                     compromission du CDN se traduit par une exécution de code
                     arbitraire dans le contexte de la page. -->
                <script src='https://cdn.example.com/support-widget.js'></script>
            </body>
            </html>";

        return new ContentResult { Content = html, ContentType = "text/html" };
    }
}
