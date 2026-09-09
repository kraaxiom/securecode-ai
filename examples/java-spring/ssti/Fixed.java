// Corrigé : Server-Side Template Injection (SSTI) (CWE-1336)
// Le template est chargé depuis un fichier statique et compilé une seule fois.
// L'entrée utilisateur n'est jamais injectée dans la structure du template :
// elle transite uniquement via le modèle de variables (`model`), qui est
// automatiquement échappé par le moteur de rendu.
package com.example.security.ssti;

import freemarker.template.Configuration;
import freemarker.template.Template;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.io.StringWriter;
import java.util.HashMap;
import java.util.Map;

@RestController
public class SstiController {

    private final Configuration freemarkerConfig;

    public SstiController(Configuration freemarkerConfig) {
        this.freemarkerConfig = freemarkerConfig;
    }

    @GetMapping("/welcome")
    public String welcome(@RequestParam String name) throws Exception {
        // Template statique, préexistant sur disque (welcome.ftl : "Bonjour ${name}, bienvenue !")
        Template template = freemarkerConfig.getTemplate("welcome.ftl");

        // La donnée utilisateur ne circule que comme variable de contexte, jamais comme code de template
        Map<String, Object> model = new HashMap<>();
        model.put("name", name);

        StringWriter out = new StringWriter();
        template.process(model, out);
        return out.toString();
    }
}
