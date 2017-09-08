package hwcloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

// SKIP deprecated
func TestAccLBV1Member_importBasic(t *testing.T) {
	resourceName := "hwcloud_lb_member_v1.member_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckDeprecated(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLBV1MemberDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLBV1Member_basic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
