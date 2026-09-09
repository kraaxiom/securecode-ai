package com.example.demo.security;

import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.security.Keys;

import javax.crypto.SecretKey;

// CORRIGÉ — le secret est une clé aléatoire de longueur suffisante
// (≥ 256 bits pour HS256, générée via un CSPRNG), chargée depuis un
// gestionnaire de secrets (variable d'environnement / vault) et jamais
// codée en dur. Une taille de clé insuffisante rend l'attaque par
// recherche exhaustive hors-ligne triviale même avec un secret "aléatoire"
// trop court.
public class Fixed {

    // Clé chargée depuis la configuration externe (ex: Vault, variables
    // d'environnement), jamais codée en dur dans le code source.
    private final SecretKey signingKey;

    public Fixed(String base64EncodedSecretFromConfig) {
        // Keys.hmacShaKeyFor exige une clé d'au moins 256 bits pour HS256 ;
        // génération recommandée : Keys.secretKeyFor(SignatureAlgorithm.HS256)
        this.signingKey = Keys.hmacShaKeyFor(java.util.Base64.getDecoder().decode(base64EncodedSecretFromConfig));
    }

    public String issueToken(String userId) {
        return Jwts.builder()
                .setSubject(userId)
                .signWith(signingKey)
                .compact();
    }
}
