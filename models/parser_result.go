package models

type ParseSelectionResult struct {
	Value      *string
	TagHref    *string
	FoundValue []ParseSelectionResult
}
