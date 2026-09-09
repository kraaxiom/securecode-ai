// Vulnérable : HTTP Request Smuggling (CWE-444)
// Le middleware fait confiance aux en-têtes Content-Length et Transfer-Encoding
// transmis par la chaîne de proxys sans vérifier leur cohérence, ouvrant la
// voie à une divergence d'interprétation entre le frontal et ce backend.
using Microsoft.AspNetCore.Http;

public class RequestForwardingMiddleware
{
    private readonly RequestDelegate _next;

    public RequestForwardingMiddleware(RequestDelegate next) => _next = next;

    public async Task InvokeAsync(HttpContext context)
    {
        // Aucune vérification de cohérence entre Content-Length et Transfer-Encoding.
        await _next(context);
    }
}
