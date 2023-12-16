package resp

import "testing"

func Test_deserializeSimpleString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "simple string with correct format",
			args: args{
				input: "+OK\r\n",
			},
			want:    "OK",
			wantErr: false,
		},
		{
			name: "wrong prefix",
			args: args{
				input: "-OK\r\n",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "empty string",
			args: args{
				input: "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deserializeSimpleString(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("deserializeSimpleString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("deserializeSimpleString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
