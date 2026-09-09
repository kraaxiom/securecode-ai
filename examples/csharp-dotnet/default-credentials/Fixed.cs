// Corrigé : Default Credentials (CWE-1392)
// Un mot de passe temporaire est généré aléatoirement à chaque provisioning,
// transmis hors-bande, et l'accès normal reste bloqué tant que l'administrateur
// n'a pas changé ce mot de passe temporaire (flag MustChangePassword).
using System.Security.Cryptography;

public class AdminProvisioningService
{
    private readonly AppDbContext _db;
    private readonly IPasswordHasher _hasher;
    private readonly ISecureOutOfBandChannel _outOfBandChannel;
    private readonly ILogger<AdminProvisioningService> _logger;

    public AdminProvisioningService(
        AppDbContext db,
        IPasswordHasher hasher,
        ISecureOutOfBandChannel outOfBandChannel,
        ILogger<AdminProvisioningService> logger)
    {
        _db = db;
        _hasher = hasher;
        _outOfBandChannel = outOfBandChannel;
        _logger = logger;
    }

    public async Task ProvisionAdminAccountAsync()
    {
        var temporaryPassword = GenerateSecureTemporaryPassword();

        var admin = new User
        {
            Email = "admin@example.com",
            PasswordHash = _hasher.Hash(temporaryPassword),
            Role = "admin",
            MustChangePassword = true
        };

        _db.Users.Add(admin);
        await _db.SaveChangesAsync();

        // Transmission hors-bande sécurisée, jamais journalisée en clair.
        await _outOfBandChannel.SendAsync(admin.Email, temporaryPassword);
        _logger.LogInformation("Compte admin créé, mot de passe temporaire transmis via canal sécurisé.");
    }

    private static string GenerateSecureTemporaryPassword()
    {
        var bytes = RandomNumberGenerator.GetBytes(24);
        return Convert.ToBase64String(bytes);
    }
}

public interface ISecureOutOfBandChannel
{
    Task SendAsync(string recipient, string secret);
}

public class User
{
    public string Email { get; set; } = "";
    public string PasswordHash { get; set; } = "";
    public string Role { get; set; } = "";
    public bool MustChangePassword { get; set; }
}
