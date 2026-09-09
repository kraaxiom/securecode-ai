<?php
/**
 * VULNÉRABLE — Utilisation de MD5 pour le hachage de mots de passe.
 *
 * Faille : MD5 est cryptographiquement cassé depuis 2004 (collisions en
 * quelques secondes) et n'offre aucune résistance à la force brute moderne
 * (GPU/ASIC) — inadapté au stockage de mots de passe.
 *
 * CWE-328: Use of Weak Hash
 * OWASP A02:2021 - Cryptographic Failures
 */

namespace App\Services;

class UserRegistrationService
{
    public function register(string $email, string $plainPassword): array
    {
        // VULNÉRABLE : hachage MD5 sans sel, cassable par table arc-en-ciel.
        $passwordHash = md5($plainPassword);

        return \App\Models\User::create([
            'email' => $email,
            'password' => $passwordHash,
        ])->toArray();
    }

    public function checkLogin(string $plainPassword, string $storedHash): bool
    {
        // VULNÉRABLE : comparaison de hachages MD5, non résistante au timing attack en plus.
        return md5($plainPassword) === $storedHash;
    }
}
