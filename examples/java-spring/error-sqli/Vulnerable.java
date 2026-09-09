// Vulnérable : Error-based SQL Injection (CWE-89)
// La requête SQL est construite par concaténation, et le message d'erreur
// détaillé du driver SQL (incluant fragments de requête ou de schéma) est
// renvoyé directement au client, permettant l'extraction d'informations
// via des erreurs provoquées volontairement.
package com.example.demo.controller;

import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/api/orders")
public class ErrorSqliController {

    private final JdbcTemplate jdbcTemplate;

    public ErrorSqliController(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }

    @GetMapping("/{id}")
    public Object getOrder(@PathVariable String id) {
        try {
            // Concaténation directe de l'entrée utilisateur dans la requête.
            String sql = "SELECT * FROM orders WHERE id = " + id;
            List<Map<String, Object>> rows = jdbcTemplate.queryForList(sql);
            return rows;
        } catch (Exception e) {
            // Fuite du message d'erreur SQL détaillé vers le client.
            return Map.of("error", "Erreur SQL: " + e.getMessage());
        }
    }
}
