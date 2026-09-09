package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — Blind XSS (CWE-79)
// Le contenu du formulaire de support est stocké tel quel, puis affiché
// SANS échappement dans le back-office consulté par un administrateur.
// L'attaquant ne voit jamais l'exécution : le payload se déclenche plus
// tard, dans le navigateur de la victime (souvent privilégiée).
@Controller
public class Vulnerable {

    private final SupportTicketRepository repository;

    public Vulnerable(SupportTicketRepository repository) {
        this.repository = repository;
    }

    @PostMapping("/support/ticket")
    public String submitTicket(@RequestParam String subject, @RequestParam String message) {
        // Aucune validation ni sanitisation avant stockage
        repository.save(new SupportTicket(subject, message));
        return "redirect:/support/thanks";
    }

    @GetMapping("/admin/tickets/{id}")
    public String viewTicket(@PathVariable Long id, Model model) {
        SupportTicket ticket = repository.findById(id);
        // Le message brut est injecté dans le modèle Thymeleaf avec th:utext
        // (unescaped text) côté template admin/ticket.html : exécution du
        // script au chargement de la page par l'administrateur.
        model.addAttribute("subject", ticket.getSubject());
        model.addAttribute("message", ticket.getMessage()); // rendu via th:utext
        return "admin/ticket";
    }

    interface SupportTicketRepository {
        void save(SupportTicket t);
        SupportTicket findById(Long id);
    }

    static class SupportTicket {
        private final String subject;
        private final String message;
        SupportTicket(String subject, String message) { this.subject = subject; this.message = message; }
        String getSubject() { return subject; }
        String getMessage() { return message; }
    }
}
