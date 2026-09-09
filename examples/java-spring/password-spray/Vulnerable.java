package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — Password Spray (CWE-307)
// Le password spraying teste UN mot de passe très répandu (ex: "Ete2026!")
// contre DE NOMBREUX comptes différents, précisément pour rester sous les
// seuils d'un éventuel verrouillage "par compte". Ici, aucune détection
// globale (multi-comptes, même mot de passe, même source) n'existe : seul
// un verrouillage par compte individuel serait de toute façon insuffisant.
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
        // Aucune corrélation entre tentatives sur des comptes différents,
        // aucune détection de mot de passe "populaire" testé en masse.
        return "invalid";
    }

    interface UserRepository { User findByUsername(String username); }
    interface User { boolean checkPassword(String password); }
}
