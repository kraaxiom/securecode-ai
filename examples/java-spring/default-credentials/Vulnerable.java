package com.example.demo.config;

// VULNÉRABLE — Default Credentials (CWE-1392 : Use of Default Credentials)
// Le compte administrateur est créé automatiquement au démarrage avec un
// identifiant et un mot de passe fixes, codés en dur, et rien n'oblige
// l'opérateur à les changer avant la mise en production. Ces identifiants
// par défaut sont publiquement documentés (README, image Docker), donc
// triviaux à deviner pour un attaquant qui scanne l'application.
public class Vulnerable {

    private final UserRepository users;

    public Vulnerable(UserRepository users) {
        this.users = users;
    }

    // Exécuté au démarrage de l'application (ex: CommandLineRunner)
    public void initAdminAccount() {
        if (users.findByUsername("admin") == null) {
            users.create("admin", "admin123"); // identifiants par défaut, jamais forcés au changement
        }
    }

    interface UserRepository {
        Object findByUsername(String username);
        void create(String username, String password);
    }
}
