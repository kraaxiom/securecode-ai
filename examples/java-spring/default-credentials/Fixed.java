package com.example.demo.config;

import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;

import java.security.SecureRandom;
import java.util.Base64;

// CORRIGÉ — un mot de passe aléatoire est généré à la première
// initialisation, affiché UNE SEULE FOIS dans les logs de démarrage (à
// récupérer par l'opérateur), stocké haché, et un flag oblige son
// changement à la première connexion. Aucune valeur fixe n'est jamais
// codée en dur ni documentée publiquement.
public class Fixed {

    private final UserRepository users;
    private final BCryptPasswordEncoder encoder = new BCryptPasswordEncoder(12);
    private final SecureRandom random = new SecureRandom();

    public Fixed(UserRepository users) {
        this.users = users;
    }

    public void initAdminAccount() {
        if (users.findByUsername("admin") == null) {
            String generatedPassword = generateRandomPassword();
            users.create("admin", encoder.encode(generatedPassword), /* mustChangePassword */ true);
            // Affiché une seule fois au démarrage, à récupérer immédiatement
            // par l'opérateur (jamais persisté en clair, jamais dans un dépôt).
            System.out.println("Mot de passe admin initial (à changer immédiatement) : " + generatedPassword);
        }
    }

    private String generateRandomPassword() {
        byte[] bytes = new byte[24];
        random.nextBytes(bytes);
        return Base64.getUrlEncoder().withoutPadding().encodeToString(bytes);
    }

    interface UserRepository {
        Object findByUsername(String username);
        void create(String username, String hashedPassword, boolean mustChangePassword);
    }
}
