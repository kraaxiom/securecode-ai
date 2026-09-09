// Vulnérable : Secrets exposés dans le code ou la configuration cloud (CWE-798)
// La clé API et la chaîne de connexion sont codées en dur dans le code source,
// donc committées dans le dépôt Git et visibles par quiconque y a accès,
// même après suppression ultérieure (elles restent dans l'historique).
using Microsoft.AspNetCore.Builder;
using Stripe;

var builder = WebApplication.CreateBuilder(args);

// Clé secrète Stripe codée en dur — VULNÉRABLE : secret en clair dans le code
StripeConfiguration.ApiKey = "sk_live_EXAMPLE_NOT_A_REAL_KEY";

// Chaîne de connexion SQL avec identifiants en clair — VULNÉRABLE
const string connectionString =
    "Server=db.example.com;Database=prod;User Id=admin;Password=REPLACE_ME_VULNERABLE_PATTERN;";

builder.Services.AddSingleton(new SqlConnectionFactory(connectionString));

var app = builder.Build();
app.MapGet("/", () => "OK");
app.Run();

public class SqlConnectionFactory
{
    private readonly string _connectionString;
    public SqlConnectionFactory(string connectionString) => _connectionString = connectionString;
}
