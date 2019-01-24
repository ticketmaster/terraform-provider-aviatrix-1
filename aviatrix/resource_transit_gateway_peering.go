package aviatrix

import (
	"fmt"
	"log"

	"github.com/AviatrixSystems/go-aviatrix/goaviatrix"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTransitGatewayPeering() *schema.Resource {
	return &schema.Resource{
		Create: resourceTransitGatewayPeeringCreate,
		Read:   resourceTransitGatewayPeeringRead,
		Update: resourceTransitGatewayPeeringUpdate,
		Delete: resourceTransitGatewayPeeringDelete,

		Schema: map[string]*schema.Schema{
			"transit_gateway_name1": {
				Type:     schema.TypeString,
				Required: true,
			},
			"transit_gateway_name2": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceTransitGatewayPeeringCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goaviatrix.Client)
	transitGatewayPeering := &goaviatrix.TransitGatewayPeering{
		TransitGatewayName1: d.Get("transit_gateway_name1").(string),
		TransitGatewayName2: d.Get("transit_gateway_name2").(string),
	}

	log.Printf("[INFO] Creating Aviatrix Transit Gateway peering: %#v", transitGatewayPeering)

	err := client.CreateTransitGatewayPeering(transitGatewayPeering)
	if err != nil {
		return fmt.Errorf("failed to create Aviatrix Transit Gateway peering: %s", err)
	}
	d.SetId(transitGatewayPeering.TransitGatewayName1 + "<->" + transitGatewayPeering.TransitGatewayName2)
	return resourceTransitGatewayPeeringRead(d, meta)
}

func resourceTransitGatewayPeeringRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goaviatrix.Client)
	transitGatewayPeering := &goaviatrix.TransitGatewayPeering{
		TransitGatewayName1: d.Get("transit_gateway_name1").(string),
		TransitGatewayName2: d.Get("transit_gateway_name2").(string),
	}

	err := client.GetTransitGatewayPeering(transitGatewayPeering)
	if err != nil {
		if err == goaviatrix.ErrNotFound {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("couldn't find Aviatrix Transit Gateway peering: %s", err)
	}
	d.SetId(transitGatewayPeering.TransitGatewayName1 + "<->" + transitGatewayPeering.TransitGatewayName2)
	log.Printf("[INFO] Found Transit Gateway peering: %#v", d)
	return nil
}

func resourceTransitGatewayPeeringUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("the AviatrixTransitGatewayPeering resource doesn't support update")
}

func resourceTransitGatewayPeeringDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goaviatrix.Client)
	transitGatewayPeering := &goaviatrix.TransitGatewayPeering{
		TransitGatewayName1: d.Get("transit_gateway_name1").(string),
		TransitGatewayName2: d.Get("transit_gateway_name2").(string),
	}

	log.Printf("[INFO] Deleting Aviatrix Transit Gateway peering: %#v", transitGatewayPeering)

	err := client.DeleteTransitGatewayPeering(transitGatewayPeering)
	if err != nil {
		return fmt.Errorf("failed to delete Aviatrix Transit Gateway peering: %s", err)
	}
	return nil
}
