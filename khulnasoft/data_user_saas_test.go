package khulnasoft

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftUserSaasManagementDatasource(t *testing.T) {

	if !isSaasEnv() {
		t.Skip("Skipping saas user test because its on prem env")
	}
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKhulnasoftUserSaasDataSource(),
				Check:  testAccCheckKhulnasoftUsersSaasDataSourceExists("data.khulnasoft_users_saas.testusers"),
			},
		},
	})
}

func testAccCheckKhulnasoftUserSaasDataSource() string {
	return `
	data "khulnasoft_users_saas" "testusers" {}
	`
}

func testAccCheckKhulnasoftUsersSaasDataSourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return NewNotFoundErrorf("%s in state", n)
		}

		if rs.Primary.ID == "" {
			return NewNotFoundErrorf("Id for %s in state", n)
		}

		return nil
	}
}
