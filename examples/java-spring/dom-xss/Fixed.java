package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

// CORRIGÉ — le script livré par le serveur utilise désormais un sink sûr
// (textContent) au lieu d'un sink dangereux (innerHTML), et une politique
// CSP stricte est envoyée par le serveur pour bloquer toute exécution de
// script inline non conforme, en défense en profondeur.
@Controller
public class Fixed {

    @GetMapping("/search")
    public String searchPage(Model model) {
        // La page renvoie désormais :
        //
        // <div id="results"></div>
        // <script>
        //   const params = new URLSearchParams(window.location.search);
        //   const q = params.get('q');
        //   // Sink sûr : textContent n'interprète jamais le contenu comme du HTML
        //   document.getElementById('results').textContent = 'Résultats pour : ' + q;
        // </script>
        //
        // Complément défense en profondeur : en-tête CSP appliqué globalement
        // via un filtre Spring Security :
        // .headers(h -> h.contentSecurityPolicy(csp -> csp
        //     .policyDirectives("default-src 'self'; script-src 'self'; object-src 'none'")))
        return "search";
    }
}
