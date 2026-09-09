// Corrigé : Error-based SQL Injection (CWE-89)
// Utilisation d'une requête préparée avec l'identifiant validé/casté en
// entier, et masquage du détail de l'exception SQL au client : seul un
// message générique est renvoyé, l'erreur complète est journalisée
// côté serveur.
package com.example.demo.controller;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/api/orders")
public class ErrorSqliController {

    private static final Logger log = LoggerFactory.getLogger(ErrorSqliController.class);

    private final JdbcTemplate jdbcTemplate;

    public ErrorSqliController(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }

    @GetMapping("/{id}")
    public ResponseEntity<?> getOrder(@PathVariable String id) {
        final long orderId;
        try {
            // Validation stricte du type attendu avant toute utilisation en requête.
            orderId = Long.parseLong(id);
        } catch (NumberFormatException e) {
            return ResponseEntity.badRequest().body(Map.of("error", "Identifiant invalide"));
        }

        try {
            // Requête paramétrée : aucune concaténation de variable.
            String sql = "SELECT * FROM orders WHERE id = ?";
            List<Map<String, Object>> rows = jdbcTemplate.queryForList(sql, orderId);
            return ResponseEntity.ok(rows);
        } catch (Exception e) {
            // Détail journalisé côté serveur uniquement, jamais renvoyé au client.
            log.error("Erreur lors de la récupération de la commande {}", orderId, e);
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR)
                    .body(Map.of("error", "Une erreur est survenue."));
        }
    }
}
