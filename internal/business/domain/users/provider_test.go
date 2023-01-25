package users

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestProvider_FindAll(t *testing.T) {
	tests := []struct {
		name       string
		dataSource io.Reader
		want       []User
		wantErr    bool
	}{
		{
			name:       "Default behavior",
			dataSource: strings.NewReader(`[{"login": "ehot_riss"}]`),
			want: []User{
				{
					Login: "ehot_riss",
				},
			},
			wantErr: false,
		},
		{
			name:       "Bad json - 1",
			dataSource: strings.NewReader(`[}]`),
			want:       []User{},
			wantErr:    true,
		},
		{
			name:       "Bad json - 1",
			dataSource: strings.NewReader(`{"login": "ehot_riss"}`),
			want:       []User{},
			wantErr:    true,
		},
		{
			name:       "Empty result",
			dataSource: strings.NewReader(`[]`),
			want:       []User{},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewProvider(tt.dataSource)
			got, err := p.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}
