package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — Broken Function Level Authorization (BFLA) (CWE-862 :
// Missing Authorization)
// L'endpoint d'administration n'exige qu'une authentification valide (un
// utilisateur connecté quelconque), sans jamais vérifier que celui-ci
// possède le RÔLE administrateur. N'importe quel utilisateur authentifié
// peut donc appeler une fonction censée être réservée aux admins.
@Controller
public class Vulnerable {

    private final UserRepository users;

    public Vulnerable(UserRepository users) {
        this.users = users;
    }

    @PostMapping("/api/admin/users/{id}/delete")
    @ResponseBody
    public String deleteUser(@PathVariable Long id, @RequestAttribute("currentUserId") Long currentUserId) {
        // Seule l'authentification est vérifiée en amont (filtre requireAuth),
        // aucun contrôle de rôle n'est effectué ici : n'importe quel
        // utilisateur connecté peut supprimer n'importe quel compte.
        users.delete(id);
        return "utilisateur supprimé";
    }

    interface UserRepository { void delete(Long id); }
}
