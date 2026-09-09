package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — MFA Bypass (CWE-287 : Improper Authentication)
// L'étape MFA est un endpoint séparé, appelé APRÈS l'authentification par
// mot de passe, mais la session est déjà considérée "connectée" avant la
// vérification du code MFA (le flag isMfaVerified n'est jamais contrôlé
// sur les routes protégées). Un attaquant qui connaît le mot de passe
// peut donc accéder directement aux ressources protégées en sautant
// l'étape /mfa/verify.
@Controller
public class Vulnerable {

    private final SessionStore sessions;

    public Vulnerable(SessionStore sessions) {
        this.sessions = sessions;
    }

    @PostMapping("/login")
    @ResponseBody
    public String login(@RequestParam String username, @RequestParam String password, @RequestParam String sessionId) {
        if (checkPassword(username, password)) {
            sessions.markAuthenticated(sessionId, username); // session pleinement active dès ce point
            return "mfa_challenge_sent";
        }
        return "invalid";
    }

    @PostMapping("/mfa/verify")
    @ResponseBody
    public String verifyMfa(@RequestParam String sessionId, @RequestParam String code) {
        // Étape optionnelle en pratique : les routes protégées ne vérifient
        // que sessions.isAuthenticated(sessionId), jamais un flag MFA dédié.
        if (code.equals(sessions.getExpectedCode(sessionId))) {
            sessions.markMfaVerified(sessionId);
        }
        return "ok";
    }

    @GetMapping("/account/sensitive-data")
    @ResponseBody
    public String sensitiveData(@RequestParam String sessionId) {
        if (!sessions.isAuthenticated(sessionId)) { // ne vérifie pas isMfaVerified !
            return "unauthorized";
        }
        return "données sensibles du compte";
    }

    private boolean checkPassword(String u, String p) { return true; }

    interface SessionStore {
        void markAuthenticated(String sessionId, String username);
        void markMfaVerified(String sessionId);
        boolean isAuthenticated(String sessionId);
        boolean isMfaVerified(String sessionId);
        String getExpectedCode(String sessionId);
    }
}
