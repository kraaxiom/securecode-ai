package com.example.demo.controller;

import org.springframework.security.core.Authentication;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.util.Set;

// CORRIGÉ — les rôles auto-attribuables par un utilisateur standard sont
// restreints à une liste blanche métier explicite qui EXCLUT tout rôle
// privilégié ("ADMIN" notamment), et l'élévation vers un rôle privilégié
// n'est possible que via un endpoint distinct, réservé aux administrateurs
// existants (@PreAuthorize dédié).
@Controller
public class Fixed {

    // Rôles qu'un utilisateur peut légitimement s'attribuer lui-même :
    // aucun rôle privilégié n'y figure.
    private static final Set<String> SELF_ASSIGNABLE_ROLES = Set.of("ACHETEUR", "VENDEUR");

    private final UserRepository users;

    public Fixed(UserRepository users) {
        this.users = users;
    }

    @PatchMapping("/api/profile/role")
    @ResponseBody
    public String updateRole(@RequestParam String role, Authentication auth) {
        if (!SELF_ASSIGNABLE_ROLES.contains(role)) {
            // Toute valeur hors liste blanche (dont "ADMIN") est rejetée,
            // quelle que soit l'identité de l'appelant.
            return "rôle non autorisé";
        }
        Long currentUserId = (Long) auth.getPrincipal();
        users.setRole(currentUserId, role);
        return "rôle mis à jour";
    }

    // Élévation vers un rôle privilégié : endpoint séparé, réservé aux
    // administrateurs existants, jamais accessible en self-service.
    @PatchMapping("/api/admin/users/{targetUserId}/role")
    @ResponseBody
    @org.springframework.security.access.prepost.PreAuthorize("hasRole('ADMIN')")
    public String promoteUser(@PathVariable Long targetUserId, @RequestParam String role) {
        users.setRole(targetUserId, role);
        return "rôle mis à jour par un administrateur";
    }

    interface UserRepository { void setRole(Long userId, String role); }
}
