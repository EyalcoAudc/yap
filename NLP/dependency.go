package NLP

import (
	. "chukuparser/Algorithm/Graph"
	"chukuparser/Util"
)

type DepNode interface {
	Vertex
	String() string
}

type DepArc interface {
	DirectedEdge
	GetModifier() int
	GetHead() int
	String() string
}

type DepRel string

type LabeledDepArc interface {
	DepArc
	GetRelation() DepRel
}

type Labeled interface {
	GetLabeledArc(int) LabeledDepArc
}

type DependencyGraph interface {
	DirectedGraph
	GetNode(int) DepNode
	GetArc(int) DepArc
	NumberOfNodes() int
	NumberOfArcs() int
	Equal(otherEq Util.Equaler) bool
	Sentence() Sentence
	TaggedSentence() TaggedSentence
}

type LabeledDependencyGraph interface {
	DependencyGraph
	Labeled
}
