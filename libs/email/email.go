package email

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type EmailRequestContent struct {
	Subject   string `json:"subject"`
	PlainText string `json:"plainText"`
	HTML      string `json:"html,omitempty"`
}

type EmailRequest struct {
	SenderAddress string `json:"senderAddress"`
	Recipients    struct {
		To []struct {
			Address string `json:"address"`
		} `json:"to"`
	} `json:"recipients"`
	Content EmailRequestContent `json:"content"`
}

func Send(recipient string, content EmailRequestContent) error {
	connectionString := os.Getenv("AZURE_EMAIL_CONNECTION_STRING")
	if connectionString == "" {
		return fmt.Errorf("AZURE_EMAIL_CONNECTION_STRING env variable cannot be null")
	}
	endpoint, accessKey := parseConnectionString(connectionString)

	senderAddress := os.Getenv("AZURE_EMAIL_SENDER_ADDRESS")
	if senderAddress == "" {
		return fmt.Errorf("AZURE_EMAIL_SENDER_ADDRESS env variable cannot be null")
	}
	payload := EmailRequest{
		SenderAddress: senderAddress,
	}

	payload.Recipients.To = append(payload.Recipients.To, struct {
		Address string `json:"address"`
	}{
		Address: recipient,
	})

	payload.Content.Subject = content.Subject
	payload.Content.PlainText = content.PlainText
	payload.Content.HTML = content.HTML

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal email payload: %w", err)
	}

	apiPath := "/emails:send?api-version=2025-09-01"

	fullUrl := endpoint + apiPath

	req, err := http.NewRequest(
		http.MethodPost,
		fullUrl,
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return fmt.Errorf("failed to create http request: %w", err)
	}

	date := time.Now().UTC().Format(http.TimeFormat)

	host, err := getHost(endpoint)
	if err != nil {
		return fmt.Errorf("failed to get host from endpoint: %w", err)
	}

	authHeader, contentHash, err := buildHMAC(http.MethodPost, apiPath, host, date, bodyBytes, accessKey)
	if err != nil {
		return fmt.Errorf("failed to build HMAC: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-content-sha256", contentHash)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send email request: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return fmt.Errorf("failed to send email, status: %s, body: %s", resp.Status, string(respBody))
	}

	fmt.Println("STATUS:", resp.Status)
	fmt.Println(string(respBody))
	return nil
}

func parseConnectionString(conn string) (endpoint string, accessKey string) {
	parts := strings.Split(conn, ";")

	for _, p := range parts {
		if strings.HasPrefix(strings.ToLower(p), "endpoint=") {
			endpoint = p[len("endpoint="):]
		}

		if strings.HasPrefix(strings.ToLower(p), "accesskey=") {
			accessKey = p[len("accesskey="):]
		}
	}

	endpoint = strings.TrimRight(endpoint, "/")

	return
}

func getHost(endpoint string) (string, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	return u.Host, nil
}

func buildHMAC(
	method string,
	path string,
	host string,
	date string,
	body []byte,
	accessKey string,
) (string, string, error) {

	hash := sha256.Sum256(body)

	contentHash := base64.StdEncoding.EncodeToString(hash[:])

	stringToSign := fmt.Sprintf(
		"%s\n%s\n%s;%s;%s",
		method,
		path,
		date,
		host,
		contentHash,
	)

	keyBytes, err := base64.StdEncoding.DecodeString(accessKey)
	if err != nil {
		return "", "", err
	}

	h := hmac.New(sha256.New, keyBytes)

	h.Write([]byte(stringToSign))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	auth := fmt.Sprintf(
		"HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=%s",
		signature,
	)

	return auth, contentHash, nil
}
