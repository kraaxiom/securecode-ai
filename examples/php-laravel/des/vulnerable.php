<?php
/**
 * VULNÉRABLE — Utilisation de DES/3DES pour chiffrer des données sensibles.
 *
 * Faille : DES a une clé effective de 56 bits, cassable par force brute en
 * quelques heures. Le 3DES (DES-EDE3) reste vulnérable à l'attaque Sweet32
 * sur les blocs de 64 bits et est retiré des standards NIST depuis 2023.
 *
 * CWE-327: Use of a Broken or Risky Cryptographic Algorithm
 * OWASP A02:2021 - Cryptographic Failures
 */

namespace App\Services;

class PayrollExportService
{
    // Clé illustrative — ne représente aucun secret réel.
    private string $key = 'CLE_SECRETE_EXEMPLE_16B';

    /**
     * Chiffre un export de bulletins de paie avant archivage.
     */
    public function encryptExport(string $csvContent): string
    {
        // VULNÉRABLE : algorithme DES/3DES obsolète, bloc de 64 bits.
        $iv = substr(md5(uniqid()), 0, 8); // IV faible en plus du reste
        $cipher = openssl_encrypt($csvContent, 'des-ede3-cbc', $this->key, 0, $iv);

        return base64_encode($cipher);
    }

    public function decryptExport(string $encoded, string $iv): string
    {
        $cipher = base64_decode($encoded);

        // VULNÉRABLE : déchiffrement avec le même algorithme cassé.
        return openssl_decrypt($cipher, 'des-ede3-cbc', $this->key, 0, $iv);
    }
}
