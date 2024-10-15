package khulnasoft

import (
	"fmt"
	"os"
	"testing"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func newTestImage() client.Image {
	return client.Image{
		Registry:   acctest.RandomWithPrefix("terraform-test"),
		Repository: "alpine",
		Tag:        "3.4",
	}
}

func TestResourceKhulnasoftImageCreate(t *testing.T) {
	//t.Parallel()
	image := newTestImage()
	rootRef := imageResourceRef("test")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: CheckDestroy("khulnasoft_image.test"),
		Steps: []resource.TestStep{
			{
				Config: getImageResource(&image),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rootRef, "registry", image.Registry),
					resource.TestCheckResourceAttr(rootRef, "registry_type", "HUB"),
					resource.TestCheckResourceAttr(rootRef, "repository", image.Repository),
					resource.TestCheckResourceAttr(rootRef, "tag", image.Tag),
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
			{
				ResourceName:            "khulnasoft_image.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"scan_date", "scan_status", "scan_error", "vulnerabilities"},
			},
		},
	})
}

func TestResourceKhulnasoftImageAllow(t *testing.T) {
	//t.Parallel()
	image := newTestImage()
	rootRef := imageResourceRef("test")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: CheckDestroy("khulnasoft_image.test"),
		Steps: []resource.TestStep{
			{
				Config: getImageResource(&image),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rootRef, "whitelisted", "false"),
					resource.TestCheckResourceAttr(rootRef, "blacklisted", "false"),
					resource.TestCheckResourceAttr(rootRef, "permission_comment", ""),
				),
			},
			{
				Config: getImageResourceAllow(&image, "This image is whitelisted from terraform test."),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rootRef, "whitelisted", "true"),
					resource.TestCheckResourceAttr(rootRef, "blacklisted", "false"),
					resource.TestCheckResourceAttr(rootRef, "permission_comment", "This image is whitelisted from terraform test."),
				),
			},
		},
	})
}

func TestResourceKhulnasoftImageBlock(t *testing.T) {
	//t.Parallel()
	image := newTestImage()
	rootRef := imageResourceRef("test")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: CheckDestroy("khulnasoft_image.test"),
		Steps: []resource.TestStep{
			{
				Config: getImageResource(&image),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rootRef, "whitelisted", "false"),
					resource.TestCheckResourceAttr(rootRef, "blacklisted", "false"),
					resource.TestCheckResourceAttr(rootRef, "permission_comment", ""),
				),
			},
			{
				Config: getImageResourceBlock(&image, "This image is blacklisted from terraform test."),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rootRef, "blacklisted", "true"),
					resource.TestCheckResourceAttr(rootRef, "whitelisted", "false"),
					resource.TestCheckResourceAttr(rootRef, "permission_comment", "This image is blacklisted from terraform test."),
				),
			},
		},
	})
}

func TestResourceKhulnasoftImageAllowAndBlock(t *testing.T) {
	//t.Parallel()
	image := newTestImage()
	rootRef := imageResourceRef("test")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: CheckDestroy("khulnasoft_image.test"),
		Steps: []resource.TestStep{
			{
				Config: getImageResource(&image),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rootRef, "whitelisted", "false"),
					resource.TestCheckResourceAttr(rootRef, "blacklisted", "false"),
					resource.TestCheckResourceAttr(rootRef, "permission_comment", ""),
				),
			},
			{
				Config: getImageResourceAllow(&image, "This image is whitelisted from terraform test."),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rootRef, "whitelisted", "true"),
					resource.TestCheckResourceAttr(rootRef, "blacklisted", "false"),
					resource.TestCheckResourceAttr(rootRef, "permission_comment", "This image is whitelisted from terraform test."),
				),
			},
			{
				Config: getImageResourceBlock(&image, "This image is blacklisted from terraform test."),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rootRef, "blacklisted", "true"),
					resource.TestCheckResourceAttr(rootRef, "whitelisted", "false"),
					resource.TestCheckResourceAttr(rootRef, "permission_comment", "This image is blacklisted from terraform test."),
				),
			},
		},
	})
}

func imageResourceRef(name string) string {
	return fmt.Sprintf("khulnasoft_image.%s", name)
}

func getImageResource(image *client.Image) string {
	return getRegistry(image.Registry) + fmt.Sprintf(`
	resource "khulnasoft_image" "test" {
		registry = khulnasoft_integration_registry.demo.id
		repository = "%s"
		tag = "%s"
	}
`, image.Repository, image.Tag)
}

func getImageResourceAllow(image *client.Image, comment string) string {
	return getRegistry(image.Registry) + fmt.Sprintf(`
	resource "khulnasoft_image" "test" {
		registry = khulnasoft_integration_registry.demo.id
		repository = "%s"
		tag = "%s"
		allow_image = true
		permission_modification_comment = "%s"
	}
`, image.Repository, image.Tag, comment)
}

func getImageResourceBlock(image *client.Image, comment string) string {
	return getRegistry(image.Registry) + fmt.Sprintf(`
	resource "khulnasoft_image" "test" {
		registry = khulnasoft_integration_registry.demo.id
		repository = "%s"
		tag = "%s"
		block_image = true
		permission_modification_comment = "%s"
	}
`, image.Repository, image.Tag, comment)
}

func getRegistry(name string) string {
	return fmt.Sprintf(`
	resource "khulnasoft_integration_registry" "demo" {
		name = "%s"
		type = "HUB"
		scanner_type = "any"
	}
`, name)
}
