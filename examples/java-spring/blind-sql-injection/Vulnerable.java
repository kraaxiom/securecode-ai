// Vulnérable : Blind SQL Injection (CWE-89)
// La requête SQL est construite par concaténation d'une entrée utilisateur.
// Le résultat n'est pas affiché directement (seulement true/false), mais un
// attaquant peut quand même déduire des informations via des conditions
// booléennes ou des délais de réponse injectés dans la chaîne.
package com.example.demo.controller;

import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/api/users")
public class UserCheckController {

    private final JdbcTemplate jdbcTemplate;

    public UserCheckController(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }

    @GetMapping("/exists")
    public Map<String, Boolean> exists(@RequestParam String username) {
        // Concaténation directe de l'entrée utilisateur dans le texte SQL.
        String sql = "SELECT 1 FROM users WHERE username = '" + username + "' AND active = 1";
        List<Map<String, Object>> rows = jdbcTemplate.queryForList(sql);
        return Map.of("exists", !rows.isEmpty());
    }
}
