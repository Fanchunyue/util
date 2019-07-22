package util

import (
	"testing"
)

func TestDeCompressZip(t *testing.T) {
	type args struct {
		filepath string
		dir      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "测试解压",
			args: args{
				dir:      "C:\\demozip\\",
				filepath: "D:\\plemis_v3.zip",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeCompressZip(tt.args.filepath, tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("DeCompressZip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCompressZip(t *testing.T) {
	type args struct {
		dir      string
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "测试压缩",
			args: args{
				dir:      "D:\\plemis_v3_190617",
				filepath: "D:\\p3.0.zip",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CompressZip(tt.args.dir, tt.args.filepath); (err != nil) != tt.wantErr {
				t.Errorf("CompressZip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
