// Vulnérable : CSV Injection / Formula Injection (CWE-1236)
// Les champs utilisateur sont écrits tels quels dans le fichier CSV exporté.
// Une valeur commençant par '=' sera interprétée comme une formule par le
// tableur à l'ouverture du fichier.
using Microsoft.AspNetCore.Mvc;
using System.Text;

[ApiController]
[Route("api/export")]
public class ExportController : ControllerBase
{
    [HttpGet("customers.csv")]
    public IActionResult ExportCustomers([FromServices] ICustomerRepository repo)
    {
        var sb = new StringBuilder();
        sb.AppendLine("Name,Comment");
        foreach (var customer in repo.GetAll())
            sb.AppendLine($"{customer.Name},{customer.Comment}");

        return File(Encoding.UTF8.GetBytes(sb.ToString()), "text/csv", "customers.csv");
    }
}
