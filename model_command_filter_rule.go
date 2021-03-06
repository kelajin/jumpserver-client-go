/*
 * Jumpserver API Docs
 *
 * Jumpserver Restful api docs
 *
 * API version: v1
 * Contact: support@fit2cloud.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CommandFilterRule struct {
	Id      string `json:"id,omitempty"`
	OrgId   string `json:"org_id,omitempty"`
	OrgName string `json:"org_name,omitempty"`
	Type_   string `json:"type,omitempty"`
	// 优先级可选范围为1-100，1最低优先级，100最高优先级
	Priority int32 `json:"priority,omitempty"`
	// 每行一个命令
	Content     string `json:"content"`
	Action      int32  `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
	DateUpdated string `json:"date_updated,omitempty"`
	CreatedBy   string `json:"created_by,omitempty"`
	Filter      string `json:"filter"`
}
