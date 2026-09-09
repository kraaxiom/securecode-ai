package com.example.demo.controller;

import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// CORRIGÉ — le contrôle d'accès est appliqué côté SERVEUR, seule frontière
// de confiance valable : le masquage côté client reste une commodité
// d'UX, jamais un mécanisme de sécurité. @PreAuthorize échoue "fermé" par
// défaut (403) si le rôle n'est pas présent.
@Controller
public class Fixed {

    private final ReportRepository reports;

    public Fixed(ReportRepository reports) {
        this.reports = reports;
    }

    @GetMapping("/api/reports/financial")
    @ResponseBody
    @PreAuthorize("hasRole('FINANCE_ADMIN')")
    public String financialReport() {
        return reports.generateFinancialReport();
    }

    interface ReportRepository { String generateFinancialReport(); }
}
