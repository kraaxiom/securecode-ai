package com.example.demo.controller;

import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// CORRIGÉ — le rôle requis pour cette fonction est vérifié explicitement
// via @PreAuthorize (Spring Security), appliqué de manière déclarative et
// centralisée sur chaque endpoint sensible plutôt que dans la logique
// métier, réduisant le risque d'oubli.
@Controller
public class Fixed {

    private final UserRepository users;

    public Fixed(UserRepository users) {
        this.users = users;
    }

    @PostMapping("/api/admin/users/{id}/delete")
    @ResponseBody
    @PreAuthorize("hasRole('ADMIN')") // fonction explicitement réservée au rôle ADMIN
    public String deleteUser(@PathVariable Long id) {
        users.delete(id);
        return "utilisateur supprimé";
    }

    interface UserRepository { void delete(Long id); }
}
