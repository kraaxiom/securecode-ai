package com.example.demo.controller;

import org.owasp.encoder.Encode;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.stream.Collectors;

// CORRIGÉ — la donnée brute reste stockée telle quelle (pour ne pas perdre
// l'information saisie), mais elle est systématiquement encodée pour le
// contexte HTML au moment de l'affichage, et le template utilise th:text
// au lieu de th:utext. L'échappement à l'affichage (et non à la saisie)
// est la bonne pratique : il s'applique quel que soit le canal de lecture.
@Controller
public class Fixed {

    private final CommentRepository repository;

    public Fixed(CommentRepository repository) {
        this.repository = repository;
    }

    @PostMapping("/products/{id}/comments")
    public String addComment(@PathVariable Long id, @RequestParam String text) {
        if (text.length() > 2000) {
            throw new IllegalArgumentException("Commentaire trop long");
        }
        repository.save(id, text);
        return "redirect:/products/" + id;
    }

    @GetMapping("/products/{id}")
    public String viewProduct(@PathVariable Long id, Model model) {
        List<String> safeComments = repository.findByProduct(id).stream()
                .map(Encode::forHtml) // échappement systématique à l'affichage
                .collect(Collectors.toList());
        model.addAttribute("comments", safeComments); // rendu via th:text
        return "product-page";
    }

    interface CommentRepository {
        void save(Long productId, String text);
        List<String> findByProduct(Long productId);
    }
}
