// Vulnérable : Code Injection (CWE-94)
// Une expression utilisateur est évaluée dynamiquement via un moteur de
// script (ici JavaScript via Nashorn/GraalJS), permettant à un attaquant
// d'exécuter du code arbitraire sur le serveur.
package com.example.demo.controller;

import org.springframework.web.bind.annotation.*;

import javax.script.ScriptEngine;
import javax.script.ScriptEngineManager;

@RestController
@RequestMapping("/api/calc")
public class CodeInjectionController {

    @GetMapping("/eval")
    public Object eval(@RequestParam String formula) throws Exception {
        // Évaluation dynamique d'une chaîne fournie par l'utilisateur.
        ScriptEngine engine = new ScriptEngineManager().getEngineByName("graal.js");
        return engine.eval(formula);
    }
}
