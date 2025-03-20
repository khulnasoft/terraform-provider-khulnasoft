---
page_title: "khulnasoft_permissions_sets_saas Data Source - terraform-provider-khulnasoft"
subcategory: ""
description: |-
  The data source khulnasoft_permissions_sets_saas provides a method to query all permissions within Khulnasoft SaaS platform.
---

# khulnasoft_permissions_sets_saas (Data Source)

The data source `khulnasoft_permissions_sets_saas` provides a method to query all permissions within Khulnasoft SaaS platform.

## Example Usage

```terraform
data "khulnasoft_permissions_sets_saas" "testpermissionsset" {}

output "permissions_sets" {
  value = data.khulnasoft_permissions_sets_saas.testpermissionsset
}

output "permissions_sets_names" {
  value = data.khulnasoft_permissions_sets_saas.testpermissionsset[*].permissions_sets[*].name
}
```

## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `permissions_sets` (List of Object) (see [below for nested schema](#nestedatt--permissions_sets))

<a id="nestedatt--permissions_sets"></a>
### Nested Schema for `permissions_sets`

Read-Only:

- `actions` (List of String)
- `description` (String)
- `name` (String)