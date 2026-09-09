// Corrigé : XML External Entity (XXE) Injection (CWE-611)
// Le traitement des DTD est interdit (DtdProcessing.Prohibit) et aucun
// résolveur d'entités externes n'est fourni (XmlResolver = null), ce qui
// désactive à la fois les déclarations DOCTYPE et la résolution d'entités
// externes générales/paramétriques.
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
            DtdProcessing = DtdProcessing.Prohibit, // rejette toute déclaration DOCTYPE
            XmlResolver = null,                      // aucune résolution d'entité externe
            MaxCharactersFromEntities = 0             // protection supplémentaire contre l'expansion
        };

        using var reader = XmlReader.Create(stream, settings);
        var doc = new XmlDocument { XmlResolver = null };
        doc.Load(reader);

        return Ok(doc.DocumentElement?.InnerText);
    }
}
