package player

import "testing"

func Test_runMPG123(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "play song",
			args: args{
				url: "http://ip.h5.ra01.sycdn.kuwo.cn/a7ffd0949b39c05098e5c0dfc09d59d6/5e22e41f/resource/n1/320/87/73/2227989589.mp3",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runMPG123(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("runMPG123() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
