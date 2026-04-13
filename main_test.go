package main

import "testing"

func TestExtractAndValidateJSON(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{
			name: "valid object",
			in:   `{"id":1,"name":"test"}`,
			want: "{\n  \"id\": 1,\n  \"name\": \"test\"\n}",
		},
		{
			name: "valid array",
			in:   `[1,2,3]`,
			want: "[\n  1,\n  2,\n  3\n]",
		},
		{
			name: "nested object",
			in:   `{"a":{"b":"c"}}`,
			want: "{\n  \"a\": {\n    \"b\": \"c\"\n  }\n}",
		},
		{
			name: "json embedded in text",
			in:   `INFO: result: {"key":"val"} done`,
			want: "{\n  \"key\": \"val\"\n}",
		},
		{
			name: "json array embedded in text",
			in:   `output: [1, 2, 3] end`,
			want: "[\n  1,\n  2,\n  3\n]",
		},
		{
			name: "incomplete object - one closing brace present",
			in:   `{"data": {"item": "value"}`,
			want: "{\n  \"data\": {\n    \"item\": \"value\"\n  }\n}",
		},
		{
			name: "incomplete object - no closing brace at all",
			in:   `{"data": {"item": "value"`,
			want: "{\n  \"data\": {\n    \"item\": \"value\"\n  }\n}",
		},
		{
			name: "incomplete array - missing closing bracket",
			in:   `[1, 2, [3, 4]`,
			want: "[\n  1,\n  2,\n  [\n    3,\n    4\n  ]\n]",
		},
		{
			name:    "no json in input",
			in:      "hello world no json here",
			wantErr: true,
		},
		{
			name:    "empty input",
			in:      "",
			wantErr: true,
		},
		{
			name: "already pretty json",
			in:   "{\n  \"key\": \"val\"\n}",
			want: "{\n  \"key\": \"val\"\n}",
		},
		{
			name: "multiline text with json",
			in:   "Some log output\n{\"status\":\"ok\",\"count\":42}\nMore text",
			want: "{\n  \"status\": \"ok\",\n  \"count\": 42\n}",
		},
		// jsonfix-powered capabilities (not possible with old regex approach)
		{
			name: "markdown code fence",
			in:   "Here is the result:\n```json\n{\"key\": \"value\"}\n```",
			want: "{\n  \"key\": \"value\"\n}",
		},
		{
			name: "single-quoted keys and values",
			in:   `{'key': 'value'}`,
			want: "{\n  \"key\": \"value\"\n}",
		},
		{
			name: "trailing comma in object",
			in:   `{"a": 1, "b": 2,}`,
			want: "{\n  \"a\": 1,\n  \"b\": 2\n}",
		},
		{
			name: "trailing comma in array",
			in:   `[1, 2, 3,]`,
			want: "[\n  1,\n  2,\n  3\n]",
		},
		{
			name: "unquoted keys",
			in:   `{name: "test", count: 42}`,
			want: "{\n  \"name\": \"test\",\n  \"count\": 42\n}",
		},
		{
			name: "python-style booleans and none",
			in:   `{"active": True, "deleted": False, "data": None}`,
			want: "{\n  \"active\": true,\n  \"deleted\": false,\n  \"data\": null\n}",
		},
		{
			name: "comments in json",
			in:   "{\"key\": \"val\",\n// comment line\n\"key2\": \"val2\"\n}",
			want: "{\n  \"key\": \"val\",\n  \"key2\": \"val2\"\n}",
		},
		{
			name: "deeply nested incomplete json",
			in:   `{"a": {"b": {"c": {"d": "value"`,
			want: "{\n  \"a\": {\n    \"b\": {\n      \"c\": {\n        \"d\": \"value\"\n      }\n    }\n  }\n}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractAndValidateJSON(tt.in)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil (result: %q)", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got:\n%s\nwant:\n%s", got, tt.want)
			}
		})
	}
}
