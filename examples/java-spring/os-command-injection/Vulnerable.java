// Vulnérable : OS Command Injection (CWE-78)
// La commande shell est construite par concaténation d'une entrée utilisateur
// puis exécutée via un interpréteur shell ("sh -c"). Un attaquant peut injecter
// des méta-caractères shell (;, |, &&, `) pour exécuter des commandes arbitraires
// sur le système.
package com.example.security.command;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.io.BufferedReader;
import java.io.InputStreamReader;

@RestController
public class OsCommandInjectionController {

    @GetMapping("/api/ping")
    public String ping(@RequestParam String host) throws Exception {
        // host = "8.8.8.8; cat /etc/passwd" exécute une seconde commande.
        String command = "ping -c 4 " + host;

        Process process = Runtime.getRuntime().exec(new String[]{"sh", "-c", command});

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
