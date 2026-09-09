package com.example.demo.controller;

import org.jsoup.Jsoup;
import org.jsoup.safety.Safelist;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;

// CORRIGÉ — utilisation d'une liste blanche minimale et non ambiguë
// (Safelist.basic, sans style ni balises "conteneurs" génériques comme
// svg/math/noscript), et surtout : la sanitisation est effectuée à
// nouveau juste avant toute insertion DOM côté client (DOMPurify côté JS,
// avec SANITIZE_DOM activé) — jamais de confiance dans un HTML "déjà
// nettoyé" qui retraverse un pipeline de parsing HTML.
@Controller
public class Fixed {

    // Liste blanche restrictive : aucune balise structurelle ambiguë,
    // aucun attribut de style, source unique de vérité pour le format autorisé.
    private final Safelist strictList = Safelist.basic()
            .removeTags("a") // pas de liens dans une bio, réduit encore la surface
            .removeAttributes(":all", "style");

    @PostMapping("/profile/bio")
    public String updateBio(@RequestParam String bio, Model model) {
        String sanitized = Jsoup.clean(bio, strictList);
        model.addAttribute("bioHtml", sanitized);
        // Côté client, l'aperçu réutilise DOMPurify.sanitize(html) juste
        // avant l'affectation à innerHTML, pour neutraliser toute mutation
        // introduite par un re-parsing HTML ultérieur :
        //
        // preview.innerHTML = DOMPurify.sanitize(html, { SANITIZE_DOM: true });
        return "profile-preview";
    }
}
