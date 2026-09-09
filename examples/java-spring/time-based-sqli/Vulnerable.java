// Vulnérable : Time-Based Blind SQL Injection (CWE-89)
// L'identifiant est concaténé dans la requête SQL. Aucune donnée n'est
// renvoyée directement au client, mais un attaquant peut injecter une
// fonction de pause (ex: "SLEEP(5)") et déduire des informations en
// mesurant le délai de réponse.
package com.example.security.timebasedsqli;

import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class TimeBasedSqliController {

    private final JdbcTemplate jdbcTemplate;

    public TimeBasedSqliController(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }

    @GetMapping("/orders/exists")
    public String exists(@RequestParam String id) {
        // Concaténation directe : aucun timeout, aucun paramètre lié
        String sql = "SELECT 1 FROM orders WHERE id = " + id;
        var rows = jdbcTemplate.queryForList(sql);
        return rows.isEmpty() ? "introuvable" : "trouvé";
    }
}
