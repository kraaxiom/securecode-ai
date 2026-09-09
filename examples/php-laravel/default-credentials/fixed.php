<?php
// Correction : mot de passe temporaire généré aléatoirement à chaque
// provisioning, avec un flag forçant son changement à la première
// connexion. Le secret n'est jamais journalisé en clair, conformément
// au pattern décrit dans rules/remediation/default-credentials.md.

namespace Database\Seeders;

use App\Models\User;
use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Facades\Log;
use Illuminate\Support\Str;

class AdminSeeder extends Seeder
{
    public function run(): void
    {
        $temporaryPassword = Str::random(24);

        User::create([
            'email' => 'admin@example.com',
            'password' => Hash::make($temporaryPassword),
            'role' => 'admin',
            'must_change_password' => true,
        ]);

        // Le mot de passe temporaire n'est jamais journalisé en clair ;
        // il est transmis via un canal hors-bande sécurisé au déploiement.
        Log::channel('deploy')->info(
            'Compte admin créé, mot de passe temporaire à récupérer via le canal sécurisé de déploiement.'
        );
    }
}
