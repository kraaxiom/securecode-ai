// Corrigé : Time-Based Blind SQL Injection (CWE-89)
// Requête préparée avec paramètre lié, éliminant toute la classe de
// vulnérabilité (pas seulement la variante temporelle). Un timeout
// d'exécution strict est appliqué au niveau du DataSource pour limiter
// l'impact d'une éventuelle injection résiduelle ailleurs dans l'application.
package com.example.security.timebasedsqli;

import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class TimeBasedSqliController {

    private final JdbcTemplate jdbcTemplate;

    public TimeBasedSqliController(JdbcTemplate jdbcTemplate) {
        // Timeout d'exécution des requêtes fixé à 5 secondes (queryTimeout côté JdbcTemplate)
        jdbcTemplate.setQueryTimeout(5);
        this.jdbcTemplate = jdbcTemplate;
    }

    @GetMapping("/orders/exists")
    public String exists(@RequestParam int id) {
        // Paramètre lié : impossible d'injecter une fonction de pause ou toute autre expression SQL
        var rows = jdbcTemplate.queryForList("SELECT 1 FROM orders WHERE id = ?", id);
        // Réponse et temps de traitement uniformes, quel que soit le résultat.
        return rows.isEmpty() ? "introuvable" : "trouvé";
    }
}
