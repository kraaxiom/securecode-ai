package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — Stored XSS (CWE-79)
// Le commentaire est stocké en base tel quel, puis réaffiché à TOUS les
// visiteurs de la page produit sans échappement. Contrairement au reflected
// XSS, la charge utile persiste et touche chaque utilisateur qui consulte
// la page, sans qu'un lien piégé ne soit nécessaire.
@Controller
public class Vulnerable {

    private final CommentRepository repository;

    public Vulnerable(CommentRepository repository) {
        this.repository = repository;
    }

    @PostMapping("/products/{id}/comments")
    public String addComment(@PathVariable Long id, @RequestParam String text) {
        repository.save(id, text); // aucune sanitisation avant stockage
        return "redirect:/products/" + id;
    }

    @GetMapping("/products/{id}")
    public String viewProduct(@PathVariable Long id, Model model) {
        // Les commentaires sont rendus via th:utext dans la vue
        model.addAttribute("comments", repository.findByProduct(id));
        return "product-page";
    }

    interface CommentRepository {
        void save(Long productId, String text);
        java.util.List<String> findByProduct(Long productId);
    }
}
