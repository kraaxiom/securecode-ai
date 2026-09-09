// Corrigé : Markdown-based Cross-Site Scripting (CWE-79)
// Le HTML brut est désactivé dans le pipeline Markdown pour tout contenu
// utilisateur non fiable, et le HTML généré est systématiquement passé dans
// une bibliothèque de sanitisation dédiée avant d'être renvoyé, avec une
// liste blanche stricte de schémas d'URL autorisés.
using Ganss.Xss;
using Markdig;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("wiki")]
public class WikiPageController : ControllerBase
{
    private readonly IWikiRepository _pages;
    private readonly HtmlSanitizer _sanitizer;

    // DisableHtml() empêche Markdig de laisser passer des balises HTML brutes
    // présentes dans le Markdown source (elles sont échappées comme texte).
    private readonly MarkdownPipeline _pipeline = new MarkdownPipelineBuilder()
        .UsePipeTables()
        .UseAutoLinks()
        .DisableHtml()
        .Build();

    public WikiPageController(IWikiRepository pages)
    {
        _pages = pages;

        _sanitizer = new HtmlSanitizer();
        _sanitizer.AllowedSchemes.Clear();
        _sanitizer.AllowedSchemes.Add("http");
        _sanitizer.AllowedSchemes.Add("https");
        _sanitizer.AllowedSchemes.Add("mailto");
    }

    [HttpGet("{slug}")]
    public ContentResult Render(string slug)
    {
        var page = _pages.GetBySlug(slug);

        var renderedHtml = Markdown.ToHtml(page.RawMarkdown, _pipeline);

        // Seconde couche de défense : sanitisation du HTML généré avant envoi,
        // au cas où une extension future réintroduirait du HTML actif.
        var safeHtml = _sanitizer.Sanitize(renderedHtml);

        return new ContentResult { Content = safeHtml, ContentType = "text/html" };
    }
}
