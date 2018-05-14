package http

import (
	"bitbucket.org/jonathanoliver/docpile/app/domain"
	"bitbucket.org/jonathanoliver/docpile/app/http/inputs"
	"bitbucket.org/jonathanoliver/docpile/generic/handlers"
	"github.com/smartystreets/detour"
)

type DocumentWriter struct {
	handler handlers.Handler
}

func NewDocumentWriter(handler handlers.Handler) *DocumentWriter {
	return &DocumentWriter{handler: handler}
}

func (this *DocumentWriter) Define(input *inputs.DefineDocument) detour.Renderer {
	return this.renderResult(domain.DefineDocument{
		Document: domain.DocumentDefinition{
			AssetID:     input.AssetID,
			AssetOffset: input.AssetOffset,
			Published:   input.Published,
			PeriodMin:   input.PeriodMin,
			PeriodMax:   input.PeriodMax,
			Tags:        input.Tags,
			Documents:   input.Documents,
			Description: input.Description,
		},
	})
}
func (this *DocumentWriter) Remove(input *inputs.IDInput) detour.Renderer {
	return this.renderResult(domain.RemoveDocument{ID: input.ID})
}

func (this *DocumentWriter) renderResult(message interface{}) detour.Renderer {
	if result := this.handler.Handle(message); result.Error == nil {
		return newEntityResult(result.ID)
	} else if result.Error == domain.AssetNotFoundError {
		return inputs.AssetDoesNotExistResult
	} else if result.Error == domain.TagNotFoundError {
		return inputs.TagDoesNotExistResult
	} else if result.Error == domain.DocumentNotFoundError {
		return inputs.DocumentDoesNotExistResult
	} else {
		return UnknownErrorResult
	}

}
