// Vulnérable : Server-Side Includes (SSI) Injection (CWE-97)
// Le commentaire utilisateur est écrit tel quel dans un fichier .shtml servi
// par le serveur web avec SSI activé. Un attaquant peut injecter une
// directive `<!--#exec cmd="..."-->` qui sera interprétée à l'affichage.
package com.example.security.ssi;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.io.FileWriter;
import java.io.IOException;

@RestController
public class SsiController {

    private static final String COMMENTS_FILE = "/var/www/public/comments.shtml";

    @PostMapping("/comments")
    public String addComment(@RequestParam String comment) throws IOException {
        // Écriture directe de l'entrée utilisateur dans un fichier interprété par SSI
        try (FileWriter writer = new FileWriter(COMMENTS_FILE, true)) {
            writer.write("<p>" + comment + "</p>\n");
        }
        return "Commentaire ajouté";
    }
}
