package com.example.demo.security;

import java.security.SecureRandom;
import java.util.Base64;

// CORRIGÉ — l'identifiant de session est généré par un générateur
// aléatoire cryptographiquement sûr (CSPRNG), avec suffisamment d'entropie
// (≥ 128 bits) pour rendre toute prédiction ou énumération computationnellement
// infaisable. En pratique, préférer déléguer entièrement la gestion des
// sessions au conteneur/framework (Spring Session), qui applique déjà
// cette génération correctement.
public class Fixed {

    private final SecureRandom secureRandom = new SecureRandom();

    public String generateSessionId() {
        byte[] randomBytes = new byte[24]; // 192 bits d'entropie
        secureRandom.nextBytes(randomBytes);
        return Base64.getUrlEncoder().withoutPadding().encodeToString(randomBytes);
    }
}
