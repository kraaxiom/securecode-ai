package com.example.demo.controller;

import org.jsoup.Jsoup;
import org.jsoup.safety.Safelist;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — Mutation XSS / mXSS (CWE-79)
// Le contenu est nettoyé une seule fois côté serveur avec une liste
// blanche trop permissive (balises de formatage "brutes" incluant des
// structures ambiguës), PUIS le résultat est réinjecté tel quel dans le
// DOM via innerHTML côté client. Le navigateur "répare" un HTML mal formé
// pendant le parsing (mutation), ce qui peut reconstituer un payload actif
// qui n'existait pas sous cette forme au moment de la sanitisation.
@Controller
public class Vulnerable {

    // Safelist "relaxed" : autorise des structures imbriquées ambiguës
    // (ex: balises noscript/svg/math, attributs de style) propices à la
    // mutation lors du re-parsing par le navigateur.
    private final Safelist permissiveList = Safelist.relaxed()
            .addTags("noscript", "svg", "math")
            .addAttributes(":all", "style");

    @PostMapping("/profile/bio")
    public String updateBio(@RequestParam String bio, Model model) {
        String sanitizedOnce = Jsoup.clean(bio, permissiveList);
        // Stocké et renvoyé tel quel ; côté template : th:utext + le
        // JavaScript client réinjecte encore via innerHTML pour l'aperçu
        // instantané, provoquant une seconde passe de parsing (mutation).
        model.addAttribute("bioHtml", sanitizedOnce);
        return "profile-preview";
    }
}
