package cloudflare

import (
	"time"

	"github.com/Sirupsen/logrus"
	cloudflare "github.com/cloudflare/cloudflare-go"
)

// Client holds a cloudflare.API to interact with cloudflare
type Client struct {
	api *cloudflare.API
}

// New returns a Client configured with a key and email
func New(key, email string) (*Client, error) {
	c, err := cloudflare.New(key, email)
	if err != nil {
		return nil, err
	}

	return &Client{
		api: c,
	}, nil
}

// KeepUpdated checks cloudflare every 30 seconds, and update the DNS record if necessary.
func (c *Client) KeepUpdated(host string) error {
	for {
		if err := c.update(host); err != nil {
			logrus.WithError(err).Error("unable to update ddns")
		}

		time.Sleep(30 * time.Second)
	}
}

func (c *Client) update(host string) error {
	zones, err := c.api.ListZones()
	if err != nil {
		return err
	}

	myIP, err := getIPAddress()
	if err != nil {
		return err
	}

	for _, zone := range zones {
		rr := cloudflare.DNSRecord{
			Name: host,
		}

		records, err := c.api.DNSRecords(zone.ID, rr)
		if err != nil {
			return err
		}

		for _, record := range records {
			if record.Type == "A" && record.Name == rr.Name {
				if myIP.String() == record.Content {
					logrus.WithField("ip", myIP.String()).Info("DNS up-to-date")
					return nil
				}

				logrus.WithField("ip", myIP.String()).Info("updating dns")
				record.Content = myIP.String()
				if err := c.api.UpdateDNSRecord(zone.ID, record.ID, record); err != nil {
					logrus.WithError(err).Errorf("unable to update record")
					return err
				}

				return nil
			}
		}

		logrus.WithField("host", rr.Name).Errorf("Record not found on Cloudflare")
	}

	return nil
}
