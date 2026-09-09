// Corrigé : OS Command Injection (CWE-78)
// L'entrée est validée selon une liste blanche stricte (format d'hôte/IP), puis
// la commande est exécutée via ProcessBuilder avec un tableau d'arguments fixe,
// sans jamais passer par un interpréteur shell. Aucune interprétation de
// méta-caractères shell n'est possible.
package com.example.security.command;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.regex.Pattern;

@RestController
public class OsCommandInjectionController {

    // Liste blanche stricte : nom d'hôte ou adresse IPv4 uniquement.
    private static final Pattern ALLOWED_HOST = Pattern.compile("^[a-zA-Z0-9.-]{1,253}$");

    @GetMapping("/api/ping")
    public String ping(@RequestParam String host) throws Exception {
        if (!ALLOWED_HOST.matcher(host).matches()) {
            throw new IllegalArgumentException("Hôte invalide");
        }

        // ProcessBuilder avec tableau d'arguments : aucune interprétation shell,
        // donc aucun méta-caractère (;, |, &&, `) ne peut être exploité.
        ProcessBuilder builder = new ProcessBuilder("ping", "-c", "4", host);
        builder.redirectErrorStream(true);
        Process process = builder.start();

        StringBuilder output = new StringBuilder();
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()))) {
            String line;
            while ((line = reader.readLine()) != null) {
                output.append(line).append("\n");
            }
        }
        return output.toString();
    }
}
