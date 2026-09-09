// Vulnérable : SMTP / Email Header Injection (CWE-93)
// Les en-têtes de l'email (Reply-To, Subject) sont construits par concaténation
// directe d'entrées utilisateur non filtrées. Un attaquant peut injecter des
// sauts de ligne (\r\n) pour ajouter des en-têtes arbitraires (Cc, Bcc) ou un
// second corps de message, transformant le formulaire en relais de spam.
package com.example.security.smtp;

import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class SmtpInjectionService {

    private final JavaMailSender mailSender;

    public SmtpInjectionService(JavaMailSender mailSender) {
        this.mailSender = mailSender;
    }

    @PostMapping("/api/contact")
    public String sendContactForm(@RequestParam String name,
                                   @RequestParam String subject,
                                   @RequestParam String body) {
        // name = "Bob\r\nBcc: victime@interne.com" injecte un en-tête Bcc caché.
        String replyToHeader = "\"" + name + "\" <contact@example.com>";

        SimpleMailMessage message = new SimpleMailMessage();
        message.setTo("dest@example.com");
        message.setReplyTo(replyToHeader);
        message.setSubject(subject);
        message.setText(body);

        mailSender.send(message);
        return "Message envoyé";
    }
}
