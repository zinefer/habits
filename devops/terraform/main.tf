data "azurerm_client_config" "current" {}

# Learn our public IP address
data "http" "icanhazip" {
   url = "http://icanhazip.com"
}

locals {
  my_ip = chomp(data.http.icanhazip.body)
  app_name = "habits-app-${random_pet.habits.id}"
  database_resource_group = var.database_resource_group == null ? azurerm_resource_group.habits.name : var.database_resource_group
}

# Create a resource group to hold all of our things
resource "azurerm_resource_group" "habits" {
  name     = local.app_name
  location = var.location

  lifecycle {
    prevent_destroy = true
  }
}

resource "random_pet" "habits" {
  length  = 1
}

resource "random_password" "database" {
  length = 16
  special = true
  override_special = "_%@"
}

resource "random_password" "database_user" {
  length = 16
  special = true
  override_special = "_%@"
}

resource "azurerm_postgresql_server" "habits" {
  name                = var.database_server
  location            = azurerm_resource_group.habits.location
  resource_group_name = local.database_resource_group

  sku {
    name     = "B_Gen5_1"
    capacity = 1
    tier     = "Basic"
    family   = "Gen5"
  }

  storage_profile {
    storage_mb            = 5120
    backup_retention_days = 7
    geo_redundant_backup  = "Disabled"
    auto_grow             = "Enabled"
  }

  administrator_login          = "postgres"
  administrator_login_password = random_password.database.result
  version                      = "10"
  ssl_enforcement              = "Enabled"
}

resource "azurerm_postgresql_firewall_rule" "allowAzure" {
  name                = "postgres-firewall-azure"
  resource_group_name = local.database_resource_group
  server_name         = azurerm_postgresql_server.habits.name
  start_ip_address    = "0.0.0.0"
  end_ip_address      = "0.0.0.0"
}

resource "azurerm_postgresql_firewall_rule" "allowMe" {
  name                = "postgres-firewall-me"
  resource_group_name = local.database_resource_group
  server_name         = azurerm_postgresql_server.habits.name
  start_ip_address    = local.my_ip
  end_ip_address      = local.my_ip
}

provider "postgresql" {
  host              = azurerm_postgresql_server.habits.fqdn
  port              = 5432
  username          = "postgres@${azurerm_postgresql_server.habits.name}"
  password          = random_password.database.result
  database_username = "postgres"
  superuser         = false
  sslmode           = "require"
  expected_version  = "10"
  connect_timeout   = 15
}

resource "postgresql_role" "habits" {
  login    = true
  name     = "habits"
  password = random_password.database_user.result
  depends_on = [azurerm_postgresql_firewall_rule.allowMe]
}

resource "postgresql_database" "habits" {
  name  = "habits_production"
  owner = postgresql_role.habits.name
  depends_on = [azurerm_postgresql_firewall_rule.allowMe]
}

resource "azurerm_app_service_plan" "habits" {
  name                = "habits-app-service"
  location            = azurerm_resource_group.habits.location
  resource_group_name = azurerm_resource_group.habits.name

  reserved = true
  kind = "Linux"

  sku {
    tier = "Basic"
    size = "B1"
  }
}

resource "azurerm_app_service" "habits" {
  name                = local.app_name
  location            = azurerm_resource_group.habits.location
  resource_group_name = azurerm_resource_group.habits.name
  app_service_plan_id = azurerm_app_service_plan.habits.id

  identity {
    type = "SystemAssigned"
  }

  logs {
    http_logs {
      file_system {
        retention_in_days = 7
        retention_in_mb   = 100
      }
    }
  }

  site_config {
    always_on        = true
    linux_fx_version = "DOCKER|zinefer/habits:latest"
  }

  app_settings = {
    HABITS_DATABASE_HOST     = azurerm_postgresql_server.habits.fqdn
    HABITS_DATABASE_USER     = "${postgresql_role.habits.name}@${azurerm_postgresql_server.habits.name}"
    HABITS_DATABASE_PASSWORD = random_password.database_user.result
    
    // OAUTH secrets
    HABITS_OAUTH_GITHUB_ID       = var.oauth_github_id
    HABITS_OAUTH_GITHUB_SECRET   = var.oauth_github_secret
    HABITS_OAUTH_FACEBOOK_ID     = var.oauth_facebook_id
    HABITS_OAUTH_FACEBOOK_SECRET = var.oauth_facebook_secret
    HABITS_OAUTH_GOOGLE_ID       = var.oauth_google_id
    HABITS_OAUTH_GOOGLE_SECRET   = var.oauth_google_secret

    HABITS_ACME_REDIRECT         = "https://${data.azurerm_storage_account.certs.primary_web_host}"
  }
}