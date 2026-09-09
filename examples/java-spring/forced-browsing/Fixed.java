package com.example.demo.controller;

import org.springframework.core.io.Resource;
import org.springframework.core.io.UrlResource;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.nio.file.Path;
import java.nio.file.Paths;

// CORRIGÉ — chaque accès exige une authentification ET une vérification
// explicite que le document demandé est bien autorisé pour l'utilisateur
// courant (liste d'accès en base, pas simple présence du fichier sur
// disque). Le nom de fichier est également validé pour empêcher toute
// traversée de répertoire en complément.
@Controller
public class Fixed {

    private final DocumentAccessRepository accessRepository;
    private final Path baseDir = Paths.get("internal-docs").toAbsolutePath().normalize();

    public Fixed(DocumentAccessRepository accessRepository) {
        this.accessRepository = accessRepository;
    }

    @GetMapping("/documents/{filename}")
    @ResponseBody
    @PreAuthorize("isAuthenticated()")
    public ResponseEntity<Resource> getDocument(@PathVariable String filename, org.springframework.security.core.Authentication auth) throws Exception {
        Long currentUserId = (Long) auth.getPrincipal();

        // Vérification explicite en base que ce document est autorisé pour
        // cet utilisateur, indépendamment de la connaissance de son nom.
        if (!accessRepository.isAuthorized(currentUserId, filename)) {
            return ResponseEntity.status(403).build();
        }

        Path target = baseDir.resolve(filename).normalize();
        if (!target.startsWith(baseDir)) {
            return ResponseEntity.badRequest().build(); // anti path traversal
        }

        return ResponseEntity.ok(new UrlResource(target.toUri()));
    }

    interface DocumentAccessRepository { boolean isAuthorized(Long userId, String filename); }
}
