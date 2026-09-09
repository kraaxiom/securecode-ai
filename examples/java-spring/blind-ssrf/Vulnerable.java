package com.example.demo.controller;

import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

// VULNÉRABLE — Blind SSRF (CWE-918 : Server-Side Request Forgery)
// La réponse de la requête sortante n'est jamais renvoyée au client,
// mais l'URL fournie est tout de même appelée côté serveur (webhook de
// notification). L'attaquant ne voit pas la réponse mais peut tout de
// même sonder le réseau interne via le délai de réponse ou des
// effets de bord (out-of-band).
@RestController
public class Vulnerable {

    private final RestTemplate restTemplate = new RestTemplate();

    @PostMapping("/webhooks/register")
    public void registerWebhook(@RequestParam String callbackUrl) {
        // Callback utilisateur appelé sans validation : SSRF aveugle.
        restTemplate.postForLocation(callbackUrl, "ping");
    }
}
