// Vulnérable : Stacked Query SQL Injection (CWE-89)
// Le nom fourni par l'utilisateur est concaténé dans une chaîne SQL exécutée
// via Statement, qui autorise l'exécution de plusieurs instructions séparées
// par des points-virgules. Un attaquant peut ainsi ajouter une instruction
// SQL arbitraire (ex: "; DROP TABLE users; --").
package com.example.security.stackedquery;

import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import javax.sql.DataSource;
import java.sql.Connection;
import java.sql.SQLException;
import java.sql.Statement;

@RestController
public class StackedQuerySqliController {

    private final DataSource dataSource;

    public StackedQuerySqliController(DataSource dataSource) {
        this.dataSource = dataSource;
    }

    @PostMapping("/users/rename")
    public String rename(@RequestParam String name) throws SQLException {
        try (Connection connection = dataSource.getConnection();
             Statement statement = connection.createStatement()) {
            // Concaténation directe : le pilote JDBC utilisé autorise le multi-instructions
            statement.execute("UPDATE users SET name = '" + name + "' WHERE id = 1");
        }
        return "Nom mis à jour";
    }
}
