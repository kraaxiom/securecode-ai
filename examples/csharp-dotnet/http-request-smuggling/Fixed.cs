// Corrigé : HTTP Request Smuggling (CWE-444)
// Le middleware rejette explicitement toute requête présentant simultanément
// Content-Length et Transfer-Encoding, en défense en profondeur du côté
// applicatif (en complément d'une configuration stricte du proxy amont).
using Microsoft.AspNetCore.Http;

public class RequestForwardingMiddleware
{
    private readonly RequestDelegate _next;

    public RequestForwardingMiddleware(RequestDelegate next) => _next = next;

    public async Task InvokeAsync(HttpContext context)
    {
        var hasContentLength = context.Request.Headers.ContainsKey("Content-Length");
        var hasTransferEncoding = context.Request.Headers.ContainsKey("Transfer-Encoding");

        if (hasContentLength && hasTransferEncoding)
        {
            context.Response.StatusCode = StatusCodes.Status400BadRequest;
            await context.Response.WriteAsync("Requête ambiguë refusée.");
            return;
        }

        await _next(context);
    }
}
