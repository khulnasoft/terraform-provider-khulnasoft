resource "khulnasoft_role_mapping" "role_mapping" {
    saml {
        role_mapping = {
            Administrator = "group1"
        }
    }
}

output "role_mapping" {
    value = khulnasoft_role_mapping.role_mapping
}