package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;

// VULNÉRABLE — Reflected XSS (CWE-79)
// Le paramètre de recherche fourni par le client est renvoyé directement
// dans la réponse HTML, sans échappement, via th:utext. Un lien piégé
// (ex: /search?q=<script>...</script>) exécute le script dès que la
// victime clique dessus : aucune persistance nécessaire.
@Controller
public class Vulnerable {

    @GetMapping("/search")
    public String search(@RequestParam(required = false, defaultValue = "") String q, Model model) {
        // Réinjecté tel quel dans le template : "Vous avez cherché : " + q
        model.addAttribute("query", q); // rendu via th:utext côté vue
        return "search-results";
    }
}
