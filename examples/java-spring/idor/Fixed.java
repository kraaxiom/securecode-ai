package com.example.demo.controller;

import org.springframework.core.io.Resource;
import org.springframework.core.io.UrlResource;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.Authentication;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.nio.file.Paths;

// CORRIGÉ — l'appartenance de la facture à l'utilisateur courant est
// vérifiée explicitement AVANT toute résolution du chemin de fichier.
// En complément, on peut remplacer les identifiants séquentiels
// prévisibles par des UUID pour réduire la surface d'énumération, mais
// cela ne dispense jamais du contrôle d'autorisation lui-même.
@Controller
public class Fixed {

    private final InvoiceRepository invoices;

    public Fixed(InvoiceRepository invoices) {
        this.invoices = invoices;
    }

    @GetMapping("/invoices/{invoiceId}/download")
    @ResponseBody
    public ResponseEntity<Resource> downloadInvoice(@PathVariable Long invoiceId, Authentication auth) throws Exception {
        Long currentUserId = (Long) auth.getPrincipal();

        if (!invoices.belongsToUser(invoiceId, currentUserId)) {
            // Réponse identique (404) que la facture existe pour un autre
            // utilisateur ou n'existe pas du tout, pour ne pas révéler
            // l'existence d'IDs valides appartenant à des tiers.
            return ResponseEntity.notFound().build();
        }

        String path = invoices.getFilePath(invoiceId);
        return ResponseEntity.ok(new UrlResource(Paths.get(path).toUri()));
    }

    interface InvoiceRepository {
        boolean belongsToUser(Long invoiceId, Long userId);
        String getFilePath(Long invoiceId);
    }
}
