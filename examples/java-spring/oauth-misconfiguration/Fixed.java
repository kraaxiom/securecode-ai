package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.security.SecureRandom;
import java.util.Base64;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;

// CORRIGÉ — la redirect_uri est validée contre une liste blanche stricte
// d'URIs enregistrées, et un paramètre "state" cryptographiquement
// aléatoire est généré à l'étape /authorize, stocké en session, puis
// obligatoirement vérifié au retour du callback avant tout échange de code.
@Controller
public class Fixed {

    private static final Set<String> ALLOWED_REDIRECT_URIS = Set.of(
            "https://app.example.com/oauth/callback"
    );

    private final SecureRandom random = new SecureRandom();
    private final ConcurrentHashMap<String, String> pendingStates = new ConcurrentHashMap<>(); // state -> sessionId

    @GetMapping("/oauth/authorize")
    public String authorize(@RequestParam String redirectUri, @RequestParam String sessionId) {
        if (!ALLOWED_REDIRECT_URIS.contains(redirectUri)) {
            throw new IllegalArgumentException("redirect_uri non autorisée");
        }

        String state = generateState();
        pendingStates.put(state, sessionId); // liaison state <-> session courante

        String authUrl = "https://provider.example/oauth/authorize"
                + "?client_id=abc123"
                + "&redirect_uri=" + redirectUri
                + "&response_type=code"
                + "&state=" + state;
        return "redirect:" + authUrl;
    }

    @GetMapping("/oauth/callback")
    @ResponseBody
    public String callback(@RequestParam String code, @RequestParam String state, @RequestParam String sessionId) {
        String expectedSession = pendingStates.remove(state); // usage unique
        if (expectedSession == null || !expectedSession.equals(sessionId)) {
            // "state" absent, inconnu, ou ne correspondant pas à la session
            // d'origine : rejet, protection contre le login CSRF.
            throw new IllegalStateException("state invalide ou expiré");
        }
        String accessToken = exchangeCodeForToken(code);
        return "connecté avec le token : " + accessToken;
    }

    private String generateState() {
        byte[] bytes = new byte[32];
        random.nextBytes(bytes);
        return Base64.getUrlEncoder().withoutPadding().encodeToString(bytes);
    }

    private String exchangeCodeForToken(String code) { return "token-" + code; }
}
