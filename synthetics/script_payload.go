package synthetics

// ScriptPayload struct representing json from Api
// end point: https://synthetics.newrelic.com/synthetics/api/v3/monitors/{id}/script
type ScriptPayload struct {
	ScriptText string `json:"scriptText"`
}
