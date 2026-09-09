package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

// VULNÉRABLE — DOM-based XSS (CWE-79)
// Le serveur ne fait ici que servir une page contenant du JavaScript qui,
// côté client, lit un paramètre d'URL et l'injecte dans le DOM via
// innerHTML sans aucune sanitisation. La vulnérabilité vit entièrement
// côté navigateur, mais le gabarit livré par le serveur en est la cause.
@Controller
public class Vulnerable {

    @GetMapping("/search")
    public String searchPage(Model model) {
        // La page renvoie ce script tel quel au navigateur :
        //
        // <div id="results"></div>
        // <script>
        //   const params = new URLSearchParams(window.location.search);
        //   const q = params.get('q');
        //   // Source non fiable (URL) -> Sink dangereux (innerHTML)
        //   document.getElementById('results').innerHTML =
        //       'Résultats pour : ' + q;
        // </script>
        return "search"; // template Thymeleaf statique contenant le script ci-dessus
    }
}
