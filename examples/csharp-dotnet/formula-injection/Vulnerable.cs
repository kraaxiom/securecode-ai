// Vulnérable : Formula Injection / CSV & XLSX (CWE-1236)
// Les commentaires utilisateur sont écrits tels quels dans une feuille de
// calcul exportée ; une valeur commençant par '=' devient une formule active.
using Microsoft.AspNetCore.Mvc;
using ClosedXML.Excel;

[ApiController]
[Route("api/export")]
public class TicketExportController : ControllerBase
{
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
            sheet.Cell(row, 1).Value = ticket.Title;
            sheet.Cell(row, 2).Value = ticket.Comment;
            row++;
        }

        using var stream = new MemoryStream();
        workbook.SaveAs(stream);
        return File(stream.ToArray(), "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "tickets.xlsx");
    }
}
