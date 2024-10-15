package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceKhulnasoftBasicFunctionRuntimePolicyCreate(t *testing.T) {
	t.Parallel()
	var runtimePolicy = client.RuntimePolicy{
		Name:        acctest.RandomWithPrefix("test-function-runtime-policy"),
		Description: "This is a test description of function runtime policy",
		Enabled:     true,
		Enforce:     true,
	}

	rootRef := functionRuntimePolicyRef("test")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: CheckDestroy("khulnasoft_function_runtime_policy.test"),
		Steps: []resource.TestStep{
			{
				Config: getFunctionRuntimePolicyResource(runtimePolicy),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rootRef, "name", runtimePolicy.Name),
					resource.TestCheckResourceAttr(rootRef, "description", runtimePolicy.Description),
					resource.TestCheckResourceAttr(rootRef, "application_scopes.#", "1"),
					resource.TestCheckResourceAttr(rootRef, "application_scopes.0", "Global"),
					resource.TestCheckResourceAttr(rootRef, "enabled", fmt.Sprintf("%v", runtimePolicy.Enabled)),
					resource.TestCheckResourceAttr(rootRef, "enforce", fmt.Sprintf("%v", runtimePolicy.Enforce)),
					//resource.TestCheckResourceAttr(rootRef, "author", os.Getenv("KHULNASOFT_USER")),
				),
			},
			{
				ResourceName:      "khulnasoft_function_runtime_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestResourceKhulnasoftFunctionRuntimePolicyUpgrade(t *testing.T) {
	t.Parallel()
	var runtimePolicy = client.RuntimePolicy{
		Name:        acctest.RandomWithPrefix("test-function-runtime-policy"),
		Description: "This is a test description of function runtime policy",
		Enabled:     true,
		Enforce:     true,
	}

	rootRef := functionRuntimePolicyRef("test")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: CheckDestroy("khulnasoft_function_runtime_policy.test"),
		Steps: []resource.TestStep{
			{
				Config: getFunctionRuntimePolicyResource(runtimePolicy),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rootRef, "name", runtimePolicy.Name),
					resource.TestCheckResourceAttr(rootRef, "description", runtimePolicy.Description),
					resource.TestCheckResourceAttr(rootRef, "application_scopes.#", "1"),
					resource.TestCheckResourceAttr(rootRef, "application_scopes.0", "Global"),
					resource.TestCheckResourceAttr(rootRef, "enabled", fmt.Sprintf("%v", runtimePolicy.Enabled)),
					resource.TestCheckResourceAttr(rootRef, "enforce", fmt.Sprintf("%v", runtimePolicy.Enforce)),
					//resource.TestCheckResourceAttr(rootRef, "author", os.Getenv("KHULNASOFT_USER")),
				),
			},
			{
				Config: getUpdatedFunctionRuntimePolicyResource(runtimePolicy),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rootRef, "name", runtimePolicy.Name),
					resource.TestCheckResourceAttr(rootRef, "description", runtimePolicy.Description),
					resource.TestCheckResourceAttr(rootRef, "application_scopes.#", "1"),
					resource.TestCheckResourceAttr(rootRef, "application_scopes.0", "Global"),
					resource.TestCheckResourceAttr(rootRef, "enabled", fmt.Sprintf("%v", runtimePolicy.Enabled)),
					resource.TestCheckResourceAttr(rootRef, "enforce", fmt.Sprintf("%v", runtimePolicy.Enforce)),
					//resource.TestCheckResourceAttr(rootRef, "author", os.Getenv("KHULNASOFT_USER")),
					resource.TestCheckResourceAttr(rootRef, "drift_prevention.0.enabled", "true"),
					resource.TestCheckResourceAttr(rootRef, "drift_prevention.0.exec_lockdown", "true"),
					resource.TestCheckResourceAttr(rootRef, "drift_prevention.0.image_lockdown", "false"),
					resource.TestCheckResourceAttr(rootRef, "drift_prevention.0.exec_lockdown_white_list.#", "1"),
					//todo: bring back after we upgrade the testing env
					//resource.TestCheckResourceAttr(rootRef, "block_malicious_executables_allowed_processes.#", "2"),
					resource.TestCheckResourceAttr(rootRef, "executable_blacklist.0.enabled", "true"),
					resource.TestCheckResourceAttr(rootRef, "executable_blacklist.0.executables.#", "2"),
				),
			},
		},
	})
}

func functionRuntimePolicyRef(name string) string {
	return fmt.Sprintf("khulnasoft_function_runtime_policy.%v", name)
}

func getFunctionRuntimePolicyResource(policy client.RuntimePolicy) string {
	return fmt.Sprintf(`
	resource "khulnasoft_function_runtime_policy" "test" {
		name = "%s"
		description = "%s"
		enabled = "%v"
		enforce = "%v"
	}
`, policy.Name, policy.Description, policy.Enabled, policy.Enforce)
}

func getUpdatedFunctionRuntimePolicyResource(policy client.RuntimePolicy) string {
	return fmt.Sprintf(`
	resource "khulnasoft_function_runtime_policy" "test" {
		name = "%s"
		description = "%s"
		enabled = "%v"
		enforce = "%v"
	drift_prevention {
		  enabled = true
		  exec_lockdown = true #block_running_executables_in_tmp_folder
		  image_lockdown = false
		  exec_lockdown_white_list = ["test"] #block_malicious_executables_allowed_processes
		}
	executable_blacklist {
		enabled = true
		executables = ["exe1","exe2"]
		}
	}
`,
		policy.Name,
		policy.Description,
		policy.Enabled,
		policy.Enforce)
}
