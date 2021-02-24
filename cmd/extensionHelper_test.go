package cmd

import "testing"

func Test_genNewFilename(t *testing.T) {
	type args struct {
		filename  string
		extension string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"basic", args{"a.out", "zip"}, "a.zip"},
		{"basic dot", args{"a.out", ".zip"}, "a.zip"},
		{"too short", args{"a.ou", "zip"}, "a.ou.zip"},
		{"compound", args{"a.tar.gz", "zip"}, "a.tar.gz.zip"},
		{"dotted", args{"a..tz", "zip"}, "a..tz.zip"},
		{"too short 2", args{".out", "zip"}, ".out.zip"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genNewFilename(tt.args.filename, tt.args.extension); got != tt.want {
				t.Errorf("genNewFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
