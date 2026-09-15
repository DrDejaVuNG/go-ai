package responses

import (
	"encoding/json"
	"testing"

	"github.com/digitallysavvy/go-ai/pkg/provider/types"
)

func TestFunctionCallOutputPairing(t *testing.T) {
	prompt := types.Prompt{Messages: []types.Message{
		{Role: types.RoleSystem, Content: []types.ContentPart{types.TextContent{Text: "sys"}}},
		{Role: types.RoleUser, Content: []types.ContentPart{types.TextContent{Text: "hi"}}},
		{Role: types.RoleAssistant, ToolCalls: []types.ToolCall{{
			ID: "call_UIvII19htpXgJmmoUDcdPcDg", ToolName: "get_trending_conversations",
			Arguments: map[string]interface{}{},
		}}, Content: []types.ContentPart{types.TextContent{Text: ""}}},
		{Role: types.RoleTool, Content: []types.ContentPart{
			types.SimpleJSONResult("call_UIvII19htpXgJmmoUDcdPcDg", "get_trending_conversations", map[string]string{"status": "ok"}),
		}},
	}}
	input := ConvertPromptToInput(prompt, "")
	b, _ := json.MarshalIndent(input, "", "  ")
	t.Logf("input:\n%s", b)
}

func TestSystemMessagePreserved(t *testing.T) {
	prompt := types.Prompt{Messages: []types.Message{
		{Role: types.RoleSystem, Content: []types.ContentPart{types.TextContent{Text: "You are Streetz AI"}}},
		{Role: types.RoleUser, Content: []types.ContentPart{types.TextContent{Text: "hi"}}},
	}}
	input := ConvertPromptToInput(prompt, "system")
	first, ok := input[0].(SystemMessage)
	if !ok {
		t.Fatalf("expected SystemMessage first, got %T", input[0])
	}
	if first.Content != "You are Streetz AI" || first.Role != "system" {
		t.Fatalf("system message not preserved: %+v", first)
	}
}
