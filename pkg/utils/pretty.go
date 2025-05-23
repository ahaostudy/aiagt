package utils

import (
	"bytes"
	"encoding/json"
)

func Pretty(v any, max ...int) string {
	resultBytes := FirstResult(json.Marshal(v))

	if len(max) > 0 && max[0] > 0 && len(resultBytes) > max[0] {
		builder := bytes.NewBuffer(resultBytes[:max[0]])
		builder.WriteString("...")

		return builder.String()
	}

	return string(resultBytes)
}

func PrettyIndent(v any, prefix string, indent string, max ...int) string {
	resultBytes := FirstResult(json.MarshalIndent(v, prefix, indent))

	if len(max) > 0 && max[0] > 0 && len(resultBytes) > max[0] {
		builder := bytes.NewBuffer(resultBytes[:max[0]])
		builder.WriteString("...")

		return builder.String()
	}

	return string(resultBytes)
}

func PrettyBytes(v []byte, max ...int) string {
	if len(max) > 0 && max[0] > 0 && len(v) > max[0] {
		return string(v[:max[0]]) + "..."
	}

	return string(v)
}
