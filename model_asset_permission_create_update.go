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

type AssetPermissionCreateUpdate struct {
	Id          string   `json:"id,omitempty"`
	OrgId       string   `json:"org_id,omitempty"`
	OrgName     string   `json:"org_name,omitempty"`
	Actions     []int32  `json:"actions,omitempty"`
	Name        string   `json:"name"`
	IsActive    bool     `json:"is_active,omitempty"`
	DateStart   string   `json:"date_start,omitempty"`
	DateExpired string   `json:"date_expired,omitempty"`
	Comment     string   `json:"comment,omitempty"`
	Users       []string `json:"users,omitempty"`
	UserGroups  []string `json:"user_groups,omitempty"`
	Assets      []string `json:"assets,omitempty"`
	Nodes       []string `json:"nodes,omitempty"`
	SystemUsers []string `json:"system_users"`
}
