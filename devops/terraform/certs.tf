locals {
  clean_lets_encrypt_azure = replace(var.lets_encrypt_azure, "-", "")
  clean_app_name = replace(local.app_name, "-", "")
}

data "azuread_service_principal" "app" {
  display_name = "Microsoft Azure App Service"
}

data "azuread_service_principal" "certs" {
  display_name = var.lets_encrypt_azure
}

data "azurerm_app_service" "certs" {
  name                = var.lets_encrypt_azure
  resource_group_name = var.lets_encrypt_azure
}

data "azurerm_storage_account" "certs_config" {
  name                = local.clean_lets_encrypt_azure
  resource_group_name = var.lets_encrypt_azure
}

resource "azurerm_storage_account" "certs" {
  name                     = local.clean_app_name
  location                 = azurerm_resource_group.habits.location
  resource_group_name      = azurerm_resource_group.habits.name
  account_tier             = "Standard"
  access_tier              = "Hot"
  account_kind             = "StorageV2"
  account_replication_type = "LRS"

  enable_https_traffic_only = true

  # Currently terraform cannot enable static-website but it is coming
  # https://github.com/terraform-providers/terraform-provider-azurerm/issues/1903
  provisioner "local-exec" {
    command = "az storage blob service-properties update --account-name ${azurerm_storage_account.certs.name} --static-website  --index-document index.html --404-document 404.html"
  }
}

data "azurerm_storage_account" "certs" {
  name                = azurerm_storage_account.certs.name
  resource_group_name = azurerm_storage_account.certs.resource_group_name
  depends_on = [azurerm_storage_account.certs]
}

resource "azurerm_key_vault" "certs" {
  name                = local.app_name
  location            = azurerm_resource_group.habits.location
  resource_group_name = azurerm_resource_group.habits.name
  tenant_id           = data.azurerm_client_config.current.tenant_id

  sku_name = "standard"

  access_policy {
    tenant_id = data.azurerm_client_config.current.tenant_id
    object_id = data.azurerm_client_config.current.object_id

    key_permissions = [
      "list"
    ]

    secret_permissions = [
      "list"
    ]

    certificate_permissions = [
      "list"
    ]
  }

  /*network_acls {
    bypass         = "AzureServices"
    default_action = "Deny"
    ip_rules       = [local.my_ip]
  }*/
}

resource "azurerm_key_vault_access_policy" "certs" {
  key_vault_id = azurerm_key_vault.certs.id

  tenant_id = data.azurerm_client_config.current.tenant_id
  object_id = data.azuread_service_principal.certs.object_id

  secret_permissions = [
    "get", "list"
  ]

  certificate_permissions = [
    "get", "list", "import", "update"
  ]
} 

resource "azurerm_key_vault_access_policy" "habits" {
  key_vault_id = azurerm_key_vault.certs.id

  tenant_id = azurerm_app_service.habits.identity[0].tenant_id
  object_id = azurerm_app_service.habits.identity[0].principal_id

  secret_permissions = [
    "get",
  ]
}

resource "azurerm_key_vault_access_policy" "app" {
  key_vault_id = azurerm_key_vault.certs.id

  tenant_id = data.azurerm_client_config.current.tenant_id
  object_id = data.azuread_service_principal.app.object_id

  certificate_permissions = [
    "get"
  ]

  secret_permissions = [
    "get",
  ]
}

resource "azurerm_role_assignment" "certsKeyvaultContributor" {
  scope                = azurerm_key_vault.certs.id
  role_definition_name = "Key Vault Contributor"
  principal_id         = data.azuread_service_principal.certs.object_id
}

resource "azurerm_role_assignment" "certsWebPlanContributor" {
  scope                = azurerm_resource_group.habits.id
  role_definition_name = "Web Plan Contributor"
  principal_id         = data.azuread_service_principal.certs.object_id
}

resource "azurerm_role_assignment" "certsResourceGroupReader" {
  scope                = azurerm_resource_group.habits.id
  role_definition_name = "Reader"
  principal_id         = data.azuread_service_principal.certs.object_id
}

resource "azurerm_role_assignment" "certsWebsiteContributor" {
  scope                = azurerm_resource_group.habits.id
  role_definition_name = "Website Contributor"
  principal_id         = data.azuread_service_principal.certs.object_id
}

resource "azurerm_role_assignment" "certsStorageContributor" {
  scope                = azurerm_storage_account.certs.id
  role_definition_name = "Storage Blob Data Contributor"
  principal_id         = data.azuread_service_principal.certs.object_id
}

resource "azurerm_app_service_custom_hostname_binding" "wwwHabits" {
  hostname            = "www.habits.watch"
  app_service_name    = azurerm_app_service.habits.name
  resource_group_name = azurerm_resource_group.habits.name
}

resource "azurerm_app_service_custom_hostname_binding" "habits" {
  hostname            = "habits.watch"
  app_service_name    = azurerm_app_service.habits.name
  resource_group_name = azurerm_resource_group.habits.name
}

resource "azurerm_storage_blob" "certs" {
  name                   = "config/habits.json"
  storage_account_name   = data.azurerm_storage_account.certs_config.name
  storage_container_name = "letsencrypt"
  type                   = "Block"
  source_content         = <<EOF
    {
        "acme": {
            "email": "zinefer@gmail.com",
            "renewXDaysBeforeExpiry": 30,
            "staging": false
        },
        "certificates": [
            {
                "hostNames": [
                    "habits.watch",
                    "www.habits.watch"
                ],
                "targetResource": {
                    "type": "appService",
                    "name": "${local.app_name}"
                }
            }
        ]
    }
  EOF
}