package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftNotificationOld(t *testing.T) {
	t.Parallel()
	user_name := "Khulnasoft"
	channel := "#general"
	webhook_url := "terraform-eg"
	enabled := true
	stype := "slack"
	name := "Slack"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: CheckDestroy("khulnasoft_notification_slack.slacknew"),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckNotificationOld(user_name, channel, webhook_url, enabled, stype, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationOldExists("khulnasoft_notification_slack.slacknew"),
				),
			},
			{
				ResourceName:      "khulnasoft_notification_slack.slacknew",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckNotificationOld(user_name string, channel string, webhook_url string, enabled bool, stype string, name string) string {
	return fmt.Sprintf(`
	resource "khulnasoft_notification_slack" "slacknew" {
		user_name = "%s"
		channel = "%s"
		webhook_url = "%s"
		enabled = "%v"
		type = "%s"
		name = "%s"
	  }`, user_name, channel, webhook_url, enabled, stype, name)

}

func testAccCheckNotificationOldExists(n string) resource.TestCheckFunc {
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
