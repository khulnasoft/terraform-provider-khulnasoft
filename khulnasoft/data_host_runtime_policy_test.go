package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestDataKhulnasoftBasicHostRuntimePolicy(t *testing.T) {
	t.Parallel()
	var basicRuntimePolicy = client.RuntimePolicy{
		Name:             acctest.RandomWithPrefix("test-host-runtime-policy"),
		Description:      "This is a test description of host runtime policy",
		Enabled:          false,
		Enforce:          false,
		EnforceAfterDays: 5,
	}

	rootRef := dataHostRuntimePolicyRef("test")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: getBasicHostRuntimePolicyData(basicRuntimePolicy),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rootRef, "name", basicRuntimePolicy.Name),
					resource.TestCheckResourceAttr(rootRef, "description", basicRuntimePolicy.Description),
					resource.TestCheckResourceAttr(rootRef, "application_scopes.#", "1"),
					resource.TestCheckResourceAttr(rootRef, "application_scopes.0", "Global"),
					resource.TestCheckResourceAttr(rootRef, "enabled", fmt.Sprintf("%v", basicRuntimePolicy.Enabled)),
					resource.TestCheckResourceAttr(rootRef, "enforce", fmt.Sprintf("%v", basicRuntimePolicy.Enforce)),
					resource.TestCheckResourceAttr(rootRef, "enforce_after_days", fmt.Sprintf("%v", basicRuntimePolicy.EnforceAfterDays)),
					//resource.TestCheckResourceAttr(rootRef, "author", os.Getenv("KHULNASOFT_USER")),
				),
			},
		},
	})
}

func TestDataKhulnasoftComplexHostRuntimePolicy(t *testing.T) {
	t.Parallel()
	var complexRuntimePolicy = client.RuntimePolicy{
		Name:        acctest.RandomWithPrefix("test-host-runtime-policy"),
		Description: "This is a test description of host runtime policy",
		Enabled:     true,
		Enforce:     true,
	}

	rootRef := dataHostRuntimePolicyRef("test")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: getComplexHostRuntimePolicyData(complexRuntimePolicy),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rootRef, "name", complexRuntimePolicy.Name),
					resource.TestCheckResourceAttr(rootRef, "description", complexRuntimePolicy.Description),
					resource.TestCheckResourceAttr(rootRef, "application_scopes.#", "1"),
					resource.TestCheckResourceAttr(rootRef, "application_scopes.0", "Global"),
					resource.TestCheckResourceAttr(rootRef, "enabled", fmt.Sprintf("%v", complexRuntimePolicy.Enabled)),
					resource.TestCheckResourceAttr(rootRef, "enforce", fmt.Sprintf("%v", complexRuntimePolicy.Enforce)),
					resource.TestCheckResourceAttr(rootRef, "enforce_after_days", fmt.Sprintf("%v", complexRuntimePolicy.EnforceAfterDays)),
					//resource.TestCheckResourceAttr(rootRef, "author", os.Getenv("KHULNASOFT_USER")),
					//todo: bring back after we upgrade the testing env
					//resource.TestCheckResourceAttr(rootRef, "block_cryptocurrency_mining", "true"),
					//resource.TestCheckResourceAttr(rootRef, "audit_brute_force_login", "true"),
					resource.TestCheckResourceAttr(rootRef, "enable_ip_reputation", "true"),
					resource.TestCheckResourceAttr(rootRef, "file_integrity_monitoring.0.enabled", "true"),
					resource.TestCheckResourceAttr(rootRef, "file_integrity_monitoring.0.monitored_files_create", "true"),
					resource.TestCheckResourceAttr(rootRef, "file_integrity_monitoring.0.monitored_files_read", "true"),
					resource.TestCheckResourceAttr(rootRef, "file_integrity_monitoring.0.monitored_files_modify", "true"),
					resource.TestCheckResourceAttr(rootRef, "file_integrity_monitoring.0.monitored_files_delete", "true"),
					resource.TestCheckResourceAttr(rootRef, "file_integrity_monitoring.0.monitored_files_attributes", "true"),
					resource.TestCheckResourceAttr(rootRef, "file_integrity_monitoring.0.monitored_files.#", "1"),
					resource.TestCheckResourceAttr(rootRef, "file_integrity_monitoring.0.exceptional_monitored_files.#", "1"),
					resource.TestCheckResourceAttr(rootRef, "file_integrity_monitoring.0.monitored_files_processes.#", "1"),
					resource.TestCheckResourceAttr(rootRef, "file_integrity_monitoring.0.exceptional_monitored_files_processes.#", "1"),
					resource.TestCheckResourceAttr(rootRef, "file_integrity_monitoring.0.monitored_files_users.#", "1"),
					resource.TestCheckResourceAttr(rootRef, "file_integrity_monitoring.0.exceptional_monitored_files_users.#", "1"),
					resource.TestCheckResourceAttr(rootRef, "auditing.0.audit_os_user_activity", "true"),
					//resource.TestCheckResourceAttr(rootRef, "auditing.0.audit_full_command_arguments", "true"),
					//todo: bring back after we upgrade the testing env
					//resource.TestCheckResourceAttr(rootRef, "audit_host_successful_login_events", "true"),
					//resource.TestCheckResourceAttr(rootRef, "audit_host_failed_login_events", "true"),
					resource.TestCheckResourceAttr(rootRef, "auditing.0.audit_success_login", "true"),
					//resource.TestCheckResourceAttr(rootRef, "port_scanning_detection", "true"),
					//resource.TestCheckResourceAttr(rootRef, "monitor_system_time_changes", "true"),
					//resource.TestCheckResourceAttr(rootRef, "monitor_windows_services", "true"),
					//resource.TestCheckResourceAttr(rootRef, "monitor_system_log_integrity", "true"),
				),
			},
		},
	})
}

