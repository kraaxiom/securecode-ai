// Corrigé : Formula Injection / CSV & XLSX (CWE-1236)
// Toute valeur commençant par un caractère déclencheur de formule est
// préfixée par une apostrophe avant d'être écrite dans la cellule.
using Microsoft.AspNetCore.Mvc;
using ClosedXML.Excel;

[ApiController]
[Route("api/export")]
public class TicketExportController : ControllerBase
{
    private static readonly char[] FormulaTriggers = { '=', '+', '-', '@', '\t', '\r' };

    private static string NeutraliserFormule(string valeur)
    {
        if (!string.IsNullOrEmpty(valeur) && FormulaTriggers.Contains(valeur[0]))
            return "'" + valeur;
        return valeur;
    }

    [HttpGet("tickets.xlsx")]
    public IActionResult ExportTickets([FromServices] ITicketRepository repo)
    {
        using var workbook = new XLWorkbook();
        var sheet = workbook.Worksheets.Add("Tickets");
        sheet.Cell(1, 1).Value = "Titre";
        sheet.Cell(1, 2).Value = "Commentaire";

        var row = 2;
        foreach (var ticket in repo.GetAll())
        {
            sheet.Cell(row, 1).Value = NeutraliserFormule(ticket.Title);
            sheet.Cell(row, 2).Value = NeutraliserFormule(ticket.Comment);
            row++;
        }

        using var stream = new MemoryStream();
        workbook.SaveAs(stream);
        return File(stream.ToArray(), "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "tickets.xlsx");
    }
}
