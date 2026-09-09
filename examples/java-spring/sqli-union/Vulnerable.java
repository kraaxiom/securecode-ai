// Vulnérable : UNION-based SQL Injection (CWE-89)
// L'identifiant produit et le tri sont concaténés directement dans la requête.
// Un attaquant peut injecter une clause UNION SELECT pour extraire des données
// d'autres tables (ex: identifiants et mots de passe de la table users) via
// une colonne d'affichage normalement destinée aux produits.
package com.example.security.sqliunion;

import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.Map;

@RestController
public class SqliUnionController {

    private final JdbcTemplate jdbcTemplate;

    public SqliUnionController(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }

    @GetMapping("/api/products")
    public List<Map<String, Object>> listByCategory(@RequestParam String categoryId,
                                                      @RequestParam(defaultValue = "name") String sort) {
        // categoryId = "1 UNION SELECT username, password FROM users -- "
        // sort provient également de l'utilisateur sans liste blanche.
        String sql = "SELECT name, price FROM products WHERE category_id = " + categoryId
                + " ORDER BY " + sort;
        return jdbcTemplate.queryForList(sql);
    }
}
