package handlers

import (
	"bitbucket.org/jonathanoliver/docpile/generic"
	"bitbucket.org/jonathanoliver/docpile/generic/applicators"
)

type DomainHandler struct {
	aggregate  generic.Aggregate
	applicator applicators.Applicator
}

func NewDomainHandler(aggregate generic.Aggregate, applicator applicators.Applicator) *DomainHandler {
	return &DomainHandler{aggregate: aggregate, applicator: applicator}
}

func (this *DomainHandler) Handle(message interface{}) generic.Result {
	result := this.aggregate.Handle(message)
	this.applicator.Apply(this.aggregate.Consume())
	return result
}