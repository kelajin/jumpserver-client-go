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

type Session struct {
	Id               string `json:"id,omitempty"`
	User             string `json:"user"`
	Asset            string `json:"asset"`
	SystemUser       string `json:"system_user"`
	LoginFrom        string `json:"login_from,omitempty"`
	LoginFromDisplay string `json:"login_from_display,omitempty"`
	RemoteAddr       string `json:"remote_addr,omitempty"`
	IsFinished       bool   `json:"is_finished,omitempty"`
	HasReplay        bool   `json:"has_replay,omitempty"`
	CanReplay        string `json:"can_replay,omitempty"`
	Protocol         string `json:"protocol,omitempty"`
	DateStart        string `json:"date_start,omitempty"`
	DateEnd          string `json:"date_end,omitempty"`
	Terminal         string `json:"terminal,omitempty"`
	CommandAmount    int32  `json:"command_amount,omitempty"`
	OrgId            string `json:"org_id"`
	OrgName          string `json:"org_name,omitempty"`
}
