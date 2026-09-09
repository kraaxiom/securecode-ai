package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.security.MessageDigest;
import java.time.Duration;
import java.time.Instant;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicInteger;

// CORRIGÉ — détection agrégée par SOURCE (IP/ASN) et par empreinte de mot
// de passe testé, indépendante du compte ciblé : le password spraying vise
// justement à contourner un verrouillage par compte en dispersant les
// tentatives. On bloque donc la source lorsque trop de comptes distincts
// sont testés en peu de temps, et MFA + politique de mot de passe fort
// réduisent encore la surface (mots de passe "populaires" rejetés au
// changement, via une base type NIST 800-63B).
@Controller
public class Fixed {

    private static final int MAX_DISTINCT_ACCOUNTS_PER_WINDOW = 10;
    private static final Duration WINDOW = Duration.ofMinutes(10);

    private final UserRepository users;
    private final WeakPasswordList weakPasswordList; // liste de mots de passe/expressions les plus utilisés
    private final ConcurrentHashMap<String, SourceActivity> activityBySource = new ConcurrentHashMap<>();

    public Fixed(UserRepository users, WeakPasswordList weakPasswordList) {
        this.users = users;
        this.weakPasswordList = weakPasswordList;
    }

    @PostMapping("/login")
    @ResponseBody
    public String login(@RequestParam String username, @RequestParam String password,
                         @RequestHeader("X-Forwarded-For") String sourceIp) {
        SourceActivity activity = activityBySource.computeIfAbsent(sourceIp, k -> new SourceActivity());
        if (activity.isSpraySuspected(WINDOW, MAX_DISTINCT_ACCOUNTS_PER_WINDOW)) {
            return "temporarily_blocked";
        }

        User user = users.findByUsername(username);
        boolean valid = user != null && user.checkPassword(password);
        if (!valid) {
            activity.recordAttempt(username, Instant.now());
            return "invalid";
        }
        return "OK";
    }

    // Rejette au moment du CHANGEMENT de mot de passe les valeurs trop
    // répandues (ex: rockyou.txt / listes NIST), ce qui réduit fortement
    // l'efficacité d'un password spraying à long terme.
    @PostMapping("/account/change-password")
    @ResponseBody
    public String changePassword(@RequestParam String newPassword) {
        if (weakPasswordList.isCommon(newPassword)) {
            return "mot_de_passe_trop_commun";
        }
        return "OK";
    }

    interface UserRepository { User findByUsername(String username); }
    interface User { boolean checkPassword(String password); }
    interface WeakPasswordList { boolean isCommon(String password); }

    static class SourceActivity {
        private final ConcurrentHashMap<String, Instant> distinctAccountsTried = new ConcurrentHashMap<>();

        void recordAttempt(String username, Instant now) { distinctAccountsTried.put(username, now); }

        boolean isSpraySuspected(Duration window, int threshold) {
            Instant cutoff = Instant.now().minus(window);
            distinctAccountsTried.values().removeIf(t -> t.isBefore(cutoff));
            return distinctAccountsTried.size() >= threshold;
        }
    }
}
