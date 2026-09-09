<?php
// Faille : un compte administrateur est créé avec un mot de passe fixe et
// prévisible ("admin123"), documenté et réutilisé tel quel en production,
// sans obligation de le changer. Un attaquant qui le connaît obtient un
// accès administrateur immédiat.

namespace Database\Seeders;

use App\Models\User;
use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\Hash;

class AdminSeeder extends Seeder
{
    public function run(): void
    {
        User::create([
            'email' => 'admin@example.com',
            'password' => Hash::make('admin123'),
            'role' => 'admin',
        ]);
    }
}
