package MachineLearning

// IDecisionTree makes a decision where and how to store files
type IDecisionTree interface {
	DecideWhereToStoreData(filename *string) string
}

// DecisionTree implemets
type DecisionTree struct {
}

// DecideWhereToStoreData decides where to store data base on type/namespace
func (instance *DecisionTree) DecideWhereToStoreData(filename *string) string {
	return ""
}
