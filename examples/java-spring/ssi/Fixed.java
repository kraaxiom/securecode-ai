// Corrigé : Server-Side Includes (SSI) Injection (CWE-97)
// Le contenu utilisateur est échappé et toute séquence de directive SSI
// résiduelle est neutralisée avant écriture. Idéalement, SSI est désactivé
// sur ce répertoire et le contenu dynamique est rendu via un moteur de
// templates applicatif plutôt que via un fichier .shtml.
package com.example.security.ssi;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.util.HtmlUtils;

import java.io.FileWriter;
import java.io.IOException;

@RestController
public class SsiController {

    private static final String COMMENTS_FILE = "/var/www/public/comments.shtml";

    @PostMapping("/comments")
    public String addComment(@RequestParam String comment) throws IOException {
        // Échappement HTML puis neutralisation de toute séquence de directive SSI
        String safeComment = stripSsiDirectives(HtmlUtils.htmlEscape(comment));

        try (FileWriter writer = new FileWriter(COMMENTS_FILE, true)) {
            writer.write("<p>" + safeComment + "</p>\n");
        }
        // Préférable : générer la page via un moteur de templates (Thymeleaf)
        // et désactiver SSI (mod_include / exec) sur ce répertoire côté serveur web.
        return "Commentaire ajouté";
    }

    private String stripSsiDirectives(String input) {
        // Neutralise toute directive SSI résiduelle du type <!--#...-->
        return input.replace("<!--#", "&lt;!--#");
    }
}
