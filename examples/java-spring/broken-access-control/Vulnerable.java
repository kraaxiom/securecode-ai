package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — Broken Access Control (CWE-284 : Improper Access Control)
// Le contrôle d'accès repose UNIQUEMENT sur le masquage côté client (le
// lien "Panneau d'administration" n'est affiché que si isAdmin=true dans
// le frontend), mais l'API serveur elle-même ne réapplique aucune
// vérification. N'importe quel client HTTP peut appeler directement
// l'endpoint sans jamais avoir vu le lien.
@Controller
public class Vulnerable {

    private final ReportRepository reports;

    public Vulnerable(ReportRepository reports) {
        this.reports = reports;
    }

    @GetMapping("/api/reports/financial")
    @ResponseBody
    public String financialReport() {
        // Aucune vérification de rôle côté serveur : la "protection" n'existe
        // que dans l'interface utilisateur (menu masqué), pas dans l'API.
        return reports.generateFinancialReport();
    }

    interface ReportRepository { String generateFinancialReport(); }
}
