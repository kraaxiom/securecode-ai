package com.example.demo.security;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;

import java.security.PublicKey;

// VULNÉRABLE — JWT Algorithm Confusion (CWE-347 : Improper Verification
// of Cryptographic Signature)
// La vérification ne restreint pas l'algorithme attendu : elle se fie au
// champ "alg" présent dans le header du token, fourni par le client.
// Un attaquant en possession de la clé publique RS256 peut alors forger un
// token signé en HS256 en utilisant cette clé publique comme secret HMAC :
// le serveur, qui accepte n'importe quel algorithme, valide la signature
// avec succès.
public class Vulnerable {

    private final PublicKey publicKey;

    public Vulnerable(PublicKey publicKey) {
        this.publicKey = publicKey;
    }

    public Claims verify(String token) {
        // setSigningKey accepte la même clé quel que soit l'algorithme
        // déclaré dans le header, sans jamais imposer RS256 explicitement.
        return Jwts.parserBuilder()
                .setSigningKey(publicKey)
                .build()
                .parseClaimsJws(token)
                .getBody();
    }
}
