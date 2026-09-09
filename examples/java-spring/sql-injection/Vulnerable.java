// Vulnérable : SQL Injection (CWE-89)
// La requête SQL est construite par concaténation directe du paramètre de
// recherche utilisateur. Un attaquant peut injecter des clauses SQL arbitraires
// (ex: name = "' OR '1'='1") pour altérer la logique de la requête et accéder
// à des données non autorisées.
package com.example.security.sqli;

import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.Map;

@RestController
public class SqlInjectionController {

    private final JdbcTemplate jdbcTemplate;

    public SqlInjectionController(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }

    @GetMapping("/api/users/search")
    public List<Map<String, Object>> search(@RequestParam String name) {
        // Concaténation directe : name = "' OR '1'='1" retourne tous les utilisateurs.
        String sql = "SELECT id, username, email FROM users WHERE name = '" + name + "'";
        return jdbcTemplate.queryForList(sql);
    }
}
