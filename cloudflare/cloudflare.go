package cloudflare

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func CreateSubdomain(subdomain string, ip string) error {
    ctx := context.Background()
	
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CF_API_TOKEN"))
    if err != nil {
        return err
    }

    zoneID, err := api.ZoneIDByName("piglin.cloud")
    if err != nil {
        return err
    }

	params := cloudflare.CreateDNSRecordParams{
		Type:		"A",
		Name:		subdomain + ".piglin.cloud",
		Content:	ip,
		TTL:		120,
	}

    rc := &cloudflare.ResourceContainer{
		Identifier: zoneID,
	}

	record, err := api.CreateDNSRecord(ctx, rc, params)
	if err != nil {
		return err
	}

	fmt.Printf("DNS Record created successfully: %+v\n", record)
    return nil
}
