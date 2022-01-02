package pesakit

import "testing"

func Test_copyToDir(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name     string
		args     args
		wantPath string
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				src: "pesakit.env",
				dst: "/home/masengwa/.pesakit",
			},
			wantErr:  false,
			wantPath: "/home/masengwa/.pesakit/pesakit.env",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, err := copyToDir(tt.args.src, tt.args.dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("copyToDir() error = %v, wantErr %v", err, tt.wantErr)
			}
			if path != tt.wantPath {
				t.Errorf("copyToDir() path = %v, want %v", path, tt.wantPath)
			}
		})
	}
}
