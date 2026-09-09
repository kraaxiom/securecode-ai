package com.example.demo.controller;

import org.springframework.http.MediaType;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.nio.file.*;

// VULNÉRABLE — XSS via SVG (CWE-79)
// Le fichier SVG uploadé est stocké et servi tel quel, avec un
// Content-Type image/svg+xml. Or un SVG peut contenir <script>,
// des gestionnaires d'événements (onload) ou des liens javascript:,
// exécutés par le navigateur lorsque le fichier est ouvert/affiché
// dans le même contexte d'origine que l'application.
@Controller
public class Vulnerable {

    private final Path uploadDir = Paths.get("uploads/avatars");

    @PostMapping("/avatar")
    public String upload(@RequestParam MultipartFile file) throws IOException {
        Path target = uploadDir.resolve(file.getOriginalFilename());
        Files.copy(file.getInputStream(), target, StandardCopyOption.REPLACE_EXISTING);
        return "redirect:/profile";
    }

    @GetMapping(value = "/avatar/{filename}", produces = "image/svg+xml")
    @ResponseBody
    public byte[] serve(@PathVariable String filename) throws IOException {
        // Servi en affichage inline, sur le même domaine que l'application,
        // avec les cookies de session actifs : un <script> dans le SVG
        // s'exécute dans le contexte de l'application.
        return Files.readAllBytes(uploadDir.resolve(filename));
    }
}
