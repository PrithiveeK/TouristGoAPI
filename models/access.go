package models

//Access model for selecting only the important piece of information
type Access struct {
	ID       int64 `json:"id"`
	RoleID   int64 `json:"role_id"`
	AccessID int64 `json:"access_id"`
}
