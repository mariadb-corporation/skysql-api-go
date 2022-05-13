// main Example code to interact with the SkySQL API
// Don't use this in production, the error handling makes no sense, all this' just minimal
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path"

	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	skysql "github.com/mariadb-corporation/skysql-api-go"
)

func main() {
	ctx := context.TODO()

	// get API Key from the user's environment
	apiKey := os.Getenv("SKYSQL_API_KEY")

	// Setup client
	client, _ := skysql.NewClientWithResponses(
		"https://api.skysql.net",
		skysql.WithRequestEditorFn(bearerToken(apiKey).Intercept),
		skysql.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("User-Agent", "skysql-client-go-example")
			return nil
		}),
	)

	// Test a random API (Provider in this case)
	res, _ := client.ReadProvidersWithResponse(ctx, &skysql.ReadProvidersParams{})

	// Print response if 200
	if res.JSON200 != nil {
		// Pretty print the available providers
		for _, provider := range *res.JSON200 {
			fmt.Println(provider.Name)
		}
	}
}

// bearerToken Get a Bearer Token from MariaDB ID
func bearerToken(apiKey string) *securityprovider.SecurityProviderBearerToken {
	mdbidURL := "https://id.mariadb.com"
	url, err := url.Parse(mdbidURL)
	if err != nil || url.String() == "" {
		panic(fmt.Sprintf("Invalid URL for MariaDB ID %s", err))
	}

	url.Path = path.Join(url.Path, "/api/v1/token")
	req, _ := http.NewRequest("POST", url.String(), nil)
	req.Header.Add("Authorization", "Token "+apiKey)

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)

	if err != nil {
		panic(fmt.Sprintf("Failure during authentication, %s", err))
	}
	defer res.Body.Close()

	var resData struct {
		Token string `json:"token"`
	}

	json.NewDecoder(res.Body).Decode(&resData)

	bearerTokenProvider, _ := securityprovider.NewSecurityProviderBearerToken(resData.Token)

	return bearerTokenProvider
}
