// Code generated by goa v3.2.6, DO NOT EDIT.
//
// assetsvc gRPC client CLI support package
//
// Command:
// $ goa gen github.com/tinyci/ci-agents/design

package client

import (
	"encoding/json"
	"fmt"

	assetsvc "github.com/tinyci/ci-agents/gen/assetsvc"
	assetsvcpb "github.com/tinyci/ci-agents/gen/grpc/assetsvc/pb"
)

// BuildGetLogPayload builds the payload for the assetsvc getLog endpoint from
// CLI flags.
func BuildGetLogPayload(assetsvcGetLogMessage string) (*assetsvc.GetLogPayload, error) {
	var err error
	var message assetsvcpb.GetLogRequest
	{
		if assetsvcGetLogMessage != "" {
			err = json.Unmarshal([]byte(assetsvcGetLogMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 8638009270681567038\n   }'")
			}
		}
	}
	v := &assetsvc.GetLogPayload{
		ID: message.Id,
	}

	return v, nil
}
