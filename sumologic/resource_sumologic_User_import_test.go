// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Sumo Logic and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------\
package sumologic

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"strconv"
	"testing"
)

func TestSumologicUser_import(t *testing.T) {
	var user User
	testFirstName := FieldsMap["User"]["firstName"]
	testLastName := FieldsMap["User"]["lastName"]
	testEmail := FieldsMap["User"]["email"]
	testRoleIds := []string{"\"" + FieldsMap["User"]["roleIds"] + "\""}
	testIsActive, _ := strconv.ParseBool(FieldsMap["User"]["isActive"])

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { TestAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUserDestroy(user),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSumologicUserConfigImported(testFirstName, testLastName, testEmail, testRoleIds, testIsActive),
			},
			{
				ResourceName:      "sumologic_user.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSumologicUserConfigImported(firstName string, lastName string, email string, roleIds []string, isActive bool) string {
	return fmt.Sprintf(`
resource "sumologic_user" "foo" {
      first_name = "%s"
      last_name = "%s"
      email = "%s"
      role_ids = %v
      is_active = %t
}
`, firstName, lastName, email, roleIds, isActive)
}
