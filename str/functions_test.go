package str

import (
	"testing"

	"github.com/dracory/base/arr"
)

func Test_fieldsFunc(t *testing.T) {
	type args struct {
		s            string
		f            func(rune) bool
		preserveFunc []func(rune) bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty string",
			args: args{
				s:            "",
				f:            func(r rune) bool { return r == ' ' },
				preserveFunc: []func(rune) bool{},
			},
			want: []string{},
		},
		{
			name: "single word",
			args: args{
				s:            "hello",
				f:            func(r rune) bool { return r == ' ' },
				preserveFunc: []func(rune) bool{},
			},
			want: []string{"hello"},
		},
		{
			name: "multiple words",
			args: args{
				s:            "hello world",
				f:            func(r rune) bool { return r == ' ' },
				preserveFunc: []func(rune) bool{},
			},
			want: []string{"hello", "world"},
		},
		{
			name: "multiple spaces",
			args: args{
				s:            "hello   world",
				f:            func(r rune) bool { return r == ' ' },
				preserveFunc: []func(rune) bool{},
			},
			want: []string{"hello", "world"},
		},
		{
			name: "preserve single char",
			args: args{
				s:            "hello-world",
				f:            func(r rune) bool { return r == ' ' },
				preserveFunc: []func(rune) bool{func(r rune) bool { return r == '-' }},
			},
			want: []string{"hello", "-world"},
		},
		{
			name: "preserve multiple chars",
			args: args{
				s:            "hello-world_again",
				f:            func(r rune) bool { return r == ' ' },
				preserveFunc: []func(rune) bool{func(r rune) bool { return r == '-' }, func(r rune) bool { return r == '_' }},
			},
			want: []string{"hello", "-world", "_again"},
		},
		{
			name: "preserve multiple chars and spaces",
			args: args{
				s:            "hello - world _ again",
				f:            func(r rune) bool { return r == ' ' },
				preserveFunc: []func(rune) bool{func(r rune) bool { return r == '-' }, func(r rune) bool { return r == '_' }},
			},
			want: []string{"hello", "-", "world", "_", "again"},
		},
		{
			name: "preserve multiple chars and spaces at start and end",
			args: args{
				s:            "- hello - world _ again _",
				f:            func(r rune) bool { return r == ' ' },
				preserveFunc: []func(rune) bool{func(r rune) bool { return r == '-' }, func(r rune) bool { return r == '_' }},
			},
			want: []string{"-", "hello", "-", "world", "_", "again", "_"},
		},
		{
			name: "custom delimiter",
			args: args{
				s:            "hello,world,again",
				f:            func(r rune) bool { return r == ',' },
				preserveFunc: []func(rune) bool{},
			},
			want: []string{"hello", "world", "again"},
		},
		{
			name: "custom delimiter and preserve",
			args: args{
				s:            "hello,world-again,test",
				f:            func(r rune) bool { return r == ',' },
				preserveFunc: []func(rune) bool{func(r rune) bool { return r == '-' }},
			},
			want: []string{"hello", "world", "-again", "test"},
		},
		{
			name: "unicode",
			args: args{
				s:            "привет мир",
				f:            func(r rune) bool { return r == ' ' },
				preserveFunc: []func(rune) bool{},
			},
			want: []string{"привет", "мир"},
		},
		{
			name: "unicode and preserve",
			args: args{
				s:            "привет-мир",
				f:            func(r rune) bool { return r == ' ' },
				preserveFunc: []func(rune) bool{func(r rune) bool { return r == '-' }},
			},
			want: []string{"привет", "-мир"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fieldsFunc(tt.args.s, tt.args.f, tt.args.preserveFunc...); !arr.Equals(got, tt.want) {
				t.Errorf("fieldsFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
