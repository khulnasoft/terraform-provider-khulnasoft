data "khulnasoft_function_runtime_policy" "function_runtime_policy" {
  name = "FunctionRuntimePolicyName"
}

output "function_runtime_policy_details" {
  value = data.khulnasoft_function_runtime_policy.function_runtime_policy
}