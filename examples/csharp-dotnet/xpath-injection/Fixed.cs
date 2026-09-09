// Corrigé : XPath Injection (CWE-643)
// L'expression XPath utilise des variables liées (via un XsltContext
// personnalisé) plutôt que la concaténation de chaînes : les valeurs
// utilisateur ne peuvent jamais altérer la structure de l'expression.
// Remarque : XPath reste déconseillé comme mécanisme d'authentification ;
// préférer une base de données avec hachage de mot de passe (ex: Argon2id).
using Microsoft.AspNetCore.Mvc;
using System.Xml;
using System.Xml.XPath;
using System.Xml.Xsl;

[ApiController]
[Route("api/auth")]
public class XmlAuthController : ControllerBase
{
    private readonly XmlDocument _users = LoadUsers();

    [HttpPost("login")]
    public IActionResult Login([FromForm] string user, [FromForm] string pass)
    {
        if (string.IsNullOrWhiteSpace(user) || string.IsNullOrWhiteSpace(pass))
            return BadRequest("Champs invalides.");

        var navigator = _users.CreateNavigator()!;
        var context = new BoundVariableContext();
        context.SetVariable("u", user);
        context.SetVariable("p", pass);

        var expression = navigator.Compile("//user[username=$u and password=$p]");
        expression.SetContext(context);

        var node = navigator.SelectSingleNode(expression);
        return node == null ? Unauthorized() : Ok("Authentifié");
    }

    private static XmlDocument LoadUsers()
    {
        var doc = new XmlDocument();
        doc.LoadXml("<users><user><username>alice</username><password>secret</password></user></users>");
        return doc;
    }
}

// XsltContext minimal permettant de lier des variables nommées à une
// expression XPath sans jamais les insérer dans le texte de l'expression.
internal sealed class BoundVariableContext : XsltContext
{
    private readonly Dictionary<string, string> _variables = new();

    public BoundVariableContext() : base(new NameTable()) { }

    public void SetVariable(string name, string value) => _variables[name] = value;

    public override IXsltContextVariable ResolveVariable(string prefix, string name) =>
        new BoundVariable(_variables[name]);

    public override int CompareDocument(string baseUri, string nextbaseUri) => 0;
    public override bool Whitespace => true;
    public override bool PreserveWhitespace(XPathNavigator node) => true;
    public override IXsltContextFunction ResolveFunction(string prefix, string name, XPathResultType[] argTypes) =>
        throw new NotSupportedException();

    private sealed class BoundVariable : IXsltContextVariable
    {
        private readonly string _value;
        public BoundVariable(string value) => _value = value;
        public bool IsLocal => false;
        public bool IsParam => false;
        public XPathResultType VariableType => XPathResultType.String;
        public object Evaluate(XsltContext context) => _value;
    }
}
