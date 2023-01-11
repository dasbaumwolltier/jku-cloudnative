package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/store/go_cache/v4"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	gocache "github.com/patrickmn/go-cache"
)

const currenciesList = `[{"value":"AFN","label":"Afghan afghani (AFN)"},{"value":"ALL","label":"Albanian lek (ALL)"},{"value":"DZD","label":"Algerian dinar (DZD)"},{"value":"AOA","label":"Angolan kwanza (AOA)"},{"value":"ARS","label":"Argentine peso (ARS)"},{"value":"AMD","label":"Armenian dram (AMD)"},{"value":"AWG","label":"Aruban florin (AWG)"},{"value":"AUD","label":"Australian dollar (AUD)"},{"value":"AZN","label":"Azerbaijani manat (AZN)"},{"value":"BSD","label":"Bahamian dollar (BSD)"},{"value":"BHD","label":"Bahraini dinar (BHD)"},{"value":"BDT","label":"Bangladeshi taka (BDT)"},{"value":"BBD","label":"Barbadian dollar (BBD)"},{"value":"BYR","label":"Belarusian ruble (BYR)"},{"value":"BZD","label":"Belize dollar (BZD)"},{"value":"BMD","label":"Bermudian dollar (BMD)"},{"value":"BTN","label":"Bhutanese ngultrum (BTN)"},{"value":"BOB","label":"Bolivian boliviano (BOB)"},{"value":"BAM","label":"Bosnia and Herzegovina konvertibilna marka (BAM)"},{"value":"BWP","label":"Botswana pula (BWP)"},{"value":"BRL","label":"Brazilian real (BRL)"},{"value":"GBP","label":"British pound (GBP)"},{"value":"BND","label":"Brunei dollar (BND)"},{"value":"BGN","label":"Bulgarian lev (BGN)"},{"value":"BIF","label":"Burundi franc (BIF)"},{"value":"XPF","label":"CFP franc (XPF)"},{"value":"KHR","label":"Cambodian riel (KHR)"},{"value":"CAD","label":"Canadian dollar (CAD)"},{"value":"CVE","label":"Cape Verdean escudo (CVE)"},{"value":"KYD","label":"Cayman Islands dollar (KYD)"},{"value":"GQE","label":"Central African CFA franc (GQE)"},{"value":"XAF","label":"Central African CFA franc (XAF)"},{"value":"CLP","label":"Chilean peso (CLP)"},{"value":"CNY","label":"Chinese/Yuan renminbi (CNY)"},{"value":"COP","label":"Colombian peso (COP)"},{"value":"KMF","label":"Comorian franc (KMF)"},{"value":"CDF","label":"Congolese franc (CDF)"},{"value":"CRC","label":"Costa Rican colon (CRC)"},{"value":"HRK","label":"Croatian kuna (HRK)"},{"value":"CUC","label":"Cuban peso (CUC)"},{"value":"CZK","label":"Czech koruna (CZK)"},{"value":"DKK","label":"Danish krone (DKK)"},{"value":"DJF","label":"Djiboutian franc (DJF)"},{"value":"DOP","label":"Dominican peso (DOP)"},{"value":"XCD","label":"East Caribbean dollar (XCD)"},{"value":"EGP","label":"Egyptian pound (EGP)"},{"value":"ERN","label":"Eritrean nakfa (ERN)"},{"value":"EEK","label":"Estonian kroon (EEK)"},{"value":"ETB","label":"Ethiopian birr (ETB)"},{"value":"EUR","label":"European Euro (EUR)"},{"value":"FKP","label":"Falkland Islands pound (FKP)"},{"value":"FJD","label":"Fijian dollar (FJD)"},{"value":"GMD","label":"Gambian dalasi (GMD)"},{"value":"GEL","label":"Georgian lari (GEL)"},{"value":"GHS","label":"Ghanaian cedi (GHS)"},{"value":"GIP","label":"Gibraltar pound (GIP)"},{"value":"GTQ","label":"Guatemalan quetzal (GTQ)"},{"value":"GNF","label":"Guinean franc (GNF)"},{"value":"GYD","label":"Guyanese dollar (GYD)"},{"value":"HTG","label":"Haitian gourde (HTG)"},{"value":"HNL","label":"Honduran lempira (HNL)"},{"value":"HKD","label":"Hong Kong dollar (HKD)"},{"value":"HUF","label":"Hungarian forint (HUF)"},{"value":"ISK","label":"Icelandic króna (ISK)"},{"value":"INR","label":"Indian rupee (INR)"},{"value":"IDR","label":"Indonesian rupiah (IDR)"},{"value":"IRR","label":"Iranian rial (IRR)"},{"value":"IQD","label":"Iraqi dinar (IQD)"},{"value":"ILS","label":"Israeli new sheqel (ILS)"},{"value":"JMD","label":"Jamaican dollar (JMD)"},{"value":"JPY","label":"Japanese yen (JPY)"},{"value":"JOD","label":"Jordanian dinar (JOD)"},{"value":"KZT","label":"Kazakhstani tenge (KZT)"},{"value":"KES","label":"Kenyan shilling (KES)"},{"value":"KWD","label":"Kuwaiti dinar (KWD)"},{"value":"KGS","label":"Kyrgyzstani som (KGS)"},{"value":"LAK","label":"Lao kip (LAK)"},{"value":"LVL","label":"Latvian lats (LVL)"},{"value":"LBP","label":"Lebanese lira (LBP)"},{"value":"LSL","label":"Lesotho loti (LSL)"},{"value":"LRD","label":"Liberian dollar (LRD)"},{"value":"LYD","label":"Libyan dinar (LYD)"},{"value":"LTL","label":"Lithuanian litas (LTL)"},{"value":"MOP","label":"Macanese pataca (MOP)"},{"value":"MKD","label":"Macedonian denar (MKD)"},{"value":"MGA","label":"Malagasy ariary (MGA)"},{"value":"MWK","label":"Malawian kwacha (MWK)"},{"value":"MYR","label":"Malaysian ringgit (MYR)"},{"value":"MVR","label":"Maldivian rufiyaa (MVR)"},{"value":"MRO","label":"Mauritanian ouguiya (MRO)"},{"value":"MUR","label":"Mauritian rupee (MUR)"},{"value":"MXN","label":"Mexican peso (MXN)"},{"value":"MDL","label":"Moldovan leu (MDL)"},{"value":"MNT","label":"Mongolian tugrik (MNT)"},{"value":"MAD","label":"Moroccan dirham (MAD)"},{"value":"MZM","label":"Mozambican metical (MZM)"},{"value":"MMK","label":"Myanma kyat (MMK)"},{"value":"NAD","label":"Namibian dollar (NAD)"},{"value":"NPR","label":"Nepalese rupee (NPR)"},{"value":"ANG","label":"Netherlands Antillean gulden (ANG)"},{"value":"TWD","label":"New Taiwan dollar (TWD)"},{"value":"NZD","label":"New Zealand dollar (NZD)"},{"value":"NIO","label":"Nicaraguan córdoba (NIO)"},{"value":"NGN","label":"Nigerian naira (NGN)"},{"value":"KPW","label":"North Korean won (KPW)"},{"value":"NOK","label":"Norwegian krone (NOK)"},{"value":"OMR","label":"Omani rial (OMR)"},{"value":"PKR","label":"Pakistani rupee (PKR)"},{"value":"PAB","label":"Panamanian balboa (PAB)"},{"value":"PGK","label":"Papua New Guinean kina (PGK)"},{"value":"PYG","label":"Paraguayan guarani (PYG)"},{"value":"PEN","label":"Peruvian nuevo sol (PEN)"},{"value":"PHP","label":"Philippine peso (PHP)"},{"value":"PLN","label":"Polish zloty (PLN)"},{"value":"QAR","label":"Qatari riyal (QAR)"},{"value":"RON","label":"Romanian leu (RON)"},{"value":"RUB","label":"Russian ruble (RUB)"},{"value":"SHP","label":"Saint Helena pound (SHP)"},{"value":"WST","label":"Samoan tala (WST)"},{"value":"SAR","label":"Saudi riyal (SAR)"},{"value":"RSD","label":"Serbian dinar (RSD)"},{"value":"SCR","label":"Seychellois rupee (SCR)"},{"value":"SLL","label":"Sierra Leonean leone (SLL)"},{"value":"SGD","label":"Singapore dollar (SGD)"},{"value":"SBD","label":"Solomon Islands dollar (SBD)"},{"value":"SOS","label":"Somali shilling (SOS)"},{"value":"ZAR","label":"South African rand (ZAR)"},{"value":"KRW","label":"South Korean won (KRW)"},{"value":"XDR","label":"Special Drawing Rights (XDR)"},{"value":"LKR","label":"Sri Lankan rupee (LKR)"},{"value":"SDG","label":"Sudanese pound (SDG)"},{"value":"SRD","label":"Surinamese dollar (SRD)"},{"value":"SZL","label":"Swazi lilangeni (SZL)"},{"value":"SEK","label":"Swedish krona (SEK)"},{"value":"CHF","label":"Swiss franc (CHF)"},{"value":"SYP","label":"Syrian pound (SYP)"},{"value":"TJS","label":"Tajikistani somoni (TJS)"},{"value":"TZS","label":"Tanzanian shilling (TZS)"},{"value":"THB","label":"Thai baht (THB)"},{"value":"TTD","label":"Trinidad and Tobago dollar (TTD)"},{"value":"TND","label":"Tunisian dinar (TND)"},{"value":"TRY","label":"Turkish new lira (TRY)"},{"value":"TMT","label":"Turkmen manat (TMT)"},{"value":"AED","label":"UAE dirham (AED)"},{"value":"UGX","label":"Ugandan shilling (UGX)"},{"value":"UAH","label":"Ukrainian hryvnia (UAH)"},{"value":"USD","label":"United States dollar (USD)"},{"value":"UYU","label":"Uruguayan peso (UYU)"},{"value":"UZS","label":"Uzbekistani som (UZS)"},{"value":"VUV","label":"Vanuatu vatu (VUV)"},{"value":"VEB","label":"Venezuelan bolivar (VEB)"},{"value":"VND","label":"Vietnamese dong (VND)"},{"value":"XOF","label":"West African CFA franc (XOF)"},{"value":"YER","label":"Yemeni rial (YER)"},{"value":"ZMK","label":"Zambian kwacha (ZMK)"},{"value":"ZWR","label":"Zimbabwean dollar (ZWR)"}]`

