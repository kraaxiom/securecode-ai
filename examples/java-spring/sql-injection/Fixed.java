// Corrigé : SQL Injection (CWE-89)
// La requête utilise un paramètre lié via JdbcTemplate (placeholder "?") : la
// valeur utilisateur est transmise séparément du texte SQL et n'est jamais
// interprétée comme du code SQL, ce qui empêche toute injection.
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
        // Requête paramétrée : le placeholder "?" garantit que la valeur est
        // liée en tant que donnée, jamais concaténée dans le texte SQL.
        String sql = "SELECT id, username, email FROM users WHERE name = ?";
        return jdbcTemplate.queryForList(sql, name);
    }
}
