/*
 * STACKIT DNS API
 *
 * This api provides dns
 *
 * API version: 1.0
 * Contact: dns@stackit.cloud
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// ResponseZoneAll for filtered zones.
type ZoneResponseZoneAll struct {
	ItemsPerPage int32 `json:"ItemsPerPage"`
	Message string `json:"message,omitempty"`
	TotalItems int32 `json:"totalItems"`
	TotalPages int32 `json:"totalPages"`
	Zones []DomainZone `json:"zones"`
}
