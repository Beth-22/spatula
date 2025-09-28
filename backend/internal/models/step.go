package models

type Step struct {
    ID         string `json:"id"`
    RecipeID   string `json:"recipe_id"`
    StepNumber int    `json:"step_number"`
    Instruction string `json:"instruction"`
}

type CreateStepRequest struct {
    StepNumber  int    `json:"step_number"`
    Instruction string `json:"instruction"`
}