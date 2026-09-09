// Corrigé : Boolean-based SQL Injection (CWE-89)
// La comparaison est réalisée via une requête préparée avec paramètre lié :
// la valeur utilisateur ne peut plus altérer la structure logique de la
// requête, quel que soit son contenu.
package com.example.demo.controller;

import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/api/products")
public class BooleanSqliController {

    private final JdbcTemplate jdbcTemplate;

    public BooleanSqliController(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }

    @GetMapping("/search")
    public List<Map<String, Object>> search(@RequestParam String name) {
        // Requête paramétrée : impossible d'injecter une expression logique SQL.
        String sql = "SELECT * FROM products WHERE name = ?";
        return jdbcTemplate.queryForList(sql, name);
    }
}
