package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mweesenaar/terraform-provider-nginxproxymanager/client/resources"
	"net/http"
)

func (c *Client) GetRedirectionHosts(ctx context.Context) (*resources.RedirectionHostCollection, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/api/nginx/redirection-hosts", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ar := resources.RedirectionHostCollection{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (c *Client) GetRedirectionHost(ctx context.Context, id *int64) (*resources.RedirectionHost, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/api/nginx/redirection-hosts/%d", c.HostURL, *id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ar := resources.RedirectionHost{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	if ar.ID == 0 {
		return nil, nil
	}

	return &ar, nil
}
