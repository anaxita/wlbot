package helpers

import "testing"

func TestFindIP4(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name         string
		args         args
		wantIp4      string
		wantWithCIDR bool
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIp4, gotWithCIDR, err := FindIP4(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindIP4() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIp4 != tt.wantIp4 {
				t.Errorf("FindIP4() gotIp4 = %v, want %v", gotIp4, tt.wantIp4)
			}
			if gotWithCIDR != tt.wantWithCIDR {
				t.Errorf("FindIP4() gotWithCIDR = %v, want %v", gotWithCIDR, tt.wantWithCIDR)
			}
		})
	}
}
