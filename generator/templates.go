package generator

import (
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/facebook/ent/schema/field"
	"github.com/go-openapi/inflect"
	"github.com/sijad/entgql/generator/internal"
)

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -o=internal/bindata.go -pkg=internal -modtime=1 ./template/...

func init() {
	templates.Funcs(funcMap)
	for _, asset := range internal.AssetNames() {
		templates = template.Must(templates.New(asset).Parse(string(internal.MustAsset(asset))))
	}
}

var templates = template.New("templates")

var (
	funcMap = template.FuncMap{
		"lowerCase":      strings.ToLower,
		"upperCase":      strings.ToUpper,
		"lowerCaseFirst": lowerCaseFirst,
		"upperCaseFirst": upperCaseFirst,
		"singular":       rules.Singularize,
		"pluralize":      rules.Pluralize,
		"camelize":       rules.Camelize,
		"scalarName":     scalarName,
		"goType":         goType,
		"quote":          strconv.Quote,
	}
	rules    = ruleset()
	acronyms = make(map[string]struct{})
)

func ruleset() *inflect.Ruleset {
	rules := inflect.NewDefaultRuleset()
	// Add common initialisms from golint and more.
	for _, w := range []string{
		"ACL", "API", "ASCII", "AWS", "CPU", "CSS", "DNS", "EOF", "GB", "GUID",
		"HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "KB", "LHS", "MAC", "MB",
		"QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "SSO", "TCP",
		"TLS", "TTL", "UDP", "UI", "UID", "URI", "URL", "UTF8", "UUID", "VM",
		"XML", "XMPP", "XSRF", "XSS",
	} {
		acronyms[w] = struct{}{}
		rules.AddAcronym(w)
	}
	return rules
}

func lowerCaseFirst(s string) string {
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func upperCaseFirst(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func scalarName(typ *field.TypeInfo) (t string) {
	switch typ.Type {
	case field.TypeBool:
		t = "Boolean"
	case field.TypeTime:
		t = "Time"
	case field.TypeJSON:
		t = "Any"
	case field.TypeUUID:
		t = "UUID"
	case field.TypeString:
		t = "String"
	case field.TypeInt, field.TypeInt8, field.TypeInt16, field.TypeInt32, field.TypeInt64:
		t = "Int"
	case field.TypeUint, field.TypeUint8, field.TypeUint16, field.TypeUint32, field.TypeUint64:
		t = "Uint"
	case field.TypeFloat32, field.TypeFloat64:
		t = "Float"
	default:
		panic(fmt.Sprintf("Type %s is not supported", typ))
	}
	return t
}

func goType(typ *field.TypeInfo) (t string) {
	if typ.Type == field.TypeInvalid {
		panic(fmt.Sprintf("Type %s is not supported", typ))
	}
	return typ.Type.String()
}
