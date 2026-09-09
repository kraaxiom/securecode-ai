<?php
/**
 * CORRIGÉ — Remplacement de DES/3DES par AES-256-GCM (chiffrement authentifié).
 *
 * CWE-327: Use of a Broken or Risky Cryptographic Algorithm
 */

namespace App\Services;

class PayrollExportService
{
    // La clé (32 octets pour AES-256) est chargée depuis un gestionnaire de
    // secrets / variable d'environnement — jamais codée en dur.
    private string $key;

    public function __construct()
    {
        $key = env('PAYROLL_ENCRYPTION_KEY');
        if (!$key) {
            throw new \RuntimeException('PAYROLL_ENCRYPTION_KEY manquant dans l\'environnement');
        }
        $this->key = base64_decode($key);
    }

    /**
     * Chiffre un export de bulletins de paie avant archivage.
     */
    public function encryptExport(string $csvContent): string
    {
        // CORRIGÉ : AES-256-GCM, chiffrement authentifié avec IV aléatoire unique.
        $iv = random_bytes(12);
        $tag = '';
        $cipher = openssl_encrypt(
            $csvContent,
            'aes-256-gcm',
            $this->key,
            OPENSSL_RAW_DATA,
            $iv,
            $tag
        );

        // IV et tag d'authentification doivent être stockés avec le texte chiffré.
        return base64_encode($iv . $tag . $cipher);
    }

    public function decryptExport(string $encoded): string
    {
        $raw = base64_decode($encoded);
        $iv = substr($raw, 0, 12);
        $tag = substr($raw, 12, 16);
        $cipher = substr($raw, 28);

        $plain = openssl_decrypt(
            $cipher,
            'aes-256-gcm',
            $this->key,
            OPENSSL_RAW_DATA,
            $iv,
            $tag
        );

        if ($plain === false) {
            throw new \RuntimeException('Échec du déchiffrement ou intégrité compromise');
        }

        return $plain;
    }
}
