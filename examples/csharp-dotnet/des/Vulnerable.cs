using System;
using System.IO;
using System.Security.Cryptography;
using System.Text;

namespace SecurityAiSkill.Examples.Crypto.Des
{
    /// <summary>
    /// EXEMPLE VULNÉRABLE — Chiffrement symétrique avec DES / 3DES.
    /// Ce fichier illustre une configuration à risque à des fins pédagogiques.
    /// Aucun code d'exploitation n'est fourni.
    /// </summary>
    public class LegacyFileEncryptor
    {
        /// <summary>
        /// VULNÉRABLE : TripleDESCryptoServiceProvider utilise une clé effective
        /// bien trop faible et un bloc de 64 bits.
        /// - DES : clé de 56 bits, cassable par force brute en quelques heures avec du matériel moderne.
        /// - 3DES (TripleDES) : bloc de 64 bits vulnérable à l'attaque Sweet32 (collisions de bloc
        ///   après ~2^32 blocs chiffrés avec la même clé), et retiré des recommandations NIST depuis 2023.
        /// - Le mode CBC utilisé ici sans authentification (pas de MAC/AEAD) expose aussi à des
        ///   attaques de type padding oracle si le texte chiffré est manipulable par un attaquant.
        /// </summary>
        public byte[] EncryptLegacy(byte[] plaintext, byte[] key, byte[] iv)
        {
            // Chaîne littérale "TripleDES" : indicateur de détection direct.
            using var tdes = TripleDES.Create();
            tdes.Key = key;   // Clé effective ~112-168 bits nominaux, mais cassée par Sweet32.
            tdes.IV = iv;
            tdes.Mode = CipherMode.CBC;
            tdes.Padding = PaddingMode.PKCS7;

            using var encryptor = tdes.CreateEncryptor();
            using var ms = new MemoryStream();
            using (var cs = new CryptoStream(ms, encryptor, CryptoStreamMode.Write))
            {
                cs.Write(plaintext, 0, plaintext.Length);
            }
            return ms.ToArray();
        }

        /// <summary>
        /// VULNÉRABLE : DESCryptoServiceProvider — algorithme DES pur (56 bits),
        /// cassable en quelques heures par force brute distribuée (voire minutes avec du matériel dédié).
        /// N'offre plus aucune garantie de confidentialité sérieuse.
        /// </summary>
        public byte[] EncryptWithRawDes(byte[] plaintext, byte[] key, byte[] iv)
        {
            using var des = DES.Create();
            des.Key = key; // Clé de 8 octets seulement -> 56 bits d'entropie utile.
            des.IV = iv;
            des.Mode = CipherMode.CBC;

            using var encryptor = des.CreateEncryptor();
            using var ms = new MemoryStream();
            using (var cs = new CryptoStream(ms, encryptor, CryptoStreamMode.Write))
            {
                cs.Write(plaintext, 0, plaintext.Length);
            }
            return ms.ToArray();
        }
    }
}
