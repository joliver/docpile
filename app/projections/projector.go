package projections

type Projector struct {
	AllTags      *AllTags
	AllDocuments *AllDocuments
	MatchingTags *MatchingTags
}

func NewProjector() *Projector {
	tags := NewAllTags()
	docs := NewAllDocuments()

	return &Projector{
		AllTags:      tags,
		AllDocuments: docs,
		MatchingTags: NewMatchingTags(docs, tags),
	}
}

///////////////////////////////////////////

func (this *Projector) Apply(messages []interface{}) {
	for _, message := range messages {
		this.apply(message)
	}
}
func (this *Projector) apply(message interface{}) {
	this.AllTags.Transform(message)
	this.AllDocuments.Transform(message)
}
