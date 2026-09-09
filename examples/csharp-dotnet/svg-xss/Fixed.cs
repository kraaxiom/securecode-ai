// Corrigé : XSS via SVG (CWE-79)
// Le contenu SVG est sanitisé (suppression des balises <script>, gestionnaires
// d'événements, schémas javascript:) avant stockage. Le fichier est ensuite
// servi en téléchargement forcé plutôt qu'en affichage inline, ce qui limite
// l'impact même en cas de contournement de la sanitisation. Dans un déploiement
// réel, servir ces fichiers depuis un sous-domaine dédié sans cookies est
// recommandé en défense en profondeur supplémentaire.
using System.Xml;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("avatars")]
public class AvatarController : ControllerBase
{
    private readonly string _uploadDir = "wwwroot/uploads/avatars";
    private static readonly string[] SafeExtensions = { ".svg" };

    [HttpPost]
    public async Task<IActionResult> Upload(IFormFile file)
    {
        var extension = Path.GetExtension(file.FileName);
        if (!SafeExtensions.Contains(extension, StringComparer.OrdinalIgnoreCase))
            return BadRequest("Extension non autorisée.");

        string rawSvg;
        using (var reader = new StreamReader(file.OpenReadStream()))
            rawSvg = await reader.ReadToEndAsync();

        var cleanSvg = SvgSanitizer.Sanitize(rawSvg);

        var safeFileName = Guid.NewGuid().ToString("N") + ".svg";
        var path = Path.Combine(_uploadDir, safeFileName);
        await System.IO.File.WriteAllTextAsync(path, cleanSvg);

        return Ok(new { FileName = safeFileName });
    }

    [HttpGet("{fileName}")]
    public IActionResult Get(string fileName)
    {
        var path = Path.Combine(_uploadDir, Path.GetFileName(fileName));
        var bytes = System.IO.File.ReadAllBytes(path);

        // Téléchargement forcé plutôt qu'affichage inline : même si un vecteur
        // de sanitisation venait à être contourné, le navigateur ne rend pas
        // le SVG comme document actif dans le contexte de l'application.
        return File(bytes, "image/svg+xml", fileDownloadName: Path.GetFileName(fileName));
    }
}

// Sanitiseur SVG minimal : retire les balises <script> et les attributs
// d'événements ("on*") ainsi que les schémas javascript: dans les attributs.
public static class SvgSanitizer
{
    public static string Sanitize(string rawSvg)
    {
        var doc = new XmlDocument { XmlResolver = null };
        doc.LoadXml(rawSvg);

        var scriptNodes = doc.GetElementsByTagName("script");
        for (int i = scriptNodes.Count - 1; i >= 0; i--)
            scriptNodes[i].ParentNode?.RemoveChild(scriptNodes[i]);

        void CleanAttributes(XmlNode node)
        {
            if (node.Attributes != null)
            {
                for (int i = node.Attributes.Count - 1; i >= 0; i--)
                {
                    var attr = node.Attributes[i];
                    var isEventHandler = attr.Name.StartsWith("on", StringComparison.OrdinalIgnoreCase);
                    var isJsScheme = attr.Value.TrimStart()
                        .StartsWith("javascript:", StringComparison.OrdinalIgnoreCase);

                    if (isEventHandler || isJsScheme)
                        node.Attributes.Remove(attr);
                }
            }

            foreach (XmlNode child in node.ChildNodes)
                CleanAttributes(child);
        }

        if (doc.DocumentElement != null)
            CleanAttributes(doc.DocumentElement);

        return doc.OuterXml;
    }
}
