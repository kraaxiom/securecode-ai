// Vulnérable : Indirect Prompt Injection (CWE-1427)
// Le contenu d'une page web externe est récupéré puis injecté tel quel
// dans le contexte du modèle, aux côtés des outils actifs de l'agent, sans
// marquage de provenance ni séparation entre contenu à analyser et
// instructions à exécuter.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/webagent")]
public class WebAgentController : ControllerBase
{
    private readonly IWebFetcher _fetcher;
    private readonly ILlmClient _llm;
    private readonly IToolRegistry _tools;

    public WebAgentController(IWebFetcher fetcher, ILlmClient llm, IToolRegistry tools)
    {
        _fetcher = fetcher;
        _llm = llm;
        _tools = tools;
    }

    [HttpPost("summarize-and-act")]
    public async Task<IActionResult> SummarizeAndAct([FromBody] WebAgentRequest request)
    {
        var pageContent = await _fetcher.FetchAsync(request.Url);

        // Le contenu externe est concaténé directement au prompt, sans
        // délimiteur ni indication qu'il s'agit de données non fiables.
        // L'agent conserve simultanément l'accès à tous ses outils actifs.
        var prompt = $"Résume ce contenu et effectue les actions pertinentes :\n{pageContent}";

        var response = await _llm.CompleteWithToolsAsync(prompt, _tools.All);
        return Ok(response);
    }
}
