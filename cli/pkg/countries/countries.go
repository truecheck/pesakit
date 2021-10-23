//

package countries

import (
	"fmt"
	"strings"
)

const (
	UGANDA                          = "UGANDA"
	NIGERIA                         = "NIGERIA"
	TANZANIA                        = "TANZANIA"
	KENYA                           = "KENYA"
	RWANDA                          = "RWANDA"
	ZAMBIA                          = "ZAMBIA"
	GABON                           = "GABON"
	NIGER                           = "NIGER"
	CONGO_BRAZZAVILLE               = "CONGO-BRAZZAVILLE"
	DR_CONGO                        = "DR CONGO"
	CHAD                            = "CHAD"
	SEYCHELLES                      = "SEYCHELLES"
	MADAGASCAR                      = "MADAGASCAR"
	MALAWI                          = "MALAWI"
	UGANDA_CODE                     = "UG"
	NIGERIA_CODE                    = "NG"
	TANZANIA_CODE                   = "TZ"
	KENYA_CODE                      = "KE"
	RWANDA_CODE                     = "RW"
	ZAMBIA_CODE                     = "ZM"
	GABON_CODE                      = "GA"
	NIGER_CODE                      = "NE"
	CONGO_BRAZZAVILLE_CODE          = "CG"
	DR_CONGO_CODE                   = "CD"
	CHAD_CODE                       = "CFA"
	SEYCHELLES_CODE                 = "SC"
	MADAGASCAR_CODE                 = "MG"
	MALAWI_CODE                     = "MW"
	UGANDA_CURRENCY_CODE            = "UGX"
	NIGERIA_CURRENCY_CODE           = "NGN"
	TANZANIA_CURRENCY_CODE          = "TZS"
	KENYA_CURRENCY_CODE             = "KES"
	RWANDA_CURRENCY_CODE            = "RWF"
	ZAMBIA_CURRENCY_CODE            = "ZMW"
	GABON_CURRENCY_CODE             = "CFA"
	NIGER_CURRENCY_CODE             = "XOF"
	CONGO_BRAZZAVILLE_CURRENCY_CODE = "XAF"
	DR_CONGO_CURRENCY_CODE          = "CDF"
	CHAD_CURRENCY_CODE              = "XAF"
	SEYCHELLES_CURRENCY_CODE        = "SCR"
	MADAGASCAR_CURRENCY_CODE        = "MGA"
	MALAWI_CURRENCY_CODE            = "MWK"
	UGANDA_CURRENCY                 = "Ugandan shilling"
	NIGERIA_CURRENCY                = "Nigerian naira"
	TANZANIA_CURRENCY               = "Tanzanian shilling"
	KENYA_CURRENCY                  = "Kenyan shilling"
	RWANDA_CURRENCY                 = "Rwandan franc"
	ZAMBIA_CURRENCY                 = "Zambian kwacha"
	GABON_CURRENCY                  = "CFA franc BEAC"
	NIGER_CURRENCY                  = "CFA franc BCEAO"
	CONGO_BRAZZAVILLE_CURRENCY      = "CFA franc BCEA"
	DR_CONGO_CURRENCY               = "Congolese franc"
	CHAD_CURRENCY                   = "CFA franc BEAC"
	SEYCHELLES_CURRENCY             = "Seychelles rupee"
	MADAGASCAR_CURRENCY             = "Malagasy ariary"
	MALAWI_CURRENCY                 = "Malawian Kwacha"
)

type (
	Country struct {
		Name         string
		CodeName     string
		Currency     string
		CurrencyCode string
	}
)

func Names() []string {
	names := []string{
		UGANDA, NIGER, NIGERIA, TANZANIA, KENYA, RWANDA,
		ZAMBIA, GABON, CONGO_BRAZZAVILLE, DR_CONGO, CHAD,
		SEYCHELLES, MADAGASCAR, MALAWI,
	}

	return names
}

