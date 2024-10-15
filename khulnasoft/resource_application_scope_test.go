package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftApplicationScope(t *testing.T) {
	t.Parallel()
	name := acctest.RandomWithPrefix("terraform-test")
	description := "Created using Terraform"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: CheckDestroy("khulnasoft_application_scope.terraformap"),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckApplicationScope(name, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationScopeExists("khulnasoft_application_scope.terraformap"),
				),
			},
			{
				ResourceName:      "khulnasoft_application_scope.terraformap",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckApplicationScope(name string, description string) string {
	return fmt.Sprintf(`
	resource "khulnasoft_application_scope" "terraformap" {
		description = "%s"
		name = "%s"
		categories {
			artifacts {
				image {
					expression = "v1 && v2 && v3"
					variables {
						attribute = "khulnasoft.registry"
						value = "test"
					}
					variables {
						attribute = "image.repo"
						value = "test123"
					}
					variables {
						attribute = "image.label"
						name = "test.label"
						value = "test.value.123"
					}
				}
			}
			workloads {
				kubernetes {
					expression = "v1 && v2"
					variables {
						attribute = "kubernetes.cluster"
						value = "test"
					}
					variables {
						attribute = "kubernetes.namespace"
						value = "test123"
					}
				}
			}
			infrastructure {
				kubernetes {
					expression = "v1"
					variables {
						attribute = "kubernetes.cluster"
						value = "lion"
					}
				}
			}
		}
	}`, description, name)

}

func testAccCheckApplicationScopeExists(n string) resource.TestCheckFunc {
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
