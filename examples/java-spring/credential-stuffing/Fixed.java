package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// CORRIGÉ — vérification du mot de passe contre une base de fuites
// connues (type HaveIBeenPwned k-anonymity API) au moment de la
// connexion réussie pour forcer un changement, exigence de MFA pour
// toute connexion, et détection de vélocité (nombreuses tentatives sur
// des comptes différents depuis une même source = signal de credential
// stuffing distinct du brute force classique ciblé sur un seul compte).
@Controller
public class Fixed {

    private final UserRepository users;
    private final BreachedPasswordChecker breachChecker;
    private final VelocityGuard velocityGuard;

    public Fixed(UserRepository users, BreachedPasswordChecker breachChecker, VelocityGuard velocityGuard) {
        this.users = users;
        this.breachChecker = breachChecker;
        this.velocityGuard = velocityGuard;
    }

    @PostMapping("/api/login")
    @ResponseBody
    public String login(@RequestParam String email, @RequestParam String password,
                         @RequestHeader("X-Forwarded-For") String sourceIp) {
        // Détecte un nombre anormal de tentatives sur des comptes DIFFÉRENTS
        // depuis la même source : signature typique du credential stuffing
        // (par opposition au brute force ciblé sur un seul compte).
        if (velocityGuard.isSuspicious(sourceIp)) {
            return "temporarily_blocked";
        }

        User user = users.findByEmail(email);
        if (user == null || !user.checkPassword(password)) {
            velocityGuard.recordFailure(sourceIp, email);
            return "invalid";
        }

        if (!user.hasMfaVerified()) {
            // MFA obligatoire : même avec des identifiants valides (issus
            // d'une fuite tierce), l'attaquant ne peut pas terminer la
            // connexion sans le second facteur.
            return "mfa_required";
        }

        if (breachChecker.isCompromised(password)) {
            // Le mot de passe, bien que correct, apparaît dans une base de
            // fuites connues : on force son renouvellement immédiat.
            return "password_reset_required";
        }

        return "token-for-" + user.hashCode();
    }

    interface UserRepository { User findByEmail(String email); }
    interface User { boolean checkPassword(String password); boolean hasMfaVerified(); }
    interface BreachedPasswordChecker { boolean isCompromised(String password); }
    interface VelocityGuard {
        boolean isSuspicious(String sourceIp);
        void recordFailure(String sourceIp, String email);
    }
}
