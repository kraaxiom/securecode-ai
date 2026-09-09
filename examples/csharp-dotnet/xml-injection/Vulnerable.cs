// Vulnérable : XML Injection (CWE-91)
// Le document XML est construit par concaténation de chaînes incluant une
// entrée utilisateur non échappée, permettant d'altérer la structure du
// document (ajout de nœuds, falsification de valeurs).
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/export")]
public class UserExportController : ControllerBase
{
    [HttpPost]
    public IActionResult ExportUser([FromForm] string name, [FromForm] string role)
    {
        // "name" contenant "</name><role>admin</role><name>" falsifie le
        // document produit.
        var xml = $"<user><name>{name}</name><role>{role}</role></user>";
        return Content(xml, "application/xml");
    }
}
