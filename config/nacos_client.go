package config

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	nacosHost     = "127.0.0.1"
	nacosPort     = 8800
	nacosUsername = "nacos"
	nacosPassword = "nacos"
)

var nacosHTTPClient = &http.Client{Timeout: 5 * time.Second}

/*
*
获取nacos的连接字符串
*/
func nacosBaseURL() string {
	return fmt.Sprintf("http://%s:%d/nacos", nacosHost, nacosPort)
}

func nacosNamespaceID() string {
	return ""
}

func nacosAuthParams(values url.Values) url.Values {
	if values == nil {
		values = url.Values{}
	}

	values.Set("username", nacosUsername)
	values.Set("password", nacosPassword)

	namespaceID := nacosNamespaceID()
	if namespaceID != "" {
		values.Set("tenant", namespaceID)
		values.Set("namespaceId", namespaceID)
	}

	return values
}

func fetchNacosConfig(dataID, group string) (string, error) {
	values := nacosAuthParams(url.Values{})
	values.Set("dataId", dataID)
	values.Set("group", group)

	resp, err := nacosHTTPClient.Get(nacosBaseURL() + "/v1/cs/configs?" + values.Encode())
	if err != nil {
		return "", fmt.Errorf("request nacos config failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read nacos config response failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("nacos config request failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	return string(body), nil
}

func registerServiceHTTP(serviceName, ip string, port uint64, group string) error {
	values := nacosAuthParams(url.Values{})
	values.Set("serviceName", serviceName)
	values.Set("ip", ip)
	values.Set("port", strconv.FormatUint(port, 10))
	values.Set("groupName", group)
	values.Set("ephemeral", "true")
	values.Set("healthy", "true")
	values.Set("enabled", "true")

	resp, err := nacosHTTPClient.PostForm(nacosBaseURL()+"/v1/ns/instance", values)
	if err != nil {
		return fmt.Errorf("request nacos register failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read nacos register response failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("nacos register request failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	return nil
}

func parseServicePort(port string) (uint64, error) {
	value, err := strconv.ParseUint(port, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid server port %q: %w", port, err)
	}
	return value, nil
}
