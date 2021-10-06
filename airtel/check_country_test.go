//

package airtel

import (
	"testing"
)

func Test_CheckCountry(t *testing.T) {

	all := map[string][]string{
		"collection":   {"Tanzania", "Kenya", "Uganda", "Rwanda"},
		"disbursement": {"Tanzania"},
		"account":      {"Rwanda", "Uganda"},
	}
	type args struct {
		api          string
		country      string
		allCountries map[string][]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check if not allowed is false",
			args: args{
				api:          CollectionApiGroup,
				country:      "Burundi",
				allCountries: all,
			},
			want: false,
		},
		{
			name: "checking if it returns true when the country is allowed",
			args: args{
				api:          CollectionApiGroup,
				country:      "tanzania",
				allCountries: all,
			},
			want: true,
		},
		{
			name: "passing empty collection name",
			args: args{
				api:          "",
				country:      "burundi",
				allCountries: all,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckCountry(tt.args.api, tt.args.country, tt.args.allCountries); got != tt.want {
				t.Errorf("CheckCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}
