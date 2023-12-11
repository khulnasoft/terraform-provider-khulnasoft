package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftresourceRegistry(t *testing.T) {
	t.Parallel()
	name := acctest.RandomWithPrefix("terraform-test")
	url := "https://docker.io"
	rtype := "HUB"
	username := ""
	password := ""
	autopull := false
	scanner_type := "any"
	description := "Terrafrom-test"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: CheckDestroy("khulnasoft_integration_registry.new"),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKhulnasoftRegistry(name, url, rtype, username, password, autopull, scanner_type, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKhulnasoftRegistryExists("khulnasoft_integration_registry.new"),
				),
			},
			{
				ResourceName:            "khulnasoft_integration_registry.new",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"prefixes", "scanner_name", "last_updated"}, //TODO: implement read prefixes
			},
		},
	})
}

func testAccCheckKhulnasoftRegistry(name string, url string, rtype string, username string, password string, autopull bool, scanner_type string, description string) string {
	return fmt.Sprintf(`
	resource "khulnasoft_integration_registry" "new" {
		name = "%s"
		url = "%s"
		type = "%s"
		username = "%s"
		password = "%s"
		auto_pull = "%v"
		scanner_type = "%s"
		description = "%s"
	}`, name, url, rtype, username, password, autopull, scanner_type, description)

}

func testAccCheckKhulnasoftRegistryExists(n string) resource.TestCheckFunc {
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
