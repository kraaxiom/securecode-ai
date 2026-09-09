package com.example.demo.controller;

import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// CORRIGÉ — application des recommandations NIST SP 800-63B : longueur
// minimale substantielle (12+ caractères) plutôt que règles de complexité
// artificielles, rejet des mots de passe figurant dans une liste de fuites
// / mots de passe communs, et hachage systématique avec un algorithme
// dédié (BCrypt) avant stockage — jamais de mot de passe en clair.
@Controller
public class Fixed {

    private static final int MIN_LENGTH = 12;

    private final UserRepository users;
    private final WeakPasswordList weakPasswordList;
    private final BCryptPasswordEncoder encoder = new BCryptPasswordEncoder(12);

    public Fixed(UserRepository users, WeakPasswordList weakPasswordList) {
        this.users = users;
        this.weakPasswordList = weakPasswordList;
    }

    @PostMapping("/register")
    @ResponseBody
    public String register(@RequestParam String username, @RequestParam String password) {
        if (password.length() < MIN_LENGTH) {
            return "mot_de_passe_trop_court";
        }
        if (weakPasswordList.isCommon(password)) {
            return "mot_de_passe_trop_commun";
        }
        users.create(username, encoder.encode(password)); // jamais stocké en clair
        return "OK";
    }

    interface UserRepository { void create(String username, String hashedPassword); }
    interface WeakPasswordList { boolean isCommon(String password); }
}
