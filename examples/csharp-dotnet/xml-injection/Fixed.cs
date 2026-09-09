// Corrigé : XML Injection (CWE-91)
// Le document est construit via l'API DOM System.Xml, qui échappe
// automatiquement le contenu textuel des nœuds : aucune concaténation de
// chaîne n'est utilisée.
using Microsoft.AspNetCore.Mvc;
using System.Xml;

[ApiController]
[Route("api/export")]
public class UserExportController : ControllerBase
{
    [HttpPost]
    public IActionResult ExportUser([FromForm] string name, [FromForm] string role)
    {
        var doc = new XmlDocument();
        var user = doc.CreateElement("user");

        var nameEl = doc.CreateElement("name");
        nameEl.InnerText = name; // échappement automatique par XmlDocument
        user.AppendChild(nameEl);

        var roleEl = doc.CreateElement("role");
        roleEl.InnerText = role; // échappement automatique par XmlDocument
        user.AppendChild(roleEl);

        doc.AppendChild(user);

        return Content(doc.OuterXml, "application/xml");
    }
}
