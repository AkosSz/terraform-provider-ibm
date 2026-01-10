// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package distributionlist

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/distributionlistv1"
	"github.com/go-openapi/strfmt"
)

func ResourceIbmDistributionListDestination() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmDistributionListDestinationCreate,
		ReadContext:   resourceIbmDistributionListDestinationRead,
		DeleteContext: resourceIbmDistributionListDestinationDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_distribution_list_destination", "account_id"),
				Description:  "The IBM Cloud account ID.",
			},
			"destination_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The GUID of the Event Notifications instance.",
			},
			"destination_type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_distribution_list_destination", "destination_type"),
				Description:  "The type of the destination.",
			},
			"email": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_distribution_list_destination", "email"),
				Description:  "The email address for the destination.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_distribution_list_destination", "name"),
				Description:  "The email name for the destination.",
			},
		},
	}
}

func ResourceIbmDistributionListDestinationValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "account_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-zA-Z]{1,100}$`,
			MinValueLength:             1,
			MaxValueLength:             100,
		},
		validate.ValidateSchema{
			Identifier:                 "destination_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "email, event_notifications",
		},
		validate.ValidateSchema{
			Identifier:                 "email",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
			MinValueLength:             3,
			MaxValueLength:             320,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.StringLenBetween,
			Type:                       validate.TypeString,
			Optional:                   true,
			MinValueLength:             3,
			MaxValueLength:             320,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_distribution_list_destination", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmDistributionListDestinationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	distributionListClient, err := meta.(conns.ClientSession).DistributionListV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bodyModelMap := map[string]interface{}{}
	createDistributionListDestinationOptions := &distributionlistv1.CreateDistributionListDestinationOptions{}

	if _, ok := d.GetOk("destination_id"); ok {
		bodyModelMap["destination_id"] = d.Get("destination_id")
	}
	bodyModelMap["destination_type"] = d.Get("destination_type")
	if _, ok := d.GetOk("name"); ok {
		bodyModelMap["name"] = d.Get("name")
	}
	if _, ok := d.GetOk("email"); ok {
		bodyModelMap["email"] = d.Get("email")
	}
	createDistributionListDestinationOptions.SetAccountID(d.Get("account_id").(string))
	convertedModel, err := ResourceIbmDistributionListDestinationMapToAddDestinationPrototype(bodyModelMap)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "create", "parse-request-body").GetDiag()
	}
	createDistributionListDestinationOptions.AddDestinationPrototype = convertedModel

	addDestinationIntf, _, err := distributionListClient.CreateDistributionListDestinationWithContext(context, createDistributionListDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDistributionListDestinationWithContext failed: %s", err.Error()), "ibm_distribution_list_destination", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if _, ok := addDestinationIntf.(*distributionlistv1.AddDestinationEmailDestination); ok {
		addDestination := addDestinationIntf.(*distributionlistv1.AddDestinationEmailDestination)
		d.SetId(fmt.Sprintf("%s/%s", *createDistributionListDestinationOptions.AccountID, *addDestination.DestinationID))
	} else if _, ok := addDestinationIntf.(*distributionlistv1.AddDestinationEventNotificationDestination); ok {
		addDestination := addDestinationIntf.(*distributionlistv1.AddDestinationEventNotificationDestination)
		d.SetId(fmt.Sprintf("%s/%s", *createDistributionListDestinationOptions.AccountID, *addDestination.DestinationID))
	} else if _, ok := addDestinationIntf.(*distributionlistv1.AddDestination); ok {
		addDestination := addDestinationIntf.(*distributionlistv1.AddDestination)
		d.SetId(fmt.Sprintf("%s/%s", *createDistributionListDestinationOptions.AccountID, *addDestination.DestinationID))
	} else {
		return flex.DiscriminatedTerraformErrorf(nil, fmt.Sprintf("Unrecognized distributionlistv1.AddDestinationIntf subtype encountered"), "ibm_distribution_list_destination", "create", "unrecognized-subtype-of-AddDestination").GetDiag()
	}

	return resourceIbmDistributionListDestinationRead(context, d, meta)
}

func resourceIbmDistributionListDestinationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	distributionListClient, err := meta.(conns.ClientSession).DistributionListV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDistributionListDestinationOptions := &distributionlistv1.GetDistributionListDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "sep-id-parts").GetDiag()
	}

	getDistributionListDestinationOptions.SetAccountID(parts[0])
	getDistributionListDestinationOptions.SetDestinationID(parts[1])

	addDestinationIntf, response, err := distributionListClient.GetDistributionListDestinationWithContext(context, getDistributionListDestinationOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDistributionListDestinationWithContext failed: %s", err.Error()), "ibm_distribution_list_destination", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if _, ok := addDestinationIntf.(*distributionlistv1.AddDestinationEmailDestination); ok {
		addDestination := addDestinationIntf.(*distributionlistv1.AddDestinationEmailDestination)
		if !core.IsNil(addDestination.DestinationID) {
			if err = d.Set("destination_id", addDestination.DestinationID); err != nil {
				err = fmt.Errorf("Error setting destination_id: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "set-destination_id").GetDiag()
			}
		}
		if err = d.Set("destination_type", addDestination.DestinationType); err != nil {
			err = fmt.Errorf("Error setting destination_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "set-destination_type").GetDiag()
		}
		if !core.IsNil(addDestination.Email) {
			if err = d.Set("email", addDestination.Email); err != nil {
				err = fmt.Errorf("Error setting email: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "set-email").GetDiag()
			}
		}
		if !core.IsNil(addDestination.Name) {
			if err = d.Set("name", addDestination.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "set-name").GetDiag()
			}
		}
	} else if _, ok := addDestinationIntf.(*distributionlistv1.AddDestinationEventNotificationDestination); ok {
		addDestination := addDestinationIntf.(*distributionlistv1.AddDestinationEventNotificationDestination)
		if !core.IsNil(addDestination.DestinationID) {
			if err = d.Set("destination_id", addDestination.DestinationID); err != nil {
				err = fmt.Errorf("Error setting destination_id: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "set-destination_id").GetDiag()
			}
		}
		if err = d.Set("destination_type", addDestination.DestinationType); err != nil {
			err = fmt.Errorf("Error setting destination_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "set-destination_type").GetDiag()
		}
	} else if _, ok := addDestinationIntf.(*distributionlistv1.AddDestination); ok {
		addDestination := addDestinationIntf.(*distributionlistv1.AddDestination)
		// parent class argument: account_id string
		if err = d.Set("account_id", getDistributionListDestinationOptions.AccountID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting account_id: %s", err))
		}
		// parent class argument: destination_id strfmt.UUID
		// parent class argument: destination_type string
		// parent class argument: email string
		// parent class argument: name string
		if !core.IsNil(addDestination.DestinationID) {
			if err = d.Set("destination_id", addDestination.DestinationID); err != nil {
				err = fmt.Errorf("Error setting destination_id: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "set-destination_id").GetDiag()
			}
		}
		if err = d.Set("destination_type", addDestination.DestinationType); err != nil {
			err = fmt.Errorf("Error setting destination_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "set-destination_type").GetDiag()
		}
		if !core.IsNil(addDestination.Email) {
			if err = d.Set("email", addDestination.Email); err != nil {
				err = fmt.Errorf("Error setting email: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "set-email").GetDiag()
			}
		}
		if !core.IsNil(addDestination.Name) {
			if err = d.Set("name", addDestination.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "read", "set-name").GetDiag()
			}
		}
	} else {
		return flex.DiscriminatedTerraformErrorf(nil, fmt.Sprintf("Unrecognized distributionlistv1.AddDestinationIntf subtype encountered"), "ibm_distribution_list_destination", "read", "unrecognized-subtype-of-AddDestination").GetDiag()
	}

	return nil
}

func resourceIbmDistributionListDestinationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	distributionListClient, err := meta.(conns.ClientSession).DistributionListV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteDistributionListDestinationOptions := &distributionlistv1.DeleteDistributionListDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_distribution_list_destination", "delete", "sep-id-parts").GetDiag()
	}

	deleteDistributionListDestinationOptions.SetAccountID(parts[0])
	deleteDistributionListDestinationOptions.SetDestinationID(parts[1])

	_, err = distributionListClient.DeleteDistributionListDestinationWithContext(context, deleteDistributionListDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDistributionListDestinationWithContext failed: %s", err.Error()), "ibm_distribution_list_destination", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmDistributionListDestinationMapToAddDestinationPrototype(modelMap map[string]interface{}) (distributionlistv1.AddDestinationPrototypeIntf, error) {
	discValue, ok := modelMap["destination_type"]
	if ok {
		if discValue == "event_notifications" {
			return ResourceIbmDistributionListDestinationMapToAddDestinationPrototypeEventNotificationDestinationPrototype(modelMap)
		} else if discValue == "email" {
			return ResourceIbmDistributionListDestinationMapToAddDestinationPrototypeEmailDestinationPrototype(modelMap)
		} else {
			return nil, fmt.Errorf("unexpected value for discriminator property 'destination_type' found in map: '%s'", discValue)
		}
	} else {
		return nil, fmt.Errorf("discriminator property 'destination_type' not found in map")
	}
}

func ResourceIbmDistributionListDestinationMapToAddDestinationPrototypeEmailDestinationPrototype(modelMap map[string]interface{}) (*distributionlistv1.AddDestinationPrototypeEmailDestinationPrototype, error) {
	model := &distributionlistv1.AddDestinationPrototypeEmailDestinationPrototype{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.Email = core.StringPtr(modelMap["email"].(string))
	model.DestinationType = core.StringPtr(modelMap["destination_type"].(string))
	return model, nil
}

func ResourceIbmDistributionListDestinationMapToAddDestinationPrototypeEventNotificationDestinationPrototype(modelMap map[string]interface{}) (*distributionlistv1.AddDestinationPrototypeEventNotificationDestinationPrototype, error) {
	model := &distributionlistv1.AddDestinationPrototypeEventNotificationDestinationPrototype{}
	model.DestinationID = core.UUIDPtr(strfmt.UUID(modelMap["destination_id"].(string)))
	model.DestinationType = core.StringPtr(modelMap["destination_type"].(string))
	return model, nil
}