func GetByName(name string) (Country, error) {
	name = strings.TrimSpace(strings.ToLower(name))
	countries := List()
	for _, country := range countries {
		n := strings.TrimSpace(strings.ToLower(country.Name))
		if name == n {
			return country, nil
		}
	}

	return Country{}, fmt.Errorf("error: the country %s is not supported", name)
}

func List() []Country {

	var countries []Country
	var (
		uganda = Country{
			Name:         UGANDA,
			CodeName:     UGANDA_CODE,
			Currency:     UGANDA_CURRENCY,
			CurrencyCode: UGANDA_CURRENCY_CODE,
		}

		sych = Country{
			Name:         SEYCHELLES,
			CodeName:     SEYCHELLES_CODE,
			Currency:     SEYCHELLES_CURRENCY,
			CurrencyCode: SEYCHELLES_CURRENCY_CODE,
		}

		brazzaville = Country{
			Name:         CONGO_BRAZZAVILLE,
			CodeName:     CONGO_BRAZZAVILLE_CODE,
			Currency:     CONGO_BRAZZAVILLE_CURRENCY,
			CurrencyCode: CONGO_BRAZZAVILLE_CURRENCY_CODE,
		}

		kenya = Country{
			Name:         KENYA,
			CodeName:     KENYA_CODE,
			Currency:     KENYA_CURRENCY,
			CurrencyCode: KENYA_CURRENCY_CODE,
		}

		nigeria = Country{
			Name:         NIGERIA,
			CodeName:     NIGERIA_CODE,
			Currency:     NIGERIA_CURRENCY,
			CurrencyCode: NIGERIA_CURRENCY_CODE,
		}

		rwanda = Country{
			Name:         RWANDA,
			CodeName:     RWANDA_CODE,
			Currency:     RWANDA_CURRENCY,
			CurrencyCode: RWANDA_CURRENCY_CODE,
		}

		niger = Country{
			Name:         NIGER,
			CodeName:     NIGER_CODE,
			Currency:     NIGER_CURRENCY,
			CurrencyCode: NIGER_CURRENCY_CODE,
		}

		chad = Country{
			Name:         CHAD,
			CodeName:     CHAD_CODE,
			Currency:     CHAD_CURRENCY,
			CurrencyCode: CHAD_CURRENCY_CODE,
		}

		congo = Country{
			Name:         DR_CONGO,
			CodeName:     DR_CONGO_CODE,
			Currency:     DR_CONGO_CURRENCY,
			CurrencyCode: DR_CONGO_CURRENCY_CODE,
		}

		madagascar = Country{
			Name:         MADAGASCAR,
			CodeName:     MADAGASCAR_CODE,
			Currency:     MADAGASCAR_CURRENCY,
			CurrencyCode: MADAGASCAR_CURRENCY_CODE,
		}

		zambia = Country{
			Name:         ZAMBIA,
			CodeName:     ZAMBIA_CODE,
			Currency:     ZAMBIA_CURRENCY,
			CurrencyCode: ZAMBIA_CURRENCY_CODE,
		}

		gabon = Country{
			Name:         GABON,
			CodeName:     GABON_CODE,
			Currency:     GABON_CURRENCY,
			CurrencyCode: GABON_CURRENCY_CODE,
		}
		tz = Country{
			Name:         TANZANIA,
			CodeName:     TANZANIA_CODE,
			Currency:     TANZANIA_CURRENCY,
			CurrencyCode: TANZANIA_CURRENCY_CODE,
		}
		malawi = Country{
			Name:         MALAWI,
			CodeName:     MALAWI_CODE,
			Currency:     MALAWI_CURRENCY,
			CurrencyCode: MALAWI_CURRENCY_CODE,
		}
	)

	countries = append(countries, uganda, malawi, tz, gabon, congo,
		brazzaville, rwanda, kenya, madagascar, zambia, sych, chad, niger, nigeria)

	return countries
}
