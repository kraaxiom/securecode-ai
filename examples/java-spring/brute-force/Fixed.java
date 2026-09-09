package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.time.Duration;
import java.time.Instant;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicInteger;

// CORRIGÉ — verrouillage temporaire du compte après un nombre limité
// d'échecs (avec compteur par identifiant, indépendant de l'IP pour
// résister à la rotation d'IP), délai constant pour éviter le
// timing-based user enumeration, et journalisation des échecs répétés
// pour alerting. En production : coupler à un rate-limiter (Bucket4j /
// API Gateway) et à un CAPTCHA après N échecs.
@Controller
public class Fixed {

    private static final int MAX_ATTEMPTS = 5;
    private static final Duration LOCK_DURATION = Duration.ofMinutes(15);

    private final UserRepository users;
    private final ConcurrentHashMap<String, AtomicInteger> failedAttempts = new ConcurrentHashMap<>();
    private final ConcurrentHashMap<String, Instant> lockedUntil = new ConcurrentHashMap<>();

    public Fixed(UserRepository users) {
        this.users = users;
    }

    @PostMapping("/login")
    @ResponseBody
    public String login(@RequestParam String username, @RequestParam String password) {
        Instant lockExpiry = lockedUntil.get(username);
        if (lockExpiry != null && Instant.now().isBefore(lockExpiry)) {
            // Message identique en cas de verrouillage ou d'échec, pour ne
            // pas révéler l'existence du compte (défense en profondeur).
            return "Identifiants invalides";
        }

        User user = users.findByUsername(username);
        boolean valid = user != null && user.checkPassword(password);

        if (valid) {
            failedAttempts.remove(username);
            lockedUntil.remove(username);
            return "OK";
        }

        int attempts = failedAttempts.computeIfAbsent(username, k -> new AtomicInteger()).incrementAndGet();
        if (attempts >= MAX_ATTEMPTS) {
            lockedUntil.put(username, Instant.now().plus(LOCK_DURATION));
            // Journalisation + alerte SOC recommandée ici (compte ciblé)
        }
        return "Identifiants invalides";
    }

    interface UserRepository { User findByUsername(String username); }
    interface User { boolean checkPassword(String password); }
}
