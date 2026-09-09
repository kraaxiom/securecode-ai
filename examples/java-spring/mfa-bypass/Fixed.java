package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// CORRIGÉ — la session reste dans un état intermédiaire "en attente de
// MFA" tant que le second facteur n'est pas vérifié, et TOUTES les routes
// protégées contrôlent explicitement isMfaVerified en plus de
// isAuthenticated. Impossible d'accéder à une ressource sensible sans
// avoir complété les deux étapes.
@Controller
public class Fixed {

    private final SessionStore sessions;

    public Fixed(SessionStore sessions) {
        this.sessions = sessions;
    }

    @PostMapping("/login")
    @ResponseBody
    public String login(@RequestParam String username, @RequestParam String password, @RequestParam String sessionId) {
        if (checkPassword(username, password)) {
            // État intermédiaire : authentifié par mot de passe MAIS pas
            // encore pleinement connecté tant que le MFA n'est pas validé.
            sessions.markPendingMfa(sessionId, username);
            return "mfa_challenge_sent";
        }
        return "invalid";
    }

    @PostMapping("/mfa/verify")
    @ResponseBody
    public String verifyMfa(@RequestParam String sessionId, @RequestParam String code) {
        if (!sessions.isPendingMfa(sessionId)) {
            return "unauthorized"; // pas de tentative de MFA sans login préalable
        }
        if (code.equals(sessions.getExpectedCode(sessionId))) {
            sessions.markFullyAuthenticated(sessionId); // seule voie vers l'état "connecté"
            return "ok";
        }
        return "invalid_code";
    }

    @GetMapping("/account/sensitive-data")
    @ResponseBody
    public String sensitiveData(@RequestParam String sessionId) {
        // Vérification explicite et obligatoire des deux facteurs sur
        // chaque route protégée, idéalement centralisée dans un filtre
        // Spring Security plutôt que dupliquée par contrôleur.
        if (!sessions.isFullyAuthenticated(sessionId)) {
            return "unauthorized";
        }
        return "données sensibles du compte";
    }

    private boolean checkPassword(String u, String p) { return true; }

    interface SessionStore {
        void markPendingMfa(String sessionId, String username);
        boolean isPendingMfa(String sessionId);
        void markFullyAuthenticated(String sessionId);
        boolean isFullyAuthenticated(String sessionId);
        String getExpectedCode(String sessionId);
    }
}
