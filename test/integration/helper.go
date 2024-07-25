package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Khan/genqlient/graphql"
)

func WaitForInput(ctx context.Context, client graphql.Client) error {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()
	for {
		result, err := getInputStatus(ctx, client, 0)
		if err != nil && !strings.Contains(err.Error(), "input not found") {
			return fmt.Errorf("failed to get input status: %w", err)
		}
		if result.Input.Status == CompletionStatusAccepted {
			return nil
		}
		select {
		case <-ticker.C:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func IncreaseTime(url string, seconds int) error {
	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "evm_increaseTime",
		"params":  []int{seconds},
		"id":      67,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error while marshaling JSON: %w", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error while making POST: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("error while decoding JSON: %w", err)
	}

	fmt.Println(result)
	return nil
}
