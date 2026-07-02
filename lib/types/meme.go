package types

type MemeStruct struct {
	Owner   string `json:"owner" binding:"required"`
	Bucket  string `json:"bucket" binding:"-"`
	ID      string `json:"id" binding:"required"`
	Message string `json:"message" binding:"required"`
	// Kind = bucket scenario for small-object uploads: "memory" (default) or
	// "knowledgebase" (RAG chunk of a Chroma collection). Empty = memory.
	Kind string `json:"kind,omitempty" binding:"-"`
}

type MemeMeta struct {
	File  string
	Start uint64
	Size  uint64
}
