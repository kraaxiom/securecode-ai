// Vulnérable : XML Injection (CWE-91)
// Le document XML est construit par concaténation de chaînes incluant
// directement l'entrée utilisateur. Un attaquant peut injecter des balises
// ou des caractères spéciaux XML pour altérer la structure du document
// (ajout de nœuds, falsification de champs non prévus par l'application).
package com.example.security.xmlinjection;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class XmlInjectionController {

    @PostMapping("/users/register")
    public String register(@RequestParam String name, @RequestParam String role) {
        // Concaténation directe : aucun échappement des caractères spéciaux XML
        String xml = "<user><name>" + name + "</name><role>" + role + "</role></user>";
        return persist(xml);
    }

    private String persist(String xml) {
        // Simule l'enregistrement du document XML généré
        return xml;
    }
}
