package sdk

import (
	"github.com/SarasovMatvey/godesk/client"
	"github.com/SarasovMatvey/godesk/sdk/api"
)

type Sdk struct {
	C2dClient client.Client

	Webhooks   api.Webhooks
	Messages   api.Messages
	Clients    api.Clients
	Dialogs    api.Dialogs
	Operators  api.Operators
	Statistics api.Statistics
}

func NewSdk(client client.Client) Sdk {
	return Sdk{
		C2dClient: client,

		Webhooks:   api.Webhooks{Client: client},
		Messages:   api.Messages{Client: client},
		Clients:    api.Clients{Client: client},
		Dialogs:    api.Dialogs{Client: client},
		Operators:  api.Operators{Client: client},
		Statistics: api.Statistics{Client: client},
	}
}
