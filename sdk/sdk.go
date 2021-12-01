package sdk

import (
	"github.com/SarasovMatvey/godesk/client"
	"github.com/SarasovMatvey/godesk/sdk/api"
)

type Sdk struct {
	C2dClient client.Client

	Webhooks api.Webhooks
	Messages api.Messages
	Clients  api.Clients
}

func NewSdk(client client.Client) Sdk {
	return Sdk{
		C2dClient: client,

		Webhooks: api.Webhooks{Client: client},
		Messages: api.Messages{Client: client},
		Clients:  api.Clients{Client: client},
	}
}
