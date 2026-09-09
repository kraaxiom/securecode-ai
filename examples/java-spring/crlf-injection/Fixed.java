// Corrigé : CRLF Injection (CWE-93)
// L'entrée est validée strictement (chemin relatif, sans caractères de
// contrôle) avant d'être utilisée, et l'API sûre du framework
// (ResponseEntity + HttpHeaders) est utilisée à la place d'une écriture
// manuelle d'en-tête.
package com.example.demo.controller;

import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.net.URI;
import java.util.regex.Pattern;

@RestController
@RequestMapping("/redirect")
public class CrlfInjectionController {

    // Un chemin relatif commençant par "/" et sans retour à la ligne ni retour chariot.
    private static final Pattern SAFE_PATH = Pattern.compile("^/[^\\r\\n]*$");

    @GetMapping
    public ResponseEntity<Void> redirectTo(@RequestParam(defaultValue = "/") String next) {
        // Toute entrée contenant \r ou \n, ou ne débutant pas par "/", est rejetée.
        String safeNext = SAFE_PATH.matcher(next).matches() ? next : "/";

        HttpHeaders headers = new HttpHeaders();
        headers.setLocation(URI.create(safeNext));
        return new ResponseEntity<>(headers, HttpStatus.FOUND);
    }
}
