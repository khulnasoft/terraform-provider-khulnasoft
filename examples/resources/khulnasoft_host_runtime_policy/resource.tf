resource "khulnasoft_host_runtime_policy" "host_runtime_policy" {
  name        = "host_runtime_policy"
  description = "host_runtime_policy"
  scope_variables {
    attribute = "kubernetes.cluster"
    value     = "default"
  }
  scope_variables {
    attribute = "kubernetes.label"
    name      = "app"
    value     = "khulnasoft"
  }

  application_scopes = [
    "Global",
  ]
  enabled                     = true
  enforce                     = false
  block_cryptocurrency_mining = true
  audit_brute_force_login     = true
  blocked_files = [
    "blocked",
  ]
  file_integrity_monitoring {
    monitor_create      = true
    monitor_read        = true
    monitor_modify      = true
    monitor_delete      = true
    monitor_attributes  = true
    monitored_paths     = ["paths"]
    excluded_paths      = ["expaths"]
    monitored_processes = ["process"]
    excluded_processes  = ["exprocess"]
    monitored_users     = ["user"]
    excluded_users      = ["expuser"]
  }
  audit_all_os_user_activity         = true
  audit_full_command_arguments       = true
  audit_host_successful_login_events = true
  audit_host_failed_login_events     = true
  audit_user_account_management      = true
  os_users_allowed = [
    "user1",
  ]
  os_groups_allowed = [
    "group1",
  ]
  os_users_blocked = [
    "user2",
  ]
  os_groups_blocked = [
    "group2",
  ]
  package_block = [
    "package1"
  ]
  monitor_system_time_changes  = true
  monitor_windows_services     = true
  monitor_system_log_integrity = true
}