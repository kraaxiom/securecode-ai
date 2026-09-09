// Corrigé : Command Injection (CWE-78)
// L'entrée est validée strictement comme adresse IP, puis le processus est
// lancé avec un tableau d'arguments distincts via ProcessBuilder, sans
// aucun passage par un interpréteur shell.
package com.example.demo.controller;

import org.springframework.web.bind.annotation.*;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.InetAddresses;

@RestController
@RequestMapping("/api/network")
public class CommandInjectionController {

    @GetMapping("/ping")
    public String ping(@RequestParam String host) throws Exception {
        // Validation stricte : seule une adresse IP littérale est acceptée.
        if (!InetAddresses.isInetAddress(host)) {
            throw new IllegalArgumentException("Hôte invalide");
        }

        // Arguments passés en tableau distinct, sans shell.
        ProcessBuilder builder = new ProcessBuilder("ping", "-c", "3", host);
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
