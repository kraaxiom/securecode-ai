// Corrigé : Universal Cross-Site Scripting / UXSS (CWE-79 / CWE-1021)
// La cause première d'un UXSS échappe souvent au code applicatif (faille de
// navigateur, extension, composant tiers largement intégré). L'application
// reste toutefois responsable de limiter sa surface d'exposition : Content
// Security Policy stricte, Subresource Integrity sur les scripts tiers,
// iframe sandboxée au strict nécessaire, et Permissions-Policy restrictive.
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

public class Startup
{
    public void Configure(IApplicationBuilder app)
    {
        // En-têtes de sécurité appliqués à toutes les réponses.
        app.Use(async (context, next) =>
        {
            context.Response.Headers["Content-Security-Policy"] =
                "default-src 'self'; " +
                "script-src 'self' https://cdn.example.com; " +
                "frame-src 'self' https://widget-tiers.example.com; " +
                "object-src 'none'; base-uri 'self'";
            context.Response.Headers["Permissions-Policy"] =
                "camera=(), microphone=(), geolocation=()";
            context.Response.Headers["X-Content-Type-Options"] = "nosniff";

            await next();
        });
    }
}

public class SupportChatController : ControllerBase
{
    [HttpGet("/support")]
    public ContentResult Page()
    {
        var html = @"
            <!DOCTYPE html>
            <html>
            <body>
                <h1>Support client</h1>

                <!-- Sandbox limité au strict nécessaire : scripts et same-origin
                     uniquement, pas de popups, pas de soumission de formulaire. -->
                <iframe
                    src='https://widget-tiers.example.com/chat'
                    sandbox='allow-scripts allow-same-origin'
                    referrerpolicy='no-referrer'>
                </iframe>

                <!-- Subresource Integrity : le script n'est exécuté que si son
                     empreinte correspond exactement à la version vérifiée. -->
                <script
                    src='https://cdn.example.com/support-widget.js'
                    integrity='sha384-<hash-du-fichier-verifie>'
                    crossorigin='anonymous'>
                </script>
            </body>
            </html>";

        return new ContentResult { Content = html, ContentType = "text/html" };
    }
}
