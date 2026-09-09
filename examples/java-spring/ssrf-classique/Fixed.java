package com.example.demo.controller;

import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

import java.net.InetAddress;
import java.net.URI;
import java.util.List;
import java.util.Set;

// CORRIGÉ — SSRF classique (CWE-918)
// La destination est validée contre une liste blanche de schémas et
// d'hôtes autorisés, puis l'IP résolue est vérifiée pour exclure les
// plages privées/loopback/link-local avant tout appel sortant.
@RestController
public class Fixed {

    private static final Set<String> ALLOWED_HOSTS = Set.of("images.partner-cdn.example.com");
    private final RestTemplate restTemplate = new RestTemplate();

    @GetMapping("/preview")
    public String previewUrl(@RequestParam String url) throws Exception {
        URI uri = URI.create(url);

        if (!"https".equalsIgnoreCase(uri.getScheme())) {
            throw new IllegalArgumentException("Schéma non autorisé");
        }
        if (!ALLOWED_HOSTS.contains(uri.getHost())) {
            throw new IllegalArgumentException("Hôte non autorisé");
        }

        // Résolution DNS et contrôle de l'IP réelle pour empêcher tout
        // détournement vers une plage privée (défense en profondeur,
        // complémentaire à la whitelist d'hôtes).
        for (InetAddress addr : InetAddress.getAllByName(uri.getHost())) {
            if (addr.isLoopbackAddress() || addr.isLinkLocalAddress() || addr.isSiteLocalAddress()) {
                throw new IllegalArgumentException("Adresse IP interne refusée");
            }
        }

        return restTemplate.getForObject(uri, String.class);
    }
}
