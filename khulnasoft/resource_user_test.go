package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftUserManagement(t *testing.T) {

	if isSaasEnv() {
		t.Skip("Skipping user test because its saas env")
	}

	t.Parallel()
	userID := acctest.RandomWithPrefix("terraform-test-user")
	password := "Pas5wo-d"
	name := "terraform"
	email := "terraform@test.com"
	newEmail := "terraform1@test.com"
	role := "Administrator"

	if isSaasEnv() {
		t.Skip("Skipping user test because its saas env")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccUserDestroy,
		Steps: []resource.TestStep{
			{
				// Config returns the test resource
				Config: testAccCheckKhulnasoftUser(userID, password, name, email, role),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKhulnasoftUsersExists("khulnasoft_user.new"),
				),
			},
			{
				// Config returns the test resource
				Config: testAccCheckKhulnasoftUser(userID, password, name, newEmail, role),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKhulnasoftUsersExists("khulnasoft_user.new"),
				),
			},
			{
				ResourceName:      "khulnasoft_user.new",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckKhulnasoftUser(userID string, password string, name string, email string, role string) string {
	return fmt.Sprintf(`
	resource "khulnasoft_user" "new" {
		user_id  = "%s"
		password = "%s"
		name     = "%s"
		email    = "%s"
		roles = [
		  "%s"
		]
	  }`, userID, password, name, email, role)
}

func testAccCheckKhulnasoftUsersExists(n string) resource.TestCheckFunc {
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

func testAccUserDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "khulnasoft_user.new" {
			continue
		}

		if rs.Primary.ID != "" {
			return fmt.Errorf("Object %q still exists", rs.Primary.ID)
		}
		return nil
	}
	return nil
}
