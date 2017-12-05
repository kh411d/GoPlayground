package main

import (
	"encoding/json"
	"fmt"
)

type xx struct {
	ProductID       string                     `json:"product_id"`
	ProductVariants map[int64]PulsaVariantInfo `json:"product_variants"`
}
type PulsaVariantInfo struct {
	EffectivePrice string `json:"effective_price"`
	PriceTag       string `json:"price_string"`
	PulsaValue     string `json:"pulsa_value"`
	PulsaVoucher   string `json:"pulsa_voucher"`
	VariantID      string `json:"variant_id"`
	VariantSKU     string `json:"variant_sku"`
}

func main() {

	ss := `{
    "5153": {
        "1157164": {
        	"product_id": "1157164",
            "product_variants": {
                "1541591": {
                    "effective_price": "24984.00",
                    "price_string": "Rp. 24.984",
                    "pulsa_value": "25,000 ",
                    "pulsa_voucher": " OXL25",
                    "variant_id": "1541591",
                    "variant_sku": "PNA0051530000000101"
                },
                "1541592": {
                    "effective_price": "49968.00",
                    "price_string": "Rp. 49.968",
                    "pulsa_value": "50,000 ",
                    "pulsa_voucher": " OXL50",
                    "variant_id": "1541592",
                    "variant_sku": "PNA0051530000000102"
                },
                "1541593": {
                    "effective_price": "99937.00",
                    "price_string": "Rp. 99.937",
                    "pulsa_value": "100,000 ",
                    "pulsa_voucher": " OXL100",
                    "variant_id": "1541593",
                    "variant_sku": "PNA0051530000000103"
                },
                "1541594": {
                    "effective_price": "199875.00",
                    "price_string": "Rp. 199.875",
                    "pulsa_value": "200,000 ",
                    "pulsa_voucher": " OXL200",
                    "variant_id": "1541594",
                    "variant_sku": "PNA0051530000000104"
                },
                "1541595": {
                    "effective_price": "299812.00",
                    "price_string": "Rp. 299.812",
                    "pulsa_value": "300,000 ",
                    "pulsa_voucher": " OXL300",
                    "variant_id": "1541595",
                    "variant_sku": "PNA0051530000000105"
                },
                "1541596": {
                    "effective_price": "499687.00",
                    "price_string": "Rp. 499.687",
                    "pulsa_value": "500,000 ",
                    "pulsa_voucher": " OXL500",
                    "variant_id": "1541596",
                    "variant_sku": "PNA0051530000000106"
                },
                "1541597": {
                    "effective_price": "999375.00",
                    "price_string": "Rp. 999.375",
                    "pulsa_value": "1,000,000 ",
                    "pulsa_voucher": " OXL1000",
                    "variant_id": "1541597",
                    "variant_sku": "PNA0051530000000107"
                }
            }
        },
        "1157185": {
        	"product_id": "1157185",
            "product_variants": {
                "1541621": {
                    "effective_price": "25498.00",
                    "price_string": "Rp. 25.498",
                    "pulsa_value": "25,000 ",
                    "pulsa_voucher": " OIS25",
                    "variant_id": "1541621",
                    "variant_sku": "PNA0051530000000201"
                },
                "1541622": {
                    "effective_price": "50737.00",
                    "price_string": "Rp. 50.737",
                    "pulsa_value": "50,000 ",
                    "pulsa_voucher": " OIS50",
                    "variant_id": "1541622",
                    "variant_sku": "PNA0051530000000202"
                },
                "1541623": {
                    "effective_price": "99937.00",
                    "price_string": "Rp. 99.937",
                    "pulsa_value": "100,000 ",
                    "pulsa_voucher": " OIS100",
                    "variant_id": "1541623",
                    "variant_sku": "PNA0051530000000203"
                },
                "1541624": {
                    "effective_price": "148368.00",
                    "price_string": "Rp. 148.368",
                    "pulsa_value": "150,000 ",
                    "pulsa_voucher": " OIS150",
                    "variant_id": "1541624",
                    "variant_sku": "PNA0051530000000204"
                },
                "1541625": {
                    "effective_price": "243437.00",
                    "price_string": "Rp. 243.437",
                    "pulsa_value": "250,000 ",
                    "pulsa_voucher": " OIS250",
                    "variant_id": "1541625",
                    "variant_sku": "PNA0051530000000205"
                },
                "1541626": {
                    "effective_price": "486875.00",
                    "price_string": "Rp. 486.875",
                    "pulsa_value": "500,000 ",
                    "pulsa_voucher": " OIS500",
                    "variant_id": "1541626",
                    "variant_sku": "PNA0051530000000206"
                },
                "1541627": {
                    "effective_price": "973750.00",
                    "price_string": "Rp. 973.750",
                    "pulsa_value": "1,000,000 ",
                    "pulsa_voucher": " OIS1000",
                    "variant_id": "1541627",
                    "variant_sku": "PNA0051530000000207"
                },
                "1606251": {
                    "effective_price": "96401.00",
                    "price_string": "Rp. 96.401",
                    "pulsa_value": "6GB ",
                    "pulsa_voucher": " OIDL",
                    "variant_id": "1606251",
                    "variant_sku": "PNA0051530000000210"
                },
                "1606546": {
                    "effective_price": "145068.00",
                    "price_string": "Rp. 145.068",
                    "pulsa_value": "10GB ",
                    "pulsa_voucher": " OIDXL",
                    "variant_id": "1606546",
                    "variant_sku": "PNA0051530000000212"
                },
                "1606548": {
                    "effective_price": "193776.00",
                    "price_string": "Rp. 193.776",
                    "pulsa_value": "20GB ",
                    "pulsa_voucher": " OIDXXL",
                    "variant_id": "1606548",
                    "variant_sku": "PNA0051530000000214"
                }
            }
        },
        "1157211": {
        	"product_id": "1157211",
            "product_variants": {
                "1541668": {
                    "effective_price": "25317.00",
                    "price_string": "Rp. 25.317",
                    "pulsa_value": "25,000 ",
                    "pulsa_voucher": " OTS25",
                    "variant_id": "1541668",
                    "variant_sku": "PNA0051530000000301"
                },
                "1541669": {
                    "effective_price": "49712.00",
                    "price_string": "Rp. 49.712",
                    "pulsa_value": "50,000 ",
                    "pulsa_voucher": " OTS50",
                    "variant_id": "1541669",
                    "variant_sku": "PNA0051530000000302"
                },
                "1541670": {
                    "effective_price": "98912.00",
                    "price_string": "Rp. 98.912",
                    "pulsa_value": "100,000 ",
                    "pulsa_voucher": " OTS100",
                    "variant_id": "1541670",
                    "variant_sku": "PNA0051530000000303"
                },
                "1541671": {
                    "effective_price": "148368.00",
                    "price_string": "Rp. 148.368",
                    "pulsa_value": "150,000 ",
                    "pulsa_voucher": " OTS150",
                    "variant_id": "1541671",
                    "variant_sku": "PNA0051530000000304"
                },
                "1541672": {
                    "effective_price": "197825.00",
                    "price_string": "Rp. 197.825",
                    "pulsa_value": "200,000 ",
                    "pulsa_voucher": " OTS200",
                    "variant_id": "1541672",
                    "variant_sku": "PNA0051530000000305"
                },
                "1541673": {
                    "effective_price": "296737.00",
                    "price_string": "Rp. 296.737",
                    "pulsa_value": "300,000 ",
                    "pulsa_voucher": " OTS300",
                    "variant_id": "1541673",
                    "variant_sku": "PNA0051530000000306"
                },
                "1579116": {
                    "effective_price": "494562.00",
                    "price_string": "Rp. 494.562",
                    "pulsa_value": "500,000 ",
                    "pulsa_voucher": " OTS500",
                    "variant_id": "1579116",
                    "variant_sku": "PNA0051530000000309"
                },
                "1579117": {
                    "effective_price": "989125.00",
                    "price_string": "Rp. 989.125",
                    "pulsa_value": "1,000,000 ",
                    "pulsa_voucher": " OTS1000",
                    "variant_id": "1579117",
                    "variant_sku": "PNA0051530000000310"
                },
                "2106872": {
                    "effective_price": "21000.00",
                    "price_string": "Rp. 21.000",
                    "pulsa_value": "20,000 ",
                    "pulsa_voucher": " OTS20",
                    "variant_id": "2106872",
                    "variant_sku": "PNA0051530000000311"
                }
            }
        }
    }
}`

	var x map[string]map[string]xx
	er := json.Unmarshal([]byte(ss), &x)
	fmt.Println(er)

	var r []xx
	for _, v := range x["5153"] {
		r = append(r, v)
	}

	fmt.Printf("%v\n", r)
}
