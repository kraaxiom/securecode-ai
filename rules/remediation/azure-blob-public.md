# Remédiation — Conteneur Azure Blob Storage public

## Principe
Désactiver l'accès public anonyme au niveau du compte de stockage et de chaque conteneur, restreindre l'accès réseau, et utiliser des SAS à durée de vie courte et permissions minimales pour tout partage ponctuel.

## Azure CLI
```bash
# Avant — vulnérable : accès blob public autorisé sur le compte et le conteneur
az storage account update \
  --name myappstorage --resource-group myrg \
  --allow-blob-public-access true

az storage container set-permission \
  --name myfiles --account-name myappstorage \
  --public-access container

# Après — sécurisé : accès public désactivé au niveau compte et conteneur
az storage account update \
  --name myappstorage --resource-group myrg \
  --allow-blob-public-access false

az storage container set-permission \
  --name myfiles --account-name myappstorage \
  --public-access off
```

## Terraform
```hcl
# Avant — vulnérable
resource "azurerm_storage_account" "myappstorage" {
  name                     = "myappstorage"
  resource_group_name      = azurerm_resource_group.rg.name
  location                 = azurerm_resource_group.rg.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
  allow_nested_items_to_be_public = true
}

resource "azurerm_storage_container" "myfiles" {
  name                  = "myfiles"
  storage_account_name  = azurerm_storage_account.myappstorage.name
  container_access_type = "container"
}

# Après — sécurisé
resource "azurerm_storage_account" "myappstorage" {
  name                     = "myappstorage"
  resource_group_name      = azurerm_resource_group.rg.name
  location                 = azurerm_resource_group.rg.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
  allow_nested_items_to_be_public = false

  network_rules {
    default_action = "Deny"
    ip_rules       = ["203.0.113.10"]
  }
}

resource "azurerm_storage_container" "myfiles" {
  name                  = "myfiles"
  storage_account_name  = azurerm_storage_account.myappstorage.name
  container_access_type = "private"
}
```

## Partage ponctuel sécurisé (SAS)
```bash
# Après — sécurisé : SAS en lecture seule, expirant dans 1 heure
az storage blob generate-sas \
  --account-name myappstorage --container-name myfiles --name rapport.pdf \
  --permissions r --expiry $(date -u -d "1 hour" '+%Y-%m-%dT%H:%MZ') \
  --https-only --output tsv
```

## Checklist de vérification post-patch
- [ ] `allowBlobPublicAccess` est à `false` au niveau du compte de stockage.
- [ ] Le niveau d'accès de chaque conteneur est réglé sur "Private" (`container_access_type = "private"`).
- [ ] Le firewall du compte de stockage restreint l'accès réseau (pas de "All networks" ouvert).
- [ ] Les SAS générés ont une expiration courte et des permissions scoped en lecture seule quand possible.
- [ ] Un accès public tenté sur l'URL du blob échoue (403 ResourceNotFound/PublicAccessNotPermitted).
- [ ] Azure Defender for Storage / Azure Policy ne remonte plus d'alerte sur ce compte.
