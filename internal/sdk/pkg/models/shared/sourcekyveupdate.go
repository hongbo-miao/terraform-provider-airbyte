// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceKyveUpdate struct {
	// The maximum amount of pages to go trough. Set to 'null' for all pages.
	MaxPages *int64 `json:"max_pages,omitempty"`
	// The pagesize for pagination, smaller numbers are used in integration tests.
	PageSize *int64 `json:"page_size,omitempty"`
	// The IDs of the KYVE storage pool you want to archive. (Comma separated)
	PoolIds string `json:"pool_ids"`
	// The start-id defines, from which bundle id the pipeline should start to extract the data (Comma separated)
	StartIds string `json:"start_ids"`
	// URL to the KYVE Chain API.
	URLBase *string `json:"url_base,omitempty"`
}