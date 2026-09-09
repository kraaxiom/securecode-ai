package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — Mass Assignment (CWE-915 : Improperly Controlled
// Modification of Dynamically-Determined Object Attributes)
// L'entité User complète (y compris ses champs sensibles "role" et
// "isAdmin") est directement désérialisée depuis le corps de la requête
// puis persistée. Un attaquant peut ajouter "role":"ADMIN" au JSON envoyé
// et s'auto-promouvoir administrateur, alors que le formulaire prévu ne
// propose que "name" et "email".
@Controller
public class Vulnerable {

    private final UserRepository users;

    public Vulnerable(UserRepository users) {
        this.users = users;
    }

    @PutMapping("/api/users/{id}")
    @ResponseBody
    public User updateUser(@PathVariable Long id, @RequestBody User update) {
        // update peut contenir n'importe quel champ de l'entité, y compris
        // "role" ou "isAdmin", jamais filtré avant persistance.
        return users.save(id, update);
    }

    interface UserRepository { User save(Long id, User user); }

    static class User {
        Long id;
        String name;
        String email;
        String role;    // ex: "USER" / "ADMIN" — ne devrait jamais être modifiable par ce endpoint
        boolean isAdmin; // idem
    }
}
