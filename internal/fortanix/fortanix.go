package fortanix

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fortanix/sdkms-client-go/sdkms"
	"github.com/rfjakob/gocryptfs/v2/internal/exitcodes"
	"github.com/rfjakob/gocryptfs/v2/internal/tlog"
)

const ENV_API string = "FORTANIX_API_KEY"

type FortanixConfig struct {
    Url         string
    SecretName  string
    ApiKey      string
}

type SecretClient struct {
	*sdkms.Client
}

func (c *SecretClient) GetSecretByName(ctx *context.Context, keyName string) (
	*sdkms.Sobject, error) {
	sobjDescriptor := sdkms.SobjectDescriptor{
		Name: &keyName,
	}
	client := c.Client
	sobject, err := client.ExportSobject(*ctx, sobjDescriptor)
	if err != nil {
		return nil, err
	}
	return sobject, nil
}
func newSecretClient(url string, apiKey string) *SecretClient {
	client := sdkms.Client{
		HTTPClient: http.DefaultClient,
		Auth:       sdkms.APIKey(apiKey),
		Endpoint:   url,
	}
	return &SecretClient{&client}
}
func GetConfig() (*FortanixConfig, error) {
	url, err := readDsmConfigTerminal("DSM URL: ")
	if err != nil {
		return nil, err
	}
	secretName, err := readDsmConfigTerminal("Secret Name: ")
	if err != nil {
		return nil, err
	}
	apiKey, err := readDsmConfigTerminal("Api Key: ")
	if err != nil {
		return nil, err
	}
    fortanixConfig := FortanixConfig{
        Url:        strings.TrimSuffix(url, "\n"),
        SecretName: strings.TrimSuffix(secretName,"\n"),
        ApiKey:     strings.TrimSuffix(apiKey,"\n"),
    }
	return &fortanixConfig, nil
}

func ReadApikey() (string){
	apiKey, err := readDsmConfigTerminal("Api Key: ")

	if err != nil {
        tlog.Fatal.Println(err)
        os.Exit(exitcodes.DSMError)
    }
	return strings.TrimSuffix(apiKey,"\n")
}

func readDsmConfigTerminal(prompt string) (string, error) {
	fmt.Fprintf(os.Stderr, prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("Could not read %v from terminal: %v\n", prompt, err)
	}
	if len(text) == 0 {
		return "", fmt.Errorf("%v is empty", prompt)
	}
	return text, nil
}

func Secret(config *FortanixConfig ) ([]byte){
    tlog.Info.Printf("DSM Secret: Retrieving secret from DSM")
    if !strings.Contains(config.Url, "http") {
        tlog.Info.Println("DSM Secret: Missing protocol scheme, resuming with https")
        config.Url = "https://" + config.Url 
    }
    client := newSecretClient(config.Url, config.ApiKey)
    ctx := context.Background()
    sobj, err:= client.GetSecretByName(&ctx, config.SecretName)
    if err != nil {
        tlog.Fatal.Println(err)
        os.Exit(exitcodes.DSMError)
    }
    return *sobj.Value
}




