data "khulnasoft_users" "users" {}

output "first_user_name" {
  value = data.khulnasoft_users.users.users[0].name // output: first_user_name = "administrator"
}