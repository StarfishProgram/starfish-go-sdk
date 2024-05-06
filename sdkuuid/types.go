package sdkuuid

import "github.com/StarfishProgram/starfish-go-sdk/sdktypes"

type UUID interface {
	Id() sdktypes.ID
	Uuid() string
	TimeID() string
}
