// Corrigé : CSV Injection / Formula Injection (CWE-1236)
// Toute cellule commençant par un caractère déclencheur de formule est
// préfixée par une apostrophe, neutralisant son interprétation par le tableur.
using Microsoft.AspNetCore.Mvc;
using System.Text;

[ApiController]
[Route("api/export")]
public class ExportController : ControllerBase
{
    private static readonly char[] FormulaTriggers = { '=', '+', '-', '@', '\t', '\r' };

    private static string SanitizeCsvCell(string value)
    {
        if (!string.IsNullOrEmpty(value) && FormulaTriggers.Contains(value[0]))
            return "'" + value;
        return value;
    }

    [HttpGet("customers.csv")]
    public IActionResult ExportCustomers([FromServices] ICustomerRepository repo)
    {
        var sb = new StringBuilder();
        sb.AppendLine("Name,Comment");
        foreach (var customer in repo.GetAll())
            sb.AppendLine($"{SanitizeCsvCell(customer.Name)},{SanitizeCsvCell(customer.Comment)}");

        return File(Encoding.UTF8.GetBytes(sb.ToString()), "text/csv", "customers.csv");
    }
}
