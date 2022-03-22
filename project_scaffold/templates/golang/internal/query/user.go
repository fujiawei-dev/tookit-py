{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"{{GOLANG_MODULE}}/internal/acl"
)

type UserResult struct {
	ID       uint     `json:"id" example:"1"` // 记录ID
	Username string   `json:"username" example:"用户名"`
	Enable    bool        `json:"enable"`                                                 // 是否启用
	CreatedAt time.Time   `json:"created_at" example:"2022-03-21T08:57:19.4615214+08:00"` // 创建时间
}
