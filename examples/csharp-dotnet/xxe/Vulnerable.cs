// Vulnérable : XML External Entity (XXE) Injection (CWE-611)
// Le XmlReader est créé avec DtdProcessing.Parse et sans restriction sur
// la résolution d'entités externes, permettant à un attaquant de définir
// une entité pointant vers un fichier local ou une URL (divulgation de
// fichiers, SSRF, déni de service par "entity expansion").
using Microsoft.AspNetCore.Mvc;
using System.Xml;

[ApiController]
[Route("api/import")]
public class XmlImportController : ControllerBase
{
    [HttpPost]
    public async Task<IActionResult> Import()
    {
        using var stream = Request.Body;

        var settings = new XmlReaderSettings
        {
            DtdProcessing = DtdProcessing.Parse, // autorise le traitement des DTD
            XmlResolver = new XmlUrlResolver()    // résout les entités externes
        };

        using var reader = XmlReader.Create(stream, settings);
        var doc = new XmlDocument();
        doc.Load(reader);

        return Ok(doc.DocumentElement?.InnerText);
    }
}
