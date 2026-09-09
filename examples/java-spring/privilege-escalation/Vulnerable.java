package com.example.demo.controller;

import org.springframework.security.core.Authentication;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — Privilege Escalation (CWE-269 : Improper Privilege
// Management)
// L'endpoint de mise à jour de profil permet à l'utilisateur de modifier
// son propre champ "role" (utile en apparence pour, par exemple, choisir
// un statut "vendeur" vs "acheteur"), sans jamais restreindre les valeurs
// acceptées ni vérifier le rôle courant de l'appelant. Un utilisateur
// standard peut ainsi s'attribuer le rôle "ADMIN" lui-même.
@Controller
public class Vulnerable {

    private final UserRepository users;

    public Vulnerable(UserRepository users) {
        this.users = users;
    }

    @PatchMapping("/api/profile/role")
    @ResponseBody
    public String updateRole(@RequestParam String role, Authentication auth) {
        Long currentUserId = (Long) auth.getPrincipal();
        // Aucune liste blanche de rôles auto-attribuables, aucune vérification
        // que le rôle courant de l'appelant l'autorise à faire ce changement.
        users.setRole(currentUserId, role); // role peut valoir "ADMIN"
        return "rôle mis à jour";
    }

    interface UserRepository { void setRole(Long userId, String role); }
}
