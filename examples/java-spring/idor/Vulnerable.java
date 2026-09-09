package com.example.demo.controller;

import org.springframework.core.io.Resource;
import org.springframework.core.io.UrlResource;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.nio.file.Paths;

// VULNÉRABLE — IDOR (Insecure Direct Object Reference) (CWE-639 :
// Authorization Bypass Through User-Controlled Key)
// La facture est récupérée directement à partir de son identifiant
// numérique séquentiel fourni par le client, sans vérifier que
// l'utilisateur connecté en est bien le destinataire. Incrémenter l'ID
// dans l'URL suffit à consulter les factures des autres clients.
@Controller
public class Vulnerable {

    private final InvoiceRepository invoices;

    public Vulnerable(InvoiceRepository invoices) {
        this.invoices = invoices;
    }

    @GetMapping("/invoices/{invoiceId}/download")
    @ResponseBody
    public Resource downloadInvoice(@PathVariable Long invoiceId) throws Exception {
        // Référence directe non vérifiée : /invoices/1001, /invoices/1002...
        String path = invoices.getFilePath(invoiceId);
        return new UrlResource(Paths.get(path).toUri());
    }

    interface InvoiceRepository { String getFilePath(Long invoiceId); }
}
