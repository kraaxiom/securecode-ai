// Vulnérable : Server-Side Template Injection (SSTI) (CWE-1336)
// Le texte du template FreeMarker est construit par concaténation d'une
// entrée utilisateur avant compilation. Un attaquant peut injecter une
// directive FreeMarker (ex: ${"freemarker.template.utility.Execute"?new()("id")})
// pour exécuter du code arbitraire côté serveur.
package com.example.security.ssti;

import freemarker.template.Configuration;
import freemarker.template.Template;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.io.StringReader;
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
        // Construction dynamique du template avec l'entrée utilisateur intégrée à sa structure
        String templateText = "Bonjour " + name + ", bienvenue !";
        Template template = new Template("dyn", new StringReader(templateText), freemarkerConfig);

        Map<String, Object> model = new HashMap<>();
        StringWriter out = new StringWriter();
        template.process(model, out);
        return out.toString();
    }
}
