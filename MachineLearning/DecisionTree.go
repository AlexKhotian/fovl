package MachineLearning

import "fovl/FileSession"

// IDecisionTree makes a decision where and how to store files
type IDecisionTree interface {
	DecideWhereToStoreData(filename *string) string
}

// DecisionTree implemets
type DecisionTree struct {
}

// DecideWhereToStoreData decides where to store data base on type/namespace
func (instance *DecisionTree) DecideWhereToStoreData(filename *string) string {
	fileType := ParseFile(filename)
	switch fileType {
	case Image:
		return FileSession.PhotoDir
	case Media:
	case Literature:
		return FileSession.MediaDir
	case Documents:
	case Presentation:
		return FileSession.DocsDir
	case SourceCode:
		return FileSession.CodeDir
	}

	return FileSession.PhotoDir
}
