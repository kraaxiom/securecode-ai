// Vulnérable : Default Credentials (CWE-1392)
// Le compte administrateur est provisionné avec un mot de passe fixe et
// documenté ("admin123"), sans jamais forcer son changement, ce qui permet
// à un attaquant connaissant la documentation d'obtenir un accès administrateur.
using Microsoft.AspNetCore.Mvc;

public class AdminProvisioningService
{
    private readonly AppDbContext _db;
    private readonly IPasswordHasher _hasher;

    public AdminProvisioningService(AppDbContext db, IPasswordHasher hasher)
    {
        _db = db;
        _hasher = hasher;
    }

    public async Task ProvisionAdminAccountAsync()
    {
        var admin = new User
        {
            Email = "admin@example.com",
            PasswordHash = _hasher.Hash("admin123"), // mot de passe par défaut en dur
            Role = "admin"
        };

        _db.Users.Add(admin);
        await _db.SaveChangesAsync();
        // Aucun flag forçant le changement de mot de passe à la première connexion.
    }
}

public class User
{
    public string Email { get; set; } = "";
    public string PasswordHash { get; set; } = "";
    public string Role { get; set; } = "";
}
