// Vulnérable : Command Injection (CWE-78)
// La commande système est construite par concaténation d'une entrée
// utilisateur puis exécutée via un shell, permettant à un attaquant
// d'injecter des métacaractères shell (`;`, `|`, `&&`) pour exécuter des
// commandes arbitraires.
package com.example.demo.controller;

import org.springframework.web.bind.annotation.*;

import java.io.BufferedReader;
import java.io.InputStreamReader;

@RestController
@RequestMapping("/api/network")
public class CommandInjectionController {

    @GetMapping("/ping")
    public String ping(@RequestParam String host) throws Exception {
        // Concaténation de l'entrée utilisateur puis passage par un shell.
        Process process = Runtime.getRuntime().exec(new String[] {
                "/bin/sh", "-c", "ping -c 3 " + host
        });
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
