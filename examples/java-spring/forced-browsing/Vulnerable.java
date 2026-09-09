package com.example.demo.controller;

import org.springframework.core.io.Resource;
import org.springframework.core.io.UrlResource;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.nio.file.Paths;

// VULNÉRABLE — Forced Browsing (CWE-425 : Direct Request / Forced
// Browsing)
// Les documents internes sont accessibles via une URL prévisible
// (/documents/{filename}), sans authentification ni autorisation, en
// misant uniquement sur le fait que l'URL exacte n'est "pas censée" être
// connue (sécurité par l'obscurité). Un attaquant qui devine ou énumère
// les noms de fichiers accède directement au contenu.
@Controller
public class Vulnerable {

    @GetMapping("/documents/{filename}")
    @ResponseBody
    public Resource getDocument(@PathVariable String filename) throws Exception {
        // Aucune vérification d'authentification, aucun contrôle
        // d'autorisation : la seule "protection" est de ne pas publier
        // les URLs, qui restent pourtant devinables (ex: rapport-2024.pdf,
        // rapport-2025.pdf, contrat-client-42.pdf).
        return new UrlResource(Paths.get("internal-docs/" + filename).toUri());
    }
}
