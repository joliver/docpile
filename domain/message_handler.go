package domain

import (
	"fmt"
	"reflect"
	"log"
)

type MessageHandler struct {
	root       *Aggregate
	applicator Applicator
}

func NewMessageHandler(app *Aggregate) *MessageHandler {
	return &MessageHandler{root: app}
}

func (this *MessageHandler) Handle(message interface{}) (uint64, error) {
	if id, err := this.handle(message); err != nil {
		return 0, err
	} else {
		this.applicator.Apply(this.root.Consume())
		return id, nil
	}
}
func (this *MessageHandler) handle(message interface{}) (uint64, error) {
	switch message := message.(type) {
	case AddTag:
		return this.root.AddTag(message.Name)
	case ImportManagedAsset:
		return this.root.ImportManagedAsset(message.Name, message.MIMEType, message.Hash)
	case ImportCloudAsset:
		return this.root.ImportCloudAsset(message.Name, message.Provider, message.Resource)
	case DefineDocument:
		return this.root.DefineDocument(message.Document)
	default:
		log.Panicf(fmt.Sprintf("MessageHandler cannot handle '%s'", reflect.TypeOf(message)))
	}

	return 0, nil
}