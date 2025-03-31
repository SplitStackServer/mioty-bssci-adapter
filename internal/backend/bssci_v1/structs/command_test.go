package structs

import "testing"

func TestCommandHeader_GetCommand(t *testing.T) {
	type fields struct {
		Command Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   Command
	}{
		{
			name: "command",
			fields: fields{
				Command: "test",
				OpId:    1,
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &CommandHeader{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); got != tt.want {
				t.Errorf("CommandHeader.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandHeader_GetOpId(t *testing.T) {
	type fields struct {
		Command Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "command",
			fields: fields{
				Command: "test",
				OpId:    1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &CommandHeader{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("CommandHeader.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}
