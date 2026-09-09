// Corrigé : Stacked Query SQL Injection (CWE-89)
// Requête préparée avec paramètre lié : la valeur utilisateur ne peut plus
// modifier la structure de la requête ni introduire d'instruction SQL
// supplémentaire. La chaîne de connexion JDBC désactive explicitement le
// support des instructions multiples (ex: "allowMultiQueries=false").
package com.example.security.stackedquery;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import javax.sql.DataSource;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;

@RestController
public class StackedQuerySqliController {

    private final DataSource dataSource;

    // La DataSource doit être configurée avec allowMultiQueries=false (MySQL) ou équivalent
    public StackedQuerySqliController(DataSource dataSource) {
        this.dataSource = dataSource;
    }

    @PostMapping("/users/rename")
    public String rename(@RequestParam String name) throws SQLException {
        try (Connection connection = dataSource.getConnection();
             PreparedStatement statement =
                     connection.prepareStatement("UPDATE users SET name = ? WHERE id = 1")) {
            // Paramètre lié : impossible d'injecter une instruction SQL distincte
            statement.setString(1, name);
            statement.executeUpdate();
        }
        return "Nom mis à jour";
    }
}
