package main

import (
	"fmt"
)
var claimes = map[string]interface {}{"citizen_id":"784199092345671", "city_id":0, "city_name":"", "country_code":"ae", "country_id":224, "dob":"1990-09-28", "dob_format":"gregorian", "email":"teuku.z.amin@gmail.com", "expired":"17/06/2022 22:33:12", "fg_payment":0, "gender":"male", "identity_id":"784199092345671", "identity_type_id":1, "language_code":"", "latitude":"24.7241504", "longitude":"46.2620616", "mobile_number":"+628161964775", "name":"Teuku Zulfikar Amin", "network_id":"1", "thirdparty_patient_id":"168873495", "thirdparty_user_id":"168873495", "user_id":0, "whitelabel_id":21}, TokenType:"bearer_token", StandardClaims:jwt.StandardClaims{Audience:"505091", ExpiresAt:1655490793, Id:"", IssuedAt:1627442714, Issuer:"okadoc.com", NotBefore:0, Subject:"teuku.z.amin@gmail.com"}

type Decorator interface {
	Decorate() string
}

type Banner struct {
	str string
}

func (self *Banner) getString() string {
	return "*" + self.str + "*"
}

// 構造体の埋込による継承
type EmbeddedDecorateBanner struct {
	*Banner
}

// 埋込による継承はgetは構造体の階層を意識しなくてもよいが
// setは明示的に埋め込んだ構造体に明示的に値を定義する必要があるため、
// 呼び出し側が階層を意識しなくてもよいようにラップするなど工夫が必要
func NewEmbeddedDecorateBanner(str string) *EmbeddedDecorateBanner {
	return &EmbeddedDecorateBanner{&Banner{str}}
}

// インターフェースの実装と埋め込んだ構造体のメソッドによるアダプタ
func (self *EmbeddedDecorateBanner) Decorate() string {
	return self.getString()
}

func main() {
	var d Decorator;
	d = NewEmbeddedDecorateBanner("abu")

	s := d.Decorate()
	fmt.Println(s)
}