package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestDataKhulnasoftFunctionAssurancePolicy(t *testing.T) {
	t.Parallel()
	name := "Default"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKhulnasoftFunctionAssurancePolicyDataSource(name),
				Check:  testAccCheckKhulnasoftFunctionAssurancePolicyDataSourceExists("data.khulnasoft_function_assurance_policy.defaultiap"),
			},
		},
	})
}

func testAccCheckKhulnasoftFunctionAssurancePolicyDataSource(name string) string {
	return fmt.Sprintf(`
	data "khulnasoft_function_assurance_policy" "defaultiap" {
		name = "%s"
	}
	`, name)

}

func testAccCheckKhulnasoftFunctionAssurancePolicyDataSourceExists(n string) resource.TestCheckFunc {
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
