package tasks

type TaskDTO struct {
	ID int `json:"id" form:"id" query:"id" param:"id" validate:"required,gte=100"`
}
