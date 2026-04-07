package telegramClient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{Timeout: 10 * time.Second},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest("getUpdates", params)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON getUpdates response: %w", err)
	}

	return res.Result, nil
}

func (c *Client) SendMessage(chatID int, text string) error {
	params := url.Values{}
	params.Add("chat_id", strconv.Itoa(chatID))
	params.Add("text", text)

	_, err := c.doRequest("sendMessage", params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SendPhoto(chatID int, photoURL string, caption string) error {
	params := url.Values{}
	params.Add("chat_id", strconv.Itoa(chatID))
	params.Add("photo", photoURL) // ← сюда URL картинки
	params.Add("caption", caption)

	_, err := c.doRequest("sendPhoto", params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) doRequest(method string, params url.Values) ([]byte, error) {
	uri := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error composing request: %w", err)
	}

	req.URL.RawQuery = params.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error during request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error during request: %w", err)
	}
	return data, nil
}
