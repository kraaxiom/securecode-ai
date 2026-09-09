// Vulnérable : Markdown-based Cross-Site Scripting (CWE-79)
// Le contenu Markdown soumis par l'utilisateur (README, commentaire de wiki) est
// converti en HTML avec le support du HTML brut activé, puis renvoyé sans
// sanitisation. Le moteur de rendu laisse passer des balises actives et des
// schémas d'URL dangereux (javascript:) présents dans le Markdown source.
using Markdig;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("wiki")]
public class WikiPageController : ControllerBase
{
    private readonly IWikiRepository _pages;

    // Pipeline Markdig avec le HTML brut autorisé dans le document source.
    private readonly MarkdownPipeline _pipeline = new MarkdownPipelineBuilder()
        .UsePipeTables()
        .UseAutoLinks()
        .Build();

    public WikiPageController(IWikiRepository pages)
    {
        _pages = pages;
    }

    [HttpGet("{slug}")]
    public ContentResult Render(string slug)
    {
        var page = _pages.GetBySlug(slug);

        // Le HTML brut inclus dans le Markdown (balises <script>, <img onerror>)
        // est conservé tel quel par Markdig et renvoyé directement au navigateur.
        var renderedHtml = Markdown.ToHtml(page.RawMarkdown, _pipeline);

        return new ContentResult { Content = renderedHtml, ContentType = "text/html" };
    }
}
