package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// VULNÉRABLE — Broken Object Level Authorization (BOLA) (CWE-639 :
// Authorization Bypass Through User-Controlled Key)
// La commande est récupérée uniquement à partir de l'ID fourni par le
// client dans l'URL, sans jamais vérifier qu'elle appartient bien à
// l'utilisateur authentifié. Changer l'ID suffit à lire ou modifier la
// commande d'un autre utilisateur.
@Controller
public class Vulnerable {

    private final OrderRepository orders;

    public Vulnerable(OrderRepository orders) {
        this.orders = orders;
    }

    @GetMapping("/api/orders/{id}")
    @ResponseBody
    public Order getOrder(@PathVariable Long id) {
        return orders.findById(id); // aucune vérification de propriétaire
    }

    @PutMapping("/api/orders/{id}")
    @ResponseBody
    public Order updateOrder(@PathVariable Long id, @RequestBody Order update) {
        return orders.update(id, update); // idem : n'importe qui peut modifier n'importe quelle commande
    }

    interface OrderRepository {
        Order findById(Long id);
        Order update(Long id, Order update);
    }
    static class Order { Long id; Long ownerId; String status; }
}
