package MachineLearning

// IDecisionTree makes a decision where and how to store files
type IDecisionTree interface {
	DecideWhereToStoreData(filename *string) string
}

// DecisionTree implemets
type DecisionTree struct {
}

func (instance *DecisionTree) DecideWhereToStoreData(filename *string) string {

}
