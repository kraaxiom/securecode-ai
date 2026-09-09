package com.example.demo.controller;

import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

// VULNÉRABLE — SSRF classique (CWE-918 : Server-Side Request Forgery)
// L'URL fournie par le client est utilisée telle quelle pour effectuer
// une requête HTTP côté serveur, sans validation ni restriction de
// destination. Un attaquant peut cibler des services internes
// (http://127.0.0.1:8080/admin, http://169.254.169.254/...).
@RestController
public class Vulnerable {

    private final RestTemplate restTemplate = new RestTemplate();

    @GetMapping("/preview")
    public String previewUrl(@RequestParam String url) {
        // Aucune vérification de l'hôte/IP cible : requête forgée possible.
        return restTemplate.getForObject(url, String.class);
    }
}
