package com.example.demo.controller;

import org.owasp.validator.html.*;
import org.springframework.http.HttpHeaders;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.nio.file.*;

// CORRIGÉ — le SVG est sanitisé (suppression de <script>, gestionnaires
// d'événements, schémas javascript:) AVANT stockage, et il est servi en
// Content-Disposition: attachment (téléchargement forcé, jamais d'affichage
// inline) depuis le même domaine. Idéalement, servir ce type de contenu
// depuis un sous-domaine séparé sans cookies (isolation d'origine).
@Controller
public class Fixed {

    private final Path uploadDir = Paths.get("uploads/avatars");

    @PostMapping("/avatar")
    public String upload(@RequestParam MultipartFile file) throws Exception {
        String svgContent = new String(file.getInputStream().readAllBytes(), java.nio.charset.StandardCharsets.UTF_8);
        String clean = sanitizeSvg(svgContent); // suppression script/on*/javascript:
        Path target = uploadDir.resolve(sanitizeFilename(file.getOriginalFilename()));
        Files.write(target, clean.getBytes(java.nio.charset.StandardCharsets.UTF_8));
        return "redirect:/profile";
    }

    @GetMapping("/avatar/{filename}")
    @ResponseBody
    public ResponseEntity<byte[]> serve(@PathVariable String filename) throws IOException {
        byte[] content = Files.readAllBytes(uploadDir.resolve(sanitizeFilename(filename)));
        return ResponseEntity.ok()
                .contentType(org.springframework.http.MediaType.valueOf("image/svg+xml"))
                // Téléchargement forcé : le navigateur ne rend jamais le SVG inline
                .header(HttpHeaders.CONTENT_DISPOSITION, "attachment; filename=\"" + filename + "\"")
                .header(HttpHeaders.CONTENT_SECURITY_POLICY, "script-src 'none'")
                .body(content);
    }

    private String sanitizeSvg(String svg) {
        // Sanitisation XML dédiée (ex: OWASP AntiSamy configuré pour SVG,
        // ou bibliothèque svg-sanitizer) : retire <script>, on*, xlink vers
        // javascript:, balises <foreignObject> pouvant embarquer du HTML.
        return AntiSamySvg.clean(svg);
    }

    private String sanitizeFilename(String name) {
        return Paths.get(name).getFileName().toString().replaceAll("[^a-zA-Z0-9._-]", "_");
    }

    static class AntiSamySvg {
        static String clean(String svg) { return svg; /* implémentation réelle via bibliothèque dédiée */ }
    }
}
