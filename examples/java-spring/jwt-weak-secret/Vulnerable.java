package com.example.demo.security;

import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;

import java.nio.charset.StandardCharsets;

// VULNÉRABLE — JWT Weak Secret (CWE-326 : Inadequate Encryption Strength)
// Le secret HMAC utilisé pour signer les tokens est un mot court, facile à
// deviner et à retrouver par recherche exhaustive hors-ligne (brute force
// sur "secret123" avec des outils comme hashcat sur des tokens HS256 en
// quelques secondes/minutes).
public class Vulnerable {

    private static final String SECRET = "secret123"; // faible, court, devinable

    public String issueToken(String userId) {
        return Jwts.builder()
                .setSubject(userId)
                .signWith(SignatureAlgorithm.HS256, SECRET.getBytes(StandardCharsets.UTF_8))
                .compact();
    }
}
