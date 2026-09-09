package com.example.demo.controller;

import org.springframework.http.ResponseEntity;
import org.springframework.security.core.Authentication;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// CORRIGÉ — le filtrage d'appartenance (ownerId) est appliqué directement
// dans la requête de données elle-même, jamais en post-traitement après
// coup : la commande d'un autre utilisateur ne peut jamais être renvoyée
// ni modifiée, quel que soit l'ID demandé.
@Controller
public class Fixed {

    private final OrderRepository orders;

    public Fixed(OrderRepository orders) {
        this.orders = orders;
    }

    @GetMapping("/api/orders/{id}")
    @ResponseBody
    public ResponseEntity<Order> getOrder(@PathVariable Long id, Authentication auth) {
        Long currentUserId = (Long) auth.getPrincipal();
        Order order = orders.findByIdAndOwner(id, currentUserId);
        if (order == null) {
            return ResponseEntity.notFound().build();
        }
        return ResponseEntity.ok(order);
    }

    @PutMapping("/api/orders/{id}")
    @ResponseBody
    public ResponseEntity<Order> updateOrder(@PathVariable Long id, @RequestBody Order update, Authentication auth) {
        Long currentUserId = (Long) auth.getPrincipal();
        Order updated = orders.updateIfOwner(id, currentUserId, update);
        if (updated == null) {
            return ResponseEntity.notFound().build();
        }
        return ResponseEntity.ok(updated);
    }

    interface OrderRepository {
        Order findByIdAndOwner(Long id, Long ownerId);
        Order updateIfOwner(Long id, Long ownerId, Order update);
    }
    static class Order { Long id; Long ownerId; String status; }
}
