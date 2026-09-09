// Corrigé : HTTP Parameter Pollution (CWE-235)
// La requête est lue via HttpServletRequest.getParameterValues() afin de
// détecter explicitement toute duplication du paramètre "role". Une
// requête contenant plusieurs occurrences est rejetée (comportement
// déterministe et documenté), éliminant toute ambiguïté d'interprétation
// entre les couches de la chaîne (proxy, WAF, backend).
package com.example.security.hpp;

import jakarta.servlet.http.HttpServletRequest;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ResponseStatusException;

@RestController
public class HttpParameterPollutionController {

    @PostMapping("/api/account/role")
    public String assignRole(HttpServletRequest request) {
        // Lecture explicite de toutes les occurrences du paramètre afin de
        // détecter une pollution avant tout traitement métier.
        String[] values = request.getParameterValues("role");

        if (values == null || values.length == 0) {
            throw new ResponseStatusException(HttpStatus.BAD_REQUEST, "Paramètre 'role' manquant");
        }
        // Politique documentée : un paramètre dupliqué est rejeté plutôt
        // que de choisir silencieusement une valeur arbitraire.
        if (values.length > 1) {
            throw new ResponseStatusException(HttpStatus.BAD_REQUEST, "Paramètre dupliqué non autorisé");
        }

        return assignerRole(values[0]);
    }

    private String assignerRole(String role) {
        return "Rôle assigné : " + role;
    }
}
