resource "khulnasoft_role" "IaC" {
  role_name   = "RoleIaC"
  description = "RoleIaC"
  permission  = "PermissionIaC"
  scopes      = ["Global"]
}
