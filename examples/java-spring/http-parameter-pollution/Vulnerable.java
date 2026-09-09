// Vulnérable : HTTP Parameter Pollution (CWE-235)
// Le paramètre "role" censé être unique n'est jamais vérifié quant à sa
// cardinalité. Si la requête contient ?role=user&role=admin, Spring MVC
// binde silencieusement une seule valeur (souvent la première), mais un
// proxy/WAF en amont peut interpréter la seconde occurrence, créant une
// divergence exploitable pour contourner un contrôle de sécurité.
package com.example.security.hpp;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HttpParameterPollutionController {

    @PostMapping("/api/account/role")
    public String assignRole(@RequestParam String role) {
        // Aucune vérification que le paramètre "role" n'a été fourni
        // qu'une seule fois : la duplication n'est jamais détectée.
        return assignerRole(role);
    }

    private String assignerRole(String role) {
        return "Rôle assigné : " + role;
    }
}
