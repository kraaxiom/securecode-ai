// Vulnérable : Expression Language Injection (CWE-917)
// L'expression SpEL est construite par concaténation directe de l'entrée
// utilisateur, puis évaluée. Un attaquant peut injecter du code SpEL
// arbitraire (ex: accès à T(java.lang.Runtime)) et l'exécuter côté serveur.
package com.example.security.el;

import org.springframework.expression.Expression;
import org.springframework.expression.ExpressionParser;
import org.springframework.expression.spel.standard.SpelExpressionParser;
import org.springframework.expression.spel.support.StandardEvaluationContext;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ElInjectionController {

    // Endpoint censé calculer une expression mathématique simple fournie
    // par l'utilisateur, ex: /api/calc?expr=2+2
    @GetMapping("/api/calc")
    public Object calc(@RequestParam String expr) {
        ExpressionParser parser = new SpelExpressionParser();
        StandardEvaluationContext context = new StandardEvaluationContext();

        // L'entrée utilisateur devient directement le code évalué : elle
        // peut contenir n'importe quelle expression SpEL, y compris des
        // appels à des classes Java arbitraires.
        Expression exp = parser.parseExpression(expr);
        return exp.getValue(context);
    }
}
