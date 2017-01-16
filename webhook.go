package box

import (
	"fmt"
	"net/http"
)

type WebhookService struct {
	*Client
}

type Webhook struct {
	Type   string `json:"type"`
	Id     string `json:"id"`
	Target struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	}
	CreatedBy struct {
		Login string `json:"login"`
		Name  string `json:"name"`
		Type  string `json:"type"`
	} `json:"created_by"`
	CreatedAt string   `json:"created_at"`
	Address   string   `json:"address"`
	Triggers  []string `json:"triggers"`
}

func (c *WebhookService) CreateWebhook(id, ftype, webhookUrl string, triggers []string) (*http.Response, *Webhook, error) {
	var dataMap = map[string]interface{}{
		"target": map[string]string{
			"id":   id,
			"type": ftype,
		},
		"address":  webhookUrl,
		"triggers": triggers,
	}
	req, err := c.NewRequest(
		"POST",
		fmt.Sprintf("/webhooks"),
		dataMap,
	)
	if err != nil {
		return nil, nil, err
	}

	var data Webhook
	resp, err := c.Do(req, &data)
	return resp, &data, err
}
func (c *WebhookService) DeleteWebhook(webhookId string) (*http.Response, error) {
	req, err := c.NewRequest(
		"DELETE",
		fmt.Sprintf("/webhooks/%s", webhookId),
		nil,
	)
	if err != nil {
		return nil, err
	}

	return c.Do(req, nil)
}
