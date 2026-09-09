package com.example.demo.controller;

import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

import java.net.InetAddress;
import java.net.URI;

// CORRIGÉ — Blind SSRF (CWE-918)
// Même sans exposer la réponse au client, toute URL fournie par un
// utilisateur doit être validée avant un appel serveur. On impose
// HTTPS, on résout l'hôte et on rejette les plages d'IP internes.
@RestController
public class Fixed {

    private final RestTemplate restTemplate = new RestTemplate();

    @PostMapping("/webhooks/register")
    public void registerWebhook(@RequestParam String callbackUrl) throws Exception {
        URI uri = URI.create(callbackUrl);

        if (!"https".equalsIgnoreCase(uri.getScheme())) {
            throw new IllegalArgumentException("Seul HTTPS est autorisé pour les webhooks");
        }

        for (InetAddress addr : InetAddress.getAllByName(uri.getHost())) {
            if (addr.isLoopbackAddress() || addr.isAnyLocalAddress()
                    || addr.isLinkLocalAddress() || addr.isSiteLocalAddress()) {
                throw new IllegalArgumentException("Destination interne refusée");
            }
        }

        // Timeout court + pas de suivi de redirection : limite l'usage du
        // endpoint comme sonde réseau interne (SSRF aveugle par délai).
        restTemplate.postForLocation(uri, "ping");
    }
}
