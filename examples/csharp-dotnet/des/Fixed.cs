using System;
using System.IO;
using System.Security.Cryptography;

namespace SecurityAiSkill.Examples.Crypto.Des
{
    /// <summary>
    /// EXEMPLE CORRIGÉ — Chiffrement symétrique avec AES-256-GCM (mode authentifié).
    /// Remplace DES/3DES par une primitive moderne, avec un nonce unique par opération.
    /// </summary>
    public class ModernFileEncryptor
    {
        private const int KeySizeBytes = 32;   // AES-256
        private const int NonceSizeBytes = 12; // Taille recommandée pour GCM
        private const int TagSizeBytes = 16;   // Tag d'authentification GCM (128 bits)

        /// <summary>
        /// SÉCURISÉ : AES-256 en mode GCM (AEAD).
        /// - Clé de 256 bits : résistant à la force brute avec les capacités actuelles et prévisibles.
        /// - Mode GCM authentifié : garantit à la fois confidentialité ET intégrité
        ///   (détecte toute modification du texte chiffré, contrairement à CBC seul).
        /// - Un nonce (IV) de 12 octets, généré aléatoirement, est utilisé une seule fois par clé
        ///   ("nonce" = number used once) : il ne doit JAMAIS être réutilisé avec la même clé.
        /// </summary>
        public byte[] Encrypt(byte[] plaintext, byte[] key)
        {
            if (key.Length != KeySizeBytes)
            {
                throw new ArgumentException($"La clé AES-256 doit faire {KeySizeBytes} octets.", nameof(key));
            }

            // Nonce généré via un CSPRNG, unique à chaque appel.
            byte[] nonce = RandomNumberGenerator.GetBytes(NonceSizeBytes);
            byte[] ciphertext = new byte[plaintext.Length];
            byte[] tag = new byte[TagSizeBytes];

            using var aesGcm = new AesGcm(key, TagSizeBytes);
            aesGcm.Encrypt(nonce, plaintext, ciphertext, tag);

            // On concatène nonce + tag + ciphertext pour le stockage/transport ;
            // les trois éléments sont nécessaires au déchiffrement.
            using var output = new MemoryStream();
            output.Write(nonce, 0, nonce.Length);
            output.Write(tag, 0, tag.Length);
            output.Write(ciphertext, 0, ciphertext.Length);
            return output.ToArray();
        }

        /// <summary>
        /// SÉCURISÉ : déchiffrement AES-256-GCM avec vérification du tag d'authentification.
        /// Toute altération du texte chiffré fait échouer l'opération (CryptographicException).
        /// </summary>
        public byte[] Decrypt(byte[] payload, byte[] key)
        {
            if (key.Length != KeySizeBytes)
            {
                throw new ArgumentException($"La clé AES-256 doit faire {KeySizeBytes} octets.", nameof(key));
            }

            byte[] nonce = payload[..NonceSizeBytes];
            byte[] tag = payload[NonceSizeBytes..(NonceSizeBytes + TagSizeBytes)];
            byte[] ciphertext = payload[(NonceSizeBytes + TagSizeBytes)..];
            byte[] plaintext = new byte[ciphertext.Length];

            using var aesGcm = new AesGcm(key, TagSizeBytes);
            aesGcm.Decrypt(nonce, ciphertext, tag, plaintext);
            return plaintext;
        }

        /// <summary>
        /// Génère une clé AES-256 à partir d'un CSPRNG. En production, la clé doit être
        /// émise et stockée par un gestionnaire de secrets (Azure Key Vault, AWS KMS, HashiCorp Vault),
        /// jamais codée en dur dans le code source.
        /// </summary>
        public static byte[] GenerateKey() => RandomNumberGenerator.GetBytes(KeySizeBytes);
    }
}
