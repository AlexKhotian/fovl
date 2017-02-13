package MachineLearning

// IDecisionTree makes a decision where and how to store files
type IDecisionTree interface {
	DecideWhereToStoreData()
}

// DecisionTree implemets
type DecisionTree struct {
}
