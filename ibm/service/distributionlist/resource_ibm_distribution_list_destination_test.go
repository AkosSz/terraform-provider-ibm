// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package distributionlist_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/distributionlist"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/distributionlistv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmDistributionListDestinationBasic(t *testing.T) {
	var conf distributionlistv1.AddDestination
	accountID := acc.DistributionListAccountId
	destinationId := acc.DistributionListDestinationId
	destinationType := "event_notifications"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmDistributionListDestinationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmDistributionListDestinationConfigBasic(accountID, destinationType, destinationId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmDistributionListDestinationExists("ibm_distribution_list_destination.distribution_list_destination_instance", conf),
					resource.TestCheckResourceAttr("ibm_distribution_list_destination.distribution_list_destination_instance", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_distribution_list_destination.distribution_list_destination_instance", "destination_type", destinationType),
					resource.TestCheckResourceAttr("ibm_distribution_list_destination.distribution_list_destination_instance", "destination_id", destinationId),
				),
			},
			{
				ResourceName:      "ibm_distribution_list_destination.distribution_list_destination_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmDistributionListDestinationConfigBasic(accountID string, destinationType string, destinationId string) string {
	return fmt.Sprintf(`
		resource "ibm_distribution_list_destination" "distribution_list_destination_instance" {
			account_id = "%s"
			destination_type = "%s"
			destination_id = "%s"
		}
	`, accountID, destinationType, destinationId)
}

func testAccCheckIbmDistributionListDestinationExists(n string, obj distributionlistv1.AddDestination) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		distributionListClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).DistributionListV1()
		if err != nil {
			return err
		}

		getDistributionListDestinationOptions := &distributionlistv1.GetDistributionListDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDistributionListDestinationOptions.SetAccountID(parts[0])
		getDistributionListDestinationOptions.SetDestinationID(parts[1])

		addDestinationIntf, _, err := distributionListClient.GetDistributionListDestination(getDistributionListDestinationOptions)
		if err != nil {
			return err
		}

		switch v := addDestinationIntf.(type) {
		case *distributionlistv1.AddDestinationEventNotificationDestination:
			obj = distributionlistv1.AddDestination{
				DestinationID:   v.DestinationID,
				DestinationType: v.DestinationType,
			}
		case *distributionlistv1.AddDestination:
			obj = *v
		default:
			return fmt.Errorf("unexpected destination type: %T", addDestinationIntf)
		}
		return nil
	}
}

func testAccCheckIbmDistributionListDestinationDestroy(s *terraform.State) error {
	distributionListClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).DistributionListV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_distribution_list_destination" {
			continue
		}

		getDistributionListDestinationOptions := &distributionlistv1.GetDistributionListDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDistributionListDestinationOptions.SetAccountID(parts[0])
		getDistributionListDestinationOptions.SetDestinationID(parts[1])

		// Try to find the key
		_, response, err := distributionListClient.GetDistributionListDestination(getDistributionListDestinationOptions)

		if err == nil {
			return fmt.Errorf("distribution_list_destination still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for distribution_list_destination (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmDistributionListDestinationMapToAddDestinationPrototype(t *testing.T) {
	// Checking the result is disabled for this model, because it has a discriminator
	// and there are separate tests for each child model below.
	model := make(map[string]interface{})
	model["destination_id"] = acc.DistributionListDestinationId
	model["destination_type"] = "event_notifications"

	_, err := distributionlist.ResourceIbmDistributionListDestinationMapToAddDestinationPrototype(model)
	assert.Nil(t, err)
}

func TestResourceIbmDistributionListDestinationMapToAddDestinationPrototypeEventNotificationDestinationPrototype(t *testing.T) {
	checkResult := func(result *distributionlistv1.AddDestinationPrototypeEventNotificationDestinationPrototype) {
		model := new(distributionlistv1.AddDestinationPrototypeEventNotificationDestinationPrototype)
		mock := strfmt.UUID("12345678-1234-1234-1234-123456789012")
		model.DestinationID = &mock
		model.DestinationType = core.StringPtr("event_notifications")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["destination_id"] = "12345678-1234-1234-1234-123456789012"
	model["destination_type"] = "event_notifications"

	result, err := distributionlist.ResourceIbmDistributionListDestinationMapToAddDestinationPrototypeEventNotificationDestinationPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}
