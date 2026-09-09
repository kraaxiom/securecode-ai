# Remédiation — Sécurité des interfaces gRPC

## Principe
Appliquer un intercepteur d'authentification/autorisation global à toutes les méthodes RPC (liste blanche explicite pour les méthodes publiques), désactiver la réflexion en production, et chiffrer systématiquement les canaux avec TLS.

## Go
```go
// Avant — vulnérable
lis, _ := net.Listen("tcp", ":50051")
s := grpc.NewServer() // aucun intercepteur, aucun TLS
reflection.Register(s) // réflexion active sans garde
pb.RegisterOrderServiceServer(s, &server{})
s.Serve(lis)

// Après — sécurisé
creds, _ := credentials.NewServerTLSFromFile("server.crt", "server.key")
s := grpc.NewServer(
    grpc.Creds(creds),
    grpc.UnaryInterceptor(authInterceptor),
)
if os.Getenv("APP_ENV") != "production" {
    reflection.Register(s)
}
pb.RegisterOrderServiceServer(s, &server{})
s.Serve(lis)

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
    publicMethods := map[string]bool{"/health.Check": true}
    if !publicMethods[info.FullMethod] {
        if err := verifyToken(ctx); err != nil {
            return nil, status.Errorf(codes.Unauthenticated, "non autorisé")
        }
    }
    return handler(ctx, req)
}
```

## Java (gRPC-Java)
```java
// Avant — vulnérable
Server server = ServerBuilder.forPort(50051)
    .addService(new OrderServiceImpl())
    .build(); // pas d'intercepteur, pas de TLS

// Après — sécurisé
Server server = ServerBuilder.forPort(50051)
    .addService(ServerInterceptors.intercept(new OrderServiceImpl(), new AuthInterceptor()))
    .useTransportSecurity(new File("server.crt"), new File("server.key"))
    .build();
```

## Python (grpcio)
```python
# Avant — vulnérable
server = grpc.server(futures.ThreadPoolExecutor())
add_OrderServiceServicer_to_server(OrderService(), server)
server.add_insecure_port('[::]:50051')

# Après — sécurisé
server = grpc.server(
    futures.ThreadPoolExecutor(),
    interceptors=[AuthInterceptor()],
)
add_OrderServiceServicer_to_server(OrderService(), server)
with open('server.key', 'rb') as f, open('server.crt', 'rb') as c:
    creds = grpc.ssl_server_credentials([(f.read(), c.read())])
server.add_secure_port('[::]:50051', creds)
```

## Checklist de vérification post-patch
- [ ] Un intercepteur global d'authentification/autorisation couvre toutes les méthodes RPC, avec liste blanche explicite pour les méthodes publiques.
- [ ] Le service de réflexion gRPC est désactivé (ou conditionné à un environnement non-production).
- [ ] Toutes les communications gRPC utilisent TLS, y compris interservices.
- [ ] Le fichier `.proto` documente le niveau d'autorisation attendu par méthode.
- [ ] Un test confirme qu'un appel RPC sans jeton valide est rejeté avec `UNAUTHENTICATED`.
