package com.example.demo.security;

import io.jsonwebtoken.*;

import java.security.PublicKey;

// CORRIGÉ — l'algorithme de signature attendu est imposé explicitement
// côté serveur (RS256 uniquement), indépendamment de ce que déclare le
// header du token. Tout token signé avec un autre algorithme (dont HS256
// fabriqué à partir de la clé publique) est rejeté avant même la
// vérification cryptographique de la signature.
public class Fixed {

    private final PublicKey publicKey;

    public Fixed(PublicKey publicKey) {
        this.publicKey = publicKey;
    }

    public Claims verify(String token) {
        Jws<Claims> jws = Jwts.parserBuilder()
                .setSigningKey(publicKey)
                .build()
                .parseClaimsJws(token);

        // Vérification explicite et redondante de l'algorithme déclaré :
        // défense en profondeur en plus de la configuration du parser.
        String alg = jws.getHeader().getAlgorithm();
        if (!"RS256".equals(alg)) {
            throw new JwtException("Algorithme non autorisé : " + alg);
        }

        return jws.getBody();
    }
}
