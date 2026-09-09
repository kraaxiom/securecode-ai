// Vulnérable : XML External Entity (XXE) Injection (CWE-611)
// Le DocumentBuilderFactory par défaut ne désactive pas le traitement des
// DTD ni la résolution des entités externes. Un attaquant peut soumettre un
// document XML avec une entité externe pointant vers un fichier local ou une
// URL interne, entraînant une divulgation d'informations ou du SSRF.
package com.example.security.xxe;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import org.w3c.dom.Document;

import javax.xml.parsers.DocumentBuilder;
import javax.xml.parsers.DocumentBuilderFactory;
import java.io.ByteArrayInputStream;
import java.nio.charset.StandardCharsets;

@RestController
public class XxeController {

    @PostMapping("/import")
    public String importXml(@RequestBody String userSuppliedXml) throws Exception {
        DocumentBuilderFactory dbf = DocumentBuilderFactory.newInstance();
        // Configuration par défaut : DTD et entités externes autorisées
        DocumentBuilder builder = dbf.newDocumentBuilder();

        Document doc = builder.parse(
                new ByteArrayInputStream(userSuppliedXml.getBytes(StandardCharsets.UTF_8)));

        return doc.getDocumentElement().getTextContent();
    }
}
