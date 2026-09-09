// Vulnérable : Header Injection (CWE-113)
// Le nom de fichier fourni par l'utilisateur est concaténé directement dans
// l'en-tête Content-Disposition. Une valeur contenant des séquences CR/LF
// encodées (%0d%0a) peut injecter des en-têtes supplémentaires arbitraires
// dans la réponse HTTP (fixation de session, cache poisoning, etc.).
package com.example.security.header;

import jakarta.servlet.http.HttpServletResponse;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HeaderInjectionController {

    @GetMapping("/api/files/download")
    public void download(@RequestParam String filename, HttpServletResponse response) {
        // Aucune validation ni suppression des caractères de contrôle :
        // filename peut contenir \r\n et injecter de nouveaux en-têtes.
        response.setHeader("Content-Disposition", "attachment; filename=" + filename);
        response.setHeader("X-Requested-File", filename);
    }
}
