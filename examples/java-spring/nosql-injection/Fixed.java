// Corrigé : NoSQL Injection (CWE-943)
// Chaque champ issu du client est validé par type avant d'être inséré dans le
// filtre MongoDB : tout objet ou opérateur (`$ne`, `$gt`, `$where`, `$regex`)
// envoyé à la place d'une chaîne scalaire est rejeté avant d'atteindre la base.
package com.example.security.nosql;

import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoDatabase;
import org.bson.Document;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.Map;

@RestController
public class NoSqlInjectionController {

    private final MongoDatabase database;

    public NoSqlInjectionController(MongoDatabase database) {
        this.database = database;
    }

    @PostMapping("/api/login")
    public boolean login(@RequestBody Map<String, Object> body) {
        MongoCollection<Document> users = database.getCollection("users");

        Object rawUsername = body.get("username");
        Object rawPassword = body.get("password");

        // Rejet strict : seule une valeur de type String est acceptée.
        // Un objet (ex: opérateur MongoDB) est refusé avant construction du filtre.
        if (!(rawUsername instanceof String) || !(rawPassword instanceof String)) {
            throw new IllegalArgumentException("Format d'identifiants invalide");
        }

        String username = (String) rawUsername;
        String password = (String) rawPassword;

        // Construction typée du filtre : les valeurs sont garanties scalaires.
        Document filter = new Document("username", username).append("password", password);

        return users.find(filter).first() != null;
    }
}
