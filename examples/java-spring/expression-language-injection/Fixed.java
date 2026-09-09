// Corrigé : Expression Language Injection (CWE-917)
// L'entrée utilisateur n'est plus jamais transformée en code évalué. Seule
// une expression statique, écrite par le développeur, est analysée par le
// parseur SpEL ; la donnée utilisateur est fournie uniquement comme variable
// du contexte d'évaluation, jamais comme fragment de code.
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

    // L'expression est fixée par le développeur : elle ne fait que
    // manipuler la variable "input", jamais du code arbitraire.
    private static final String STATIC_EXPRESSION = "#input.length()";

    @GetMapping("/api/calc")
    public Object calc(@RequestParam String value) {
        ExpressionParser parser = new SpelExpressionParser();
        Expression exp = parser.parseExpression(STATIC_EXPRESSION);

        StandardEvaluationContext context = new StandardEvaluationContext();
        // La valeur utilisateur est liée en tant que simple donnée, pas
        // comme code évalué : elle ne peut jamais changer la sémantique
        // de l'expression exécutée.
        context.setVariable("input", value);

        return exp.getValue(context);
    }
}
