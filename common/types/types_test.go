package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAny(t *testing.T) {
	a := NewAny()
	assert.NotNil(t, a)
	assert.NotNil(t, a.value)
}

func TestAny_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		any     *Any
		want    []byte
		wantErr bool
	}{
		{
			name: "nil Any",
			any:  nil,
			want: []byte("null"),
		},
		{
			name: "nil value",
			any:  &Any{value: nil},
			want: []byte("null"),
		},
		{
			name: "string value",
			any:  &Any{value: "test"},
			want: []byte(`"test"`),
		},
		{
			name: "int value",
			any:  &Any{value: 123},
			want: []byte(`123`),
		},
		{
			name: "map value",
			any:  &Any{value: map[string]string{"key": "value"}},
			want: []byte(`{"key":"value"}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.any.MarshalJSON()
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAny_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    *Any
		wantErr bool
	}{
		{
			name: "string value",
			data: []byte(`"test"`),
			want: &Any{value: "test"},
		},
		{
			name: "int value",
			data: []byte(`123`),
			want: &Any{value: float64(123)}, // JSON unmarshals numbers to float64
		},
		{
			name: "map value",
			data: []byte(`{"key":"value"}`),
			want: &Any{value: map[string]interface{}{"key": "value"}},
		},
		{
			name:    "nil Any",
			data:    []byte(`"test"`),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var a *Any
			if tt.want != nil {
				a = &Any{}
			}
			err := a.UnmarshalJSON(tt.data)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want.value, a.value)
		})
	}
}

func TestAny_Value(t *testing.T) {
	tests := []struct {
		name string
		any  *Any
		want any
	}{
		{
			name: "string value",
			any:  &Any{value: "test"},
			want: "test",
		},
		{
			name: "int value",
			any:  &Any{value: 123},
			want: 123,
		},
		{
			name: "nil value",
			any:  &Any{value: nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.any.Value()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAny_Decode(t *testing.T) {
	type TestStruct struct {
		Name  string
		Value int
	}

	tests := []struct {
		name    string
		any     *Any
		out     *TestStruct
		want    *TestStruct
		wantErr bool
	}{
		{
			name: "valid map",
			any: &Any{value: map[string]interface{}{
				"name":  "test",
				"value": 123,
			}},
			out: &TestStruct{},
			want: &TestStruct{
				Name:  "test",
				Value: 123,
			},
		},
		{
			name: "invalid map",
			any: &Any{value: map[string]interface{}{
				"name":  "test",
				"value": "not an int",
			}},
			out:     &TestStruct{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.any.Decode(tt.out)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, tt.out)
		})
	}
}
