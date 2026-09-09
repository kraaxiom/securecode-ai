// Vulnérable : NoSQL Injection (CWE-943)
// Le filtre de requête MongoDB est construit directement à partir du corps JSON
// brut envoyé par le client. Un attaquant peut remplacer une valeur scalaire
// attendue par un objet contenant un opérateur MongoDB (ex: {"$ne": null}),
// modifiant ainsi la logique de la requête (contournement d'authentification).
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

        // Le corps JSON brut est injecté tel quel dans le filtre : un attaquant
        // peut envoyer {"username": "admin", "password": {"$ne": null}}.
        Document filter = new Document();
        filter.put("username", body.get("username"));
        filter.put("password", body.get("password"));

        return users.find(filter).first() != null;
    }
}
