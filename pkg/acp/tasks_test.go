package acp

import (
	"encoding/json"
	"testing"
)

func TestTranslateTask(t *testing.T) {
	r := NewRouter()
	input, _ := json.Marshal(map[string]string{
		"text": "Hello", "source": "en", "target": "ar",
	})
	out, err := r.Handle(t.Context(), TaskTranslate, input)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]interface{}
	json.Unmarshal(out, &result)
	engine := result["engine"].(string)
	if engine != "mymemory.translated.net" && engine != "stub/v1-fallback" {
		t.Fatalf("unexpected engine: %v", engine)
	}
}

func TestExecuteTask(t *testing.T) {
	r := NewRouter()
	input, _ := json.Marshal(map[string]interface{}{
		"action": "hash",
		"params": map[string]string{"text": "test"},
	})
	out, err := r.Handle(t.Context(), TaskExecute, input)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]string
	json.Unmarshal(out, &result)
	if len(result["hash"]) != 64 {
		t.Fatalf("expected sha256 hex, got %v", result)
	}

	// إجراء غير مسموح
	bad, _ := json.Marshal(map[string]string{"action": "rm_rf"})
	if _, err := r.Handle(t.Context(), TaskExecute, bad); err == nil {
		t.Fatal("disallowed action should fail")
	}
}
