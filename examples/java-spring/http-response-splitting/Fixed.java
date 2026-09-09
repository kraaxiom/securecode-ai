// Corrigé : HTTP Response Splitting (CWE-113)
// La destination de redirection est restreinte à une liste blanche de
// chemins internes connus, et l'API haut niveau sendRedirect() du conteneur
// Servlet est utilisée plutôt qu'une écriture brute de l'en-tête Location.
// Toute valeur hors liste blanche (y compris contenant CR/LF ou pointant
// vers un hôte externe) est remplacée par une destination par défaut sûre.
package com.example.security.responsesplitting;

import jakarta.servlet.http.HttpServletResponse;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.io.IOException;
import java.util.Set;

@RestController
public class HttpResponseSplittingController {

    private static final Set<String> ALLOWED_PATHS = Set.of("/dashboard", "/profile", "/account");
    private static final String DEFAULT_PATH = "/dashboard";

    @GetMapping("/redirect")
    public void redirect(@RequestParam(required = false) String next, HttpServletResponse response) throws IOException {
        // Liste blanche stricte : toute valeur absente de cet ensemble
        // (y compris une tentative d'injection CR/LF ou une redirection
        // ouverte vers un hôte externe) retombe sur la destination par défaut.
        String target = (next != null && ALLOWED_PATHS.contains(next)) ? next : DEFAULT_PATH;

        // sendRedirect() utilise l'API du conteneur Servlet, qui encode
        // correctement l'en-tête Location plutôt qu'une écriture brute.
        response.sendRedirect(target);
    }
}
