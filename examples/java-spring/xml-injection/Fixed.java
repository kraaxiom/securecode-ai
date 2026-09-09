// Corrigé : XML Injection (CWE-91)
// Le document XML est construit via l'API DOM, qui échappe automatiquement
// le contenu textuel des nœuds. L'entrée utilisateur ne peut plus modifier
// la structure du document, quel que soit son contenu (`<`, `>`, `&`, etc.).
package com.example.security.xmlinjection;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.w3c.dom.Document;
import org.w3c.dom.Element;

import javax.xml.parsers.DocumentBuilder;
import javax.xml.parsers.DocumentBuilderFactory;
import javax.xml.transform.OutputKeys;
import javax.xml.transform.Transformer;
import javax.xml.transform.TransformerFactory;
import javax.xml.transform.dom.DOMSource;
import javax.xml.transform.stream.StreamResult;
import java.io.StringWriter;

@RestController
public class XmlInjectionController {

    @PostMapping("/users/register")
    public String register(@RequestParam String name, @RequestParam String role) throws Exception {
        DocumentBuilderFactory dbf = DocumentBuilderFactory.newInstance();
        DocumentBuilder builder = dbf.newDocumentBuilder();
        Document doc = builder.newDocument();

        Element user = doc.createElement("user");
        Element nameEl = doc.createElement("name");
        nameEl.appendChild(doc.createTextNode(name)); // échappement automatique par l'API DOM
        Element roleEl = doc.createElement("role");
        roleEl.appendChild(doc.createTextNode(role)); // échappement automatique par l'API DOM

        user.appendChild(nameEl);
        user.appendChild(roleEl);
        doc.appendChild(user);

        return serialize(doc);
    }

    private String serialize(Document doc) throws Exception {
        Transformer transformer = TransformerFactory.newInstance().newTransformer();
        transformer.setOutputProperty(OutputKeys.OMIT_XML_DECLARATION, "yes");
        StringWriter writer = new StringWriter();
        transformer.transform(new DOMSource(doc), new StreamResult(writer));
        return writer.toString();
    }
}
