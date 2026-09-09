package com.example.demo.controller;

import org.owasp.encoder.Encode;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;

// CORRIGÉ — le contenu utilisateur reste stocké brut (pour préserver la
// donnée d'origine) mais n'est JAMAIS rendu en unescaped : le template
// admin utilise th:text (échappement automatique par défaut de Thymeleaf)
// et, en défense en profondeur, la valeur est explicitement encodée pour
// le contexte HTML avant d'être exposée au modèle. Une politique CSP
// stricte est également recommandée sur les pages du back-office.
@Controller
public class Fixed {

    private final SupportTicketRepository repository;

    public Fixed(SupportTicketRepository repository) {
        this.repository = repository;
    }

    @PostMapping("/support/ticket")
    public String submitTicket(@RequestParam String subject, @RequestParam String message) {
        // Validation de longueur / format avant stockage (défense en profondeur)
        if (subject.length() > 200 || message.length() > 5000) {
            throw new IllegalArgumentException("Contenu trop long");
        }
        repository.save(new SupportTicket(subject, message));
        return "redirect:/support/thanks";
    }

    @GetMapping("/admin/tickets/{id}")
    public String viewTicket(@PathVariable Long id, Model model) {
        SupportTicket ticket = repository.findById(id);
        // Encodage explicite pour le contexte HTML : neutralise tout script
        // stocké, quel que soit le moteur de template utilisé côté vue
        // (défense en profondeur en plus de th:text côté Thymeleaf).
        model.addAttribute("subject", Encode.forHtml(ticket.getSubject()));
        model.addAttribute("message", Encode.forHtml(ticket.getMessage())); // rendu via th:text
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
