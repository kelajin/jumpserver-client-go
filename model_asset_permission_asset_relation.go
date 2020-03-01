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

type AssetPermissionAssetRelation struct {
	Id int32 `json:"id,omitempty"`
	Asset string `json:"asset"`
	AssetDisplay string `json:"asset_display,omitempty"`
	Assetpermission string `json:"assetpermission"`
	AssetpermissionDisplay string `json:"assetpermission_display,omitempty"`
}