package com.example.demo.security;

import com.fasterxml.jackson.databind.ObjectMapper;

import java.util.Base64;
import java.util.Map;

// VULNÉRABLE — JWT "alg: none" (CWE-347 : Improper Verification of
// Cryptographic Signature)
// Le décodage personnalisé lit le header pour connaître l'algorithme, puis
// ne vérifie AUCUNE signature lorsque l'algorithme déclaré est "none".
// Un attaquant peut forger un token avec {"alg":"none"} et un payload
// arbitraire (ex: role=admin), sans posséder ni clé secrète ni clé privée.
public class Vulnerable {

    private final ObjectMapper mapper = new ObjectMapper();

    public Map<String, Object> decode(String token) throws Exception {
        String[] parts = token.split("\\.");
        Map<String, Object> header = mapper.readValue(Base64.getUrlDecoder().decode(parts[0]), Map.class);
        Map<String, Object> payload = mapper.readValue(Base64.getUrlDecoder().decode(parts[1]), Map.class);

        String alg = (String) header.get("alg");
        if ("none".equalsIgnoreCase(alg)) {
            // Aucune vérification de signature effectuée : le payload est
            // accepté tel quel, l'attaquant contrôle entièrement les claims.
            return payload;
        }
        // ... vérification de signature pour les autres algorithmes (omise ici)
        return payload;
    }
}
