package v3

type ConfigOperation string

var (
	ConfigDeleteOperation ConfigOperation = "DELETE"
)

type ConfigBatchAlterData struct {
	Name      string          `json:"name"`
	Value     string          `json:"value,omitempty"`
	Operation ConfigOperation `json:"operation,omitempty"`
}

type ConfigBatchAlterRequest struct {
	Data []ConfigBatchAlterData `json:"data"`
}
