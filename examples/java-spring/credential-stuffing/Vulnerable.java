package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — Credential Stuffing (CWE-307)
// Le point de connexion accepte un volume illimité de tentatives, sans
// détection de vitesse anormale, sans MFA, et sans vérification des
// identifiants contre une base de mots de passe compromis. Un attaquant
// peut ainsi rejouer automatiquement des paires identifiant/mot de passe
// fuitées d'autres services contre CE service, sans friction.
@Controller
public class Vulnerable {

    private final UserRepository users;

    public Vulnerable(UserRepository users) {
        this.users = users;
    }

    @PostMapping("/api/login")
    @ResponseBody
    public String login(@RequestParam String email, @RequestParam String password) {
        User user = users.findByEmail(email);
        if (user != null && user.checkPassword(password)) {
            return issueSessionToken(user); // pas de MFA, pas de contrôle de vélocité
        }
        return "invalid";
    }

    private String issueSessionToken(User user) { return "token-for-" + user.hashCode(); }

    interface UserRepository { User findByEmail(String email); }
    interface User { boolean checkPassword(String password); }
}
