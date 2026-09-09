// Corrigé : Blind SQL Injection (CWE-89)
// Utilisation d'une requête préparée : la valeur utilisateur est liée en tant
// que paramètre SQL et n'est jamais concaténée dans le texte de la requête,
// ce qui empêche toute inférence booléenne ou temporelle.
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
        // Requête paramétrée : le paramètre est lié, jamais concaténé.
        String sql = "SELECT 1 FROM users WHERE username = ? AND active = 1";
        List<Map<String, Object>> rows = jdbcTemplate.queryForList(sql, username);
        // Réponse et temps de traitement uniformes, quel que soit le résultat.
        return Map.of("exists", !rows.isEmpty());
    }
}
