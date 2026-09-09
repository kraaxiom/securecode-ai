package com.example.demo.controller;

import org.commonmark.node.Node;
import org.commonmark.parser.Parser;
import org.commonmark.renderer.html.HtmlRenderer;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — Markdown XSS (CWE-79)
// Le Markdown fourni par l'utilisateur est converti en HTML puis rendu
// SANS SANITISATION. Or le Markdown autorise nativement l'injection de
// HTML brut (ex: <script>, <img onerror=...>) qui traverse le rendu tel
// quel, contournant toute intuition de "texte simple sûr".
@Controller
public class Vulnerable {

    private final Parser parser = Parser.builder().build();
    private final HtmlRenderer renderer = HtmlRenderer.builder().build();

    @PostMapping("/comments")
    public String preview(@RequestParam String markdown, Model model) {
        Node document = parser.parse(markdown);
        String html = renderer.render(document);
        // HTML issu du Markdown injecté directement, non échappé, via th:utext
        model.addAttribute("renderedHtml", html);
        return "comment-preview";
    }
}
