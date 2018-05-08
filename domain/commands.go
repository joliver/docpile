package domain

import (
	"bitbucket.org/jonathanoliver/docpile/events"
	"io"
)

type AddTag struct {
	Name string
}

type ImportManagedStreamingAsset struct {
	Name     string
	MIMEType string
	Body     io.ReadCloser
}
type ImportManagedAsset struct {
	Name     string
	MIMEType string
	Hash     events.SHA256Hash
}

type ImportCloudAsset struct {
	Name     string
	Provider string
	Resource string
}

type DefineDocument struct {
	Document DocumentDefinition
}