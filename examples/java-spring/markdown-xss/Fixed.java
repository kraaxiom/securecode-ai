package com.example.demo.controller;

import org.commonmark.node.Node;
import org.commonmark.parser.Parser;
import org.commonmark.renderer.html.HtmlRenderer;
import org.owasp.html.HtmlPolicyBuilder;
import org.owasp.html.PolicyFactory;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;

// CORRIGÉ — le HTML issu du rendu Markdown est passé dans une liste
// blanche stricte de balises/attributs (OWASP Java HTML Sanitizer) avant
// d'être exposé à la vue. Le rendu de Markdown en HTML brut n'est jamais
// sûr par défaut : la sanitisation post-rendu est obligatoire.
@Controller
public class Fixed {

    private final Parser parser = Parser.builder().build();
    private final HtmlRenderer renderer = HtmlRenderer.builder().build();

    // Liste blanche : mise en forme basique uniquement, aucun script/event handler/iframe
    private final PolicyFactory sanitizer = new HtmlPolicyBuilder()
            .allowElements("p", "b", "i", "em", "strong", "a", "ul", "ol", "li", "code", "pre", "blockquote")
            .allowUrlProtocols("https")
            .allowAttributes("href").onElements("a")
            .requireRelNofollowOnLinks()
            .toFactory();

    @PostMapping("/comments")
    public String preview(@RequestParam String markdown, Model model) {
        Node document = parser.parse(markdown);
        String html = renderer.render(document);
        String safeHtml = sanitizer.sanitize(html);
        model.addAttribute("renderedHtml", safeHtml); // rendu via th:utext, désormais sûr
        return "comment-preview";
    }
}
