package khulnasoft

import (
	"fmt"
	"os"
	"testing"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var imageData = client.Image{
	Registry:   acctest.RandomWithPrefix("terraform-test"),
	Repository: "alpine",
	Tag:        "3.13",
}

func TestDataSourceKhulnasoftImage(t *testing.T) {
	t.Parallel()
	rootRef := imageDataRef("test")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: getImageDataSource(&imageData),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rootRef, "registry", imageData.Registry),
					resource.TestCheckResourceAttr(rootRef, "registry_type", "HUB"),
					resource.TestCheckResourceAttr(rootRef, "repository", imageData.Repository),
					resource.TestCheckResourceAttr(rootRef, "tag", imageData.Tag),
					resource.TestCheckResourceAttrSet(rootRef, "scan_status"),
					resource.TestCheckResourceAttrSet(rootRef, "disallowed"),
					resource.TestCheckResourceAttrSet(rootRef, "scan_date"),
					resource.TestCheckResourceAttr(rootRef, "scan_error", ""),
					resource.TestCheckResourceAttrSet(rootRef, "critical_vulnerabilities"),
					resource.TestCheckResourceAttrSet(rootRef, "high_vulnerabilities"),
					resource.TestCheckResourceAttrSet(rootRef, "medium_vulnerabilities"),
					resource.TestCheckResourceAttrSet(rootRef, "low_vulnerabilities"),
					resource.TestCheckResourceAttrSet(rootRef, "negligible_vulnerabilities"),
					resource.TestCheckResourceAttrSet(rootRef, "total_vulnerabilities"),
					resource.TestCheckResourceAttr(rootRef, "author", os.Getenv("KHULNASOFT_USER")),
					resource.TestCheckResourceAttrSet(rootRef, "image_size"),
				),
			},
		},
	})
}

func imageDataRef(name string) string {
	return fmt.Sprintf("data.khulnasoft_image.%s", name)
}

func getImageDataSource(image *client.Image) string {
	return getRegistry(image.Registry) + fmt.Sprintf(`
	resource "khulnasoft_image" "test" {
		registry = khulnasoft_integration_registry.demo.id
		repository = "%s"
		tag = "%s"
	}

	data "khulnasoft_image" "test" {
		registry = split("/", khulnasoft_image.test.id).0
		repository = split(":", split("/", khulnasoft_image.test.id).1).0
		tag = split(":", split("/", khulnasoft_image.test.id).1).1
	}
`, image.Repository, image.Tag)
}
