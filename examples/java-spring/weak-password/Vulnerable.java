package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — Weak Password Policy (CWE-521 : Weak Password Requirements)
// Aucune règle minimale n'est imposée à l'inscription : une longueur d'un
// seul caractère est acceptée, aucune vérification contre les mots de
// passe les plus communs, aucune estimation de robustesse. Les comptes
// créés sont donc triviaux à compromettre par brute force ou dictionnaire.
@Controller
public class Vulnerable {

    private final UserRepository users;

    public Vulnerable(UserRepository users) {
        this.users = users;
    }

    @PostMapping("/register")
    @ResponseBody
    public String register(@RequestParam String username, @RequestParam String password) {
        // Aucune contrainte de longueur, de complexité, ni de liste noire
        users.create(username, password); // stocké sans même être validé
        return "OK";
    }

    interface UserRepository { void create(String username, String password); }
}
