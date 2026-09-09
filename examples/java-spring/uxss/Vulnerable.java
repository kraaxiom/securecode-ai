package com.example.demo.controller;

import org.springframework.http.HttpHeaders;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — Universal XSS / UXSS (CWE-79)
// L'application embarque un widget tiers dans une <iframe> sans aucune
// restriction (pas de sandbox, pas de CSP frame-ancestors/frame-src), et
// n'envoie pas d'en-têtes limitant l'intégration inter-origines. Une faille
// dans le moteur de rendu ou dans le widget tiers embarqué peut alors
// permettre à un contenu contrôlé par l'attaquant de s'exécuter dans le
// contexte d'origine de l'application elle-même (évasion de la sandbox).
@Controller
public class Vulnerable {

    @GetMapping("/dashboard")
    @ResponseBody
    public ResponseEntity<String> dashboard() {
        String html = "<html><body>"
                + "<iframe src=\"https://widget-tiers.example/embed\"></iframe>"
                // Pas d'attribut sandbox, pas de restriction allow=...
                + "</body></html>";
        return ResponseEntity.ok()
                // Aucun en-tête CSP ni X-Frame-Options : l'application elle-même
                // peut aussi être embarquée par un tiers malveillant (clickjacking
                // combiné à UXSS).
                .body(html);
    }
}
