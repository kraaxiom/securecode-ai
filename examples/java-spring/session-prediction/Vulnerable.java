package com.example.demo.security;

import java.util.concurrent.atomic.AtomicLong;

// VULNÉRABLE — Session Prediction (CWE-330 : Use of Insufficiently Random
// Values)
// L'identifiant de session est un compteur séquentiel (ou dérivé d'un
// timestamp prévisible). Un attaquant qui observe un identifiant de
// session valide (le sien) peut en déduire ou énumérer d'autres
// identifiants valides, et ainsi usurper la session d'autres utilisateurs
// sans jamais connaître leur mot de passe.
public class Vulnerable {

    private final AtomicLong counter = new AtomicLong(100000);

    public String generateSessionId() {
        // Séquentiel et prévisible : session-100001, session-100002, ...
        return "session-" + counter.incrementAndGet();
    }
}
