// Corrigé : Secrets exposés dans le code ou la configuration cloud (CWE-798)
// Aucun secret n'apparaît dans le code source. Les valeurs sensibles sont
// récupérées à l'exécution depuis Azure Key Vault, avec repli sur des
// variables d'environnement injectées par la plateforme d'hébergement.
using Azure.Identity;
using Microsoft.AspNetCore.Builder;
using Stripe;

var builder = WebApplication.CreateBuilder(args);

// Charge la configuration depuis Azure Key Vault via une identité managée
// (aucune credential Azure codée en dur, aucun secret en clair).
var keyVaultUri = builder.Configuration["KeyVault:Uri"];
if (!string.IsNullOrEmpty(keyVaultUri))
{
    builder.Configuration.AddAzureKeyVault(
        new Uri(keyVaultUri),
        new DefaultAzureCredential());
}

// La clé API est lue depuis la configuration (Key Vault en priorité,
// variable d'environnement en repli) — jamais stockée dans le code.
StripeConfiguration.ApiKey = builder.Configuration["Stripe:SecretKey"]
    ?? throw new InvalidOperationException("Stripe:SecretKey non configuré (secret manquant).");

var connectionString = builder.Configuration["Db:ConnectionString"]
    ?? throw new InvalidOperationException("Db:ConnectionString non configuré (secret manquant).");

builder.Services.AddSingleton(new SqlConnectionFactory(connectionString));

var app = builder.Build();
app.MapGet("/", () => "OK");
app.Run();

public class SqlConnectionFactory
{
    private readonly string _connectionString;
    public SqlConnectionFactory(string connectionString) => _connectionString = connectionString;
}
