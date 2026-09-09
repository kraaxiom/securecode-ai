package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — Brute Force (CWE-307 : Improper Restriction of Excessive
// Authentication Attempts)
// Aucune limitation du nombre de tentatives, aucun verrouillage de compte,
// aucun délai progressif : un attaquant peut essayer des milliers de mots
// de passe par seconde sur le même compte sans être ralenti ni bloqué.
@Controller
public class Vulnerable {

    private final UserRepository users;

    public Vulnerable(UserRepository users) {
        this.users = users;
    }

    @PostMapping("/login")
    @ResponseBody
    public String login(@RequestParam String username, @RequestParam String password) {
        User user = users.findByUsername(username);
        if (user != null && user.checkPassword(password)) {
            return "OK";
        }
        // Aucun compteur d'échecs, aucun verrouillage, réponse immédiate :
        // le brute force n'est ralenti par rien côté serveur.
        return "Identifiants invalides";
    }

    interface UserRepository { User findByUsername(String username); }
    interface User { boolean checkPassword(String password); }
}
