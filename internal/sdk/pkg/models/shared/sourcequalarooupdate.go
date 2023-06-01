// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceQualarooUpdate struct {
	// A Qualaroo token. See the <a href="https://help.qualaroo.com/hc/en-us/articles/201969438-The-REST-Reporting-API">docs</a> for instructions on how to generate it.
	Key string `json:"key"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate string `json:"start_date"`
	// IDs of the surveys from which you'd like to replicate data. If left empty, data from all surveys to which you have access will be replicated.
	SurveyIds []string `json:"survey_ids,omitempty"`
	// A Qualaroo token. See the <a href="https://help.qualaroo.com/hc/en-us/articles/201969438-The-REST-Reporting-API">docs</a> for instructions on how to generate it.
	Token string `json:"token"`
}