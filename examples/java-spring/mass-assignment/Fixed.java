package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// CORRIGÉ — un DTO d'entrée dédié n'expose que les champs réellement
// modifiables par l'utilisateur (liste blanche explicite). Les champs
// sensibles ("role", "isAdmin") n'existent tout simplement pas dans ce
// DTO : ils ne peuvent structurellement pas être désérialisés depuis le
// corps de la requête, quelle que soit sa valeur.
@Controller
public class Fixed {

    private final UserRepository users;

    public Fixed(UserRepository users) {
        this.users = users;
    }

    @PutMapping("/api/users/{id}")
    @ResponseBody
    public User updateUser(@PathVariable Long id, @RequestBody UpdateUserRequest request) {
        User existing = users.findById(id);
        existing.name = request.name;   // affectation champ par champ, explicite
        existing.email = request.email; // "role"/"isAdmin" restent inchangés, jamais exposés ici
        return users.save(id, existing);
    }

    interface UserRepository {
        User findById(Long id);
        User save(Long id, User user);
    }

    // DTO d'entrée : liste blanche stricte, aucun champ sensible présent
    static class UpdateUserRequest {
        String name;
        String email;
    }

    static class User {
        Long id;
        String name;
        String email;
        String role;
        boolean isAdmin;
    }
}