func dataHostRuntimePolicyRef(name string) string {
	return fmt.Sprintf("data.khulnasoft_host_runtime_policy.%v", name)
}

func getBasicHostRuntimePolicyData(policy client.RuntimePolicy) string {
	return fmt.Sprintf(`
	resource "khulnasoft_host_runtime_policy" "test" {
		name = "%s"
		description = "%s"
		enabled = "%v"
		enforce = "%v"
		enforce_after_days = "%d"
	}
	
	data "khulnasoft_host_runtime_policy" "test" {
		name = khulnasoft_host_runtime_policy.test.id
	}
`, policy.Name, policy.Description, policy.Enabled, policy.Enforce, policy.EnforceAfterDays)
}

func getComplexHostRuntimePolicyData(policy client.RuntimePolicy) string {
	return fmt.Sprintf(`
	resource "khulnasoft_host_runtime_policy" "test" {
		name = "%s"
		description = "%s"
		enabled = "%v"
		enforce = "%v"
		# block_cryptocurrency_mining = true
		# audit_brute_force_login = true
	file_integrity_monitoring {
		enabled                                = true
		monitored_files_create                 = true
		monitored_files_read                   = true
		monitored_files_modify                 = true
		monitored_files_delete                 = true
		monitored_files_attributes             = true
		monitored_files                        = ["paths"]
		exceptional_monitored_files            = ["expaths"]
		monitored_files_processes              = ["process"]
		exceptional_monitored_files_processes  = ["exprocess"]
		monitored_files_users                  = ["user"]
		exceptional_monitored_files_users      = ["expuser"]
	  }
	  auditing {
		audit_os_user_activity        = true
		audit_user_account_management = true
		audit_success_login = true
	  }
	  enable_ip_reputation = true
	  enable_port_scan_protection     = true
}
	data "khulnasoft_host_runtime_policy" "test" {
		name = khulnasoft_host_runtime_policy.test.id
	}

`,
		policy.Name,
		policy.Description,
		policy.Enabled,
		policy.Enforce)
}
