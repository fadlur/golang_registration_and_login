package response

import (
	"registration/model"
	"time"
)

type RoleResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ConvertRoleResponse(role model.Role) RoleResponse {
	var roleResponse RoleResponse
	roleResponse.Id = int(role.ID)
	roleResponse.Name = role.Name
	roleResponse.Description = role.Description
	roleResponse.CreatedAt = role.CreatedAt
	roleResponse.UpdatedAt = role.UpdatedAt
	return roleResponse
}