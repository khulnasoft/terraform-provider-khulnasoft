data "khulnasoft_container_runtime_policy" "container_runtime_policy" {
  name = "FunctionRuntimePolicyName"
}

output "container_runtime_policy_details" {
  value = data.khulnasoft_container_runtime_policy.container_runtime_policy
}