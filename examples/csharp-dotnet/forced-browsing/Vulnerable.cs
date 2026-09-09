// Vulnérable : Forced Browsing (CWE-425)
// Les exports de sauvegarde sont servis depuis un répertoire statique public,
// sans middleware d'authentification ni d'autorisation : seule l'absence
// de lien visible dans l'interface protège la ressource.
using Microsoft.AspNetCore.Builder;

var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

// Le contenu de wwwroot/backups/ est exposé publiquement à quiconque devine
// ou énumère un nom de fichier, sans aucun contrôle d'accès côté serveur.
app.UseStaticFiles(); // sert wwwroot/, y compris wwwroot/backups/export-2026-01.zip

app.Run();
