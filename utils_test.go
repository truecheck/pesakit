package pesakit

import "testing"

func Test_copyToDir(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				src: "pesakit.env",
				dst: "$HOME/.pesakit",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := copyToDir(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("copyToDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
