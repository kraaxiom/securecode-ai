// Corrigé : XML External Entity (XXE) Injection (CWE-611)
// Le traitement des DOCTYPE/DTD est explicitement interdit et la résolution
// des entités externes (générales et paramétriques) est désactivée sur le
// DocumentBuilderFactory, ce qui neutralise toute exploitation XXE quelle
// que soit l'entrée fournie.
package com.example.security.xxe;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import org.w3c.dom.Document;

import javax.xml.XMLConstants;
import javax.xml.parsers.DocumentBuilder;
import javax.xml.parsers.DocumentBuilderFactory;
import java.io.ByteArrayInputStream;
import java.nio.charset.StandardCharsets;

@RestController
public class XxeController {

    @PostMapping("/import")
    public String importXml(@RequestBody String userSuppliedXml) throws Exception {
        DocumentBuilderFactory dbf = DocumentBuilderFactory.newInstance();

        // Interdit toute déclaration DOCTYPE : neutralise la classe XXE dans son ensemble
        dbf.setFeature("http://apache.org/xml/features/disallow-doctype-decl", true);
        // Désactive la résolution des entités externes générales et paramétriques
        dbf.setFeature("http://xml.org/sax/features/external-general-entities", false);
        dbf.setFeature("http://xml.org/sax/features/external-parameter-entities", false);
        dbf.setFeature("http://apache.org/xml/features/nonvalidating/load-external-dtd", false);
        // Traitement sécurisé additionnel (limites de taille, protection XXE/entity expansion)
        dbf.setFeature(XMLConstants.FEATURE_SECURE_PROCESSING, true);
        dbf.setXIncludeAware(false);
        dbf.setExpandEntityReferences(false);

        DocumentBuilder builder = dbf.newDocumentBuilder();
        Document doc = builder.parse(
                new ByteArrayInputStream(userSuppliedXml.getBytes(StandardCharsets.UTF_8)));

        return doc.getDocumentElement().getTextContent();
    }
}
