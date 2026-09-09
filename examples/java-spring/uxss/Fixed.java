package com.example.demo.controller;

import org.springframework.http.HttpHeaders;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// CORRIGÉ — le contenu tiers est isolé via l'attribut sandbox (sans
// allow-scripts ni allow-same-origin combinés, ce qui empêcherait toute
// évasion vers l'origine parente), et l'application envoie des en-têtes
// CSP stricts (frame-src limité au domaine de confiance, frame-ancestors
// 'none' pour se protéger elle-même du clickjacking/UXSS inverse).
@Controller
public class Fixed {

    @GetMapping("/dashboard")
    @ResponseBody
    public ResponseEntity<String> dashboard() {
        String html = "<html><body>"
                // sandbox restreint : scripts autorisés mais SANS allow-same-origin,
                // ce qui empêche le contenu embarqué d'accéder à l'origine parente
                // même en cas de faille dans le widget tiers.
                + "<iframe src=\"https://widget-tiers.example/embed\" sandbox=\"allow-scripts\"></iframe>"
                + "</body></html>";
        return ResponseEntity.ok()
                .header(HttpHeaders.CONTENT_SECURITY_POLICY,
                        "default-src 'self'; frame-src https://widget-tiers.example; frame-ancestors 'none'")
                .header("X-Frame-Options", "DENY")
                .body(html);
    }
}
