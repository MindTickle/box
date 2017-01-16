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
		Id   string
		Type string
	}
	CreatedBy struct {
		Login string
		Name  string
		Type  string
	} `json:"created_by"`
	CreatedAt string `json:"created_at"`
	Address   string
	Triggers  []string
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
func (c *WebhookService) DeleteTask(webhookId string) (*http.Response, error) {
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
