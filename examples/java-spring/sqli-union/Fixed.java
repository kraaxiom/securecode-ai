// Corrigé : UNION-based SQL Injection (CWE-89)
// L'identifiant de catégorie est casté en entier et lié via un paramètre
// PreparedStatement (JdbcTemplate), empêchant toute clause UNION injectée.
// Le champ de tri, qui ne peut pas être paramétré en SQL, passe par une liste
// blanche stricte d'identifiants de colonnes autorisés.
package com.example.security.sqliunion;

import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.Map;
import java.util.Set;

@RestController
public class SqliUnionController {

    private final JdbcTemplate jdbcTemplate;

    // Liste blanche stricte des colonnes de tri autorisées.
    private static final Set<String> ALLOWED_SORT_COLUMNS = Set.of("name", "price");

    public SqliUnionController(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }

    @GetMapping("/api/products")
    public List<Map<String, Object>> listByCategory(@RequestParam String categoryId,
                                                      @RequestParam(defaultValue = "name") String sort) {
        // Typage strict : un identifiant non numérique lève une exception
        // avant d'atteindre la requête (empêche toute clause UNION).
        int categoryIdInt;
        try {
            categoryIdInt = Integer.parseInt(categoryId);
        } catch (NumberFormatException e) {
            throw new IllegalArgumentException("Identifiant de catégorie invalide");
        }

        // Le tri ne peut pas être paramétré en SQL standard : liste blanche obligatoire.
        String safeSort = ALLOWED_SORT_COLUMNS.contains(sort) ? sort : "name";

        String sql = "SELECT name, price FROM products WHERE category_id = ? ORDER BY " + safeSort;
        return jdbcTemplate.queryForList(sql, categoryIdInt);
    }
}
