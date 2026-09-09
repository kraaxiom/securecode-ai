// Corrigé : SMTP / Email Header Injection (CWE-93)
// Tout caractère de contrôle (\r, \n) est supprimé des champs libres avant
// construction de l'email, et l'API MimeMessageHelper de Spring/JavaMail est
// utilisée : elle échappe correctement les en-têtes, empêchant l'ajout
// d'en-têtes arbitraires ou d'un second corps de message.
package com.example.security.smtp;

import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.mail.javamail.MimeMessageHelper;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import jakarta.mail.internet.MimeMessage;

@RestController
public class SmtpInjectionService {

    private final JavaMailSender mailSender;

    public SmtpInjectionService(JavaMailSender mailSender) {
        this.mailSender = mailSender;
    }

    // Supprime tout caractère de contrôle susceptible d'injecter un en-tête.
    private static String stripControlChars(String value) {
        return value == null ? "" : value.replace("\r", "").replace("\n", "");
    }

    @PostMapping("/api/contact")
    public String sendContactForm(@RequestParam String name,
                                   @RequestParam String subject,
                                   @RequestParam String body) throws Exception {
        String safeName = stripControlChars(name);
        String safeSubject = stripControlChars(subject);

        MimeMessage mimeMessage = mailSender.createMimeMessage();
        // MimeMessageHelper échappe correctement les en-têtes MIME.
        MimeMessageHelper helper = new MimeMessageHelper(mimeMessage, false, "UTF-8");
        helper.setTo("dest@example.com");
        helper.setReplyTo("contact@example.com", safeName);
        helper.setSubject(safeSubject);
        helper.setText(body, false);

        mailSender.send(mimeMessage);
        return "Message envoyé";
    }
}
