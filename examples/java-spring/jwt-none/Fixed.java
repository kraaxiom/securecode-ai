package com.example.demo.security;

import io.jsonwebtoken.*;

import javax.crypto.SecretKey;

// CORRIGÉ — utilisation d'une bibliothèque JWT standard (jjwt) configurée
// avec une liste fermée d'algorithmes acceptés, "none" n'étant jamais
// autorisé. Toute tentative de token avec alg=none est rejetée par la
// bibliothèque elle-même, avant tout accès au payload.
public class Fixed {

    private final SecretKey signingKey;

    public Fixed(SecretKey signingKey) {
        this.signingKey = signingKey;
    }

    public Claims decode(String token) {
        // parseClaimsJws exige une signature valide vérifiable avec la clé
        // fournie ; un token "alg: none" (non signé) est automatiquement
        // rejeté avec une UnsupportedJwtException, sans traitement manuel.
        Jws<Claims> jws = Jwts.parserBuilder()
                .setSigningKey(signingKey)
                .build()
                .parseClaimsJws(token);
        return jws.getBody();
    }
}
