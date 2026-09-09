package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — OAuth Misconfiguration (CWE-287 : Improper Authentication)
// Le callback OAuth accepte n'importe quelle redirect_uri fournie par le
// client (pas de liste blanche stricte), et le paramètre "state" (censé
// prévenir le CSRF sur le flux OAuth) n'est ni généré aléatoirement ni
// vérifié au retour. Un attaquant peut rediriger le code d'autorisation
// vers un domaine qu'il contrôle, ou forcer une victime à se connecter
// avec le compte OAuth de l'attaquant (login CSRF).
@Controller
public class Vulnerable {

    @GetMapping("/oauth/authorize")
    public String authorize(@RequestParam String redirectUri, @RequestParam(required = false) String state) {
        // Aucune validation de redirectUri contre une liste blanche connue
        String authUrl = "https://provider.example/oauth/authorize"
                + "?client_id=abc123"
                + "&redirect_uri=" + redirectUri // accepté tel quel
                + "&response_type=code";
        return "redirect:" + authUrl;
    }

    @GetMapping("/oauth/callback")
    @ResponseBody
    public String callback(@RequestParam String code, @RequestParam(required = false) String state) {
        // "state" reçu n'est jamais comparé à une valeur générée et stockée
        // en session lors de l'étape /authorize : le CSRF sur le flux OAuth
        // n'est pas empêché.
        String accessToken = exchangeCodeForToken(code);
        return "connecté avec le token : " + accessToken;
    }

    private String exchangeCodeForToken(String code) { return "token-" + code; }
}