func main() {
	r := gin.Default()

	client := gocache.New(5*time.Minute, 10*time.Minute)
	store := go_cache.NewGoCache(client)

	cache := cache.New[*string](store)

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	r.GET("/api/currencies", func(ctx *gin.Context) {
		currencies(ctx)
	})

	r.GET("/api/convert", func(ctx *gin.Context) {
		convert(ctx, cache)
	})

	r.Run()
}

func currencies(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "application/json", []byte(currenciesList))
	return
}

func convert(ctx *gin.Context, cache *cache.Cache[*string]) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	val := ctx.Query("val")

	if from == "" || to == "" || val == "" {
		ctx.Status(400)
		return
	}

	f64, err := strconv.ParseFloat(val, 64)

	cachKey := fmt.Sprintf("convert:%s:%s", from, to)
	cached, err := cache.Get(ctx, cachKey)

	res := 0.0
	if err == nil {
		res, err = strconv.ParseFloat(*cached, 64)

		if err != nil {
			panic("invalid cache")
		}
	} else {
		res = fetchConversion(ctx, &from, &to)

		if res == 0 {
			ctx.Status(400)
			return
		}

		cacheVal := fmt.Sprintf("%f", res)
		cache.Set(ctx, cachKey, &cacheVal)
	}

	resres := f64 * res

	ctx.Data(200, "application/json", []byte(fmt.Sprintf("{\"value\": %f}", resres)))
}

func fetchCurrencies(ctx *gin.Context) string {
	res, err := http.Get("https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies.min.json")

	if err != nil {
		return "[]"
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "[]"
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}

func fetchConversion(ctx *gin.Context, from *string, to *string) float64 {
	res, err := http.Get(fmt.Sprintf("https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies/%s/%s.min.json", *from, *to))

	if err != nil {
		return 0
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return 0
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	parsed := map[string]float64{}
	json.Unmarshal(body, &parsed)

	return parsed[*to]
}
