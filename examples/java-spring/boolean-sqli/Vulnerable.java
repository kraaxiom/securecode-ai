// Vulnérable : Boolean-based SQL Injection (CWE-89)
// La clause WHERE est construite par concaténation d'une entrée utilisateur,
// permettant à un attaquant d'altérer la logique booléenne de la requête
// (ex: injection d'un `OR 1=1`) pour extraire ou modifier le jeu de résultats.
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
        // Concaténation directe de l'entrée utilisateur dans la clause WHERE.
        String sql = "SELECT * FROM products WHERE name = '" + name + "'";
        return jdbcTemplate.queryForList(sql);
    }
}
