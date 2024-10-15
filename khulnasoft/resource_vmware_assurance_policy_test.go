package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftVMwareAssurancePolicy(t *testing.T) {
	t.Parallel()
	description := "Created using Terraform"
	name := acctest.RandomWithPrefix("terraform-test")
	application_scopes := "Global"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: CheckDestroy("khulnasoft_vmware_assurance_policy.terraformiap"),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckVMwareAssurancePolicy(description, name, application_scopes),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVMwareAssurancePolicyExists("khulnasoft_vmware_assurance_policy.terraformiap"),
				),
			},
			{
				ResourceName:      "khulnasoft_vmware_assurance_policy.terraformiap",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckVMwareAssurancePolicy(description string, name string, application_scopes string) string {
	return fmt.Sprintf(`
	resource "khulnasoft_vmware_assurance_policy" "terraformiap" {
		description = "%s"
		name = "%s"
		application_scopes = [
			"%s"
		]
	}`, description, name, application_scopes)

}

func testAccCheckVMwareAssurancePolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return NewNotFoundErrorf("%s in state", n)
		}

		if rs.Primary.ID == "" {
			return NewNotFoundErrorf("ID for %s in state", n)
		}

		return nil
	}
}
