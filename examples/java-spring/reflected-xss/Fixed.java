package com.example.demo.controller;

import org.owasp.encoder.Encode;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;

// CORRIGÉ — le paramètre est échappé pour le contexte HTML avant d'être
// exposé au modèle, et le template utilise th:text (échappement
// automatique) au lieu de th:utext. La combinaison des deux constitue une
// défense en profondeur contre tout contournement d'un des deux niveaux.
@Controller
public class Fixed {

    @GetMapping("/search")
    public String search(@RequestParam(required = false, defaultValue = "") String q, Model model) {
        model.addAttribute("query", Encode.forHtml(q)); // rendu via th:text côté vue
        return "search-results";
    }
}
