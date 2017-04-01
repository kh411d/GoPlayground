package main

import (
    "fmt"
    "reflect"
    "strings"
    "unicode"
)

const (
    Email          string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
    CreditCard     string = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"
    ISBN10         string = "^(?:[0-9]{9}X|[0-9]{10})$"
    ISBN13         string = "^(?:[0-9]{13})$"
    UUID3          string = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
    UUID4          string = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
    UUID5          string = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
    UUID           string = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
    Alpha          string = "^[a-zA-Z]+$"
    Alphanumeric   string = "^[a-zA-Z0-9]+$"
    Numeric        string = "^[-+]?[0-9]+$"
    Int            string = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
    Float          string = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
    Hexadecimal    string = "^[0-9a-fA-F]+$"
    Hexcolor       string = "^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
    RGBcolor       string = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
    ASCII          string = "^[\x00-\x7F]+$"
    Multibyte      string = "[^\x00-\x7F]"
    FullWidth      string = "[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
    HalfWidth      string = "[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
    Base64         string = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
    PrintableASCII string = "^[\x20-\x7E]+$"
    DataURI        string = "^data:.+\\/(.+);base64$"
    Latitude       string = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
    Longitude      string = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
    DNSName        string = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9]{1}[a-zA-Z0-9_-]{1,62})*$`
    IP             string = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
    URLSchema      string = `((ftp|tcp|udp|wss?|https?):\/\/)`
    URLUsername    string = `(\S+(:\S*)?@)`
    Hostname       string = ``
    URLPath        string = `((\/|\?|#)[^\s]*)`
    URLPort        string = `(:(\d{1,5}))`
    URLIP          string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))`
    URLSubdomain   string = `((www\.)|([a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*))`
    URL            string = `^` + URLSchema + `?` + URLUsername + `?` + `((` + URLIP + `|(\[` + IP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))` + URLPort + `?` + URLPath + `?$`
    SSN            string = `^\d{3}[- ]?\d{2}[- ]?\d{4}$`
    WinPath        string = `^[a-zA-Z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*$`
    UnixPath       string = `^(/[^/\x00]*)+/?$`
    Semver         string = "^v?(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"
    tagName        string = "valid"
)

var fieldsRequiredByDefault bool

// Errors is an array of multiple errors and conforms to the error interface.
type Errors []error

type tagOptionsMap map[string]string

// Errors returns itself.
func (es Errors) Errors() []error {
    return es
}

func (es Errors) Error() string {
    var err string
    for _, e := range es {
        err += e.Error() + ";"
    }
    return err
}

// Error encapsulates a name, an error and whether there's a custom error message or not.
type Error struct {
    Name                     string
    Err                      error
    CustomErrorMessageExists bool
}

func (e Error) Error() string {
    if e.CustomErrorMessageExists {
        return e.Err.Error()
    }
    return e.Name + ": " + e.Err.Error()
}

// this struct definition will fail govalidator.ValidateStruct() (and the field values do not matter):
type exampleStruct struct {
    Name  string ``
    Email string `valid:"email"`
}

// lastly, this will only fail when Email is an invalid email address but not when it's empty:
type exampleStruct2 struct {
    Name  string `valid:"required"`
    Email string `valid:"email,optional"`
}

func ValidateStruct(s interface{}) (bool, error) {
    if s == nil {
        return true, nil
    }
    result := true
    var err error
    var errs Errors
    val := reflect.ValueOf(s)
    if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
        val = val.Elem()
    }
    // we only accept structs
    if val.Kind() != reflect.Struct {
        return false, fmt.Errorf("function only accepts structs; got %s", val.Kind())
    }

    //var errs Errors
    for i := 0; i < val.NumField(); i++ {
        valueField := val.Field(i)
        typeField := val.Type().Field(i)

        fmt.Println("ini value field")
        fmt.Println(valueField)
        fmt.Println("Ini TypeField")
        fmt.Println(typeField)

        if typeField.PkgPath != "" {
            continue // Private field
        }
        resultField, err2 := typeCheck(valueField, typeField, val)
        if err2 != nil {
            errs = append(errs, err2)
        }

        result = result && resultField
        fmt.Println(result)
    }
    if len(errs) > 0 {
        err = errs[1]
    }
    return result, err
}

func typeCheck(v reflect.Value, t reflect.StructField, o reflect.Value) (bool, error) {
    if !v.IsValid() {
        return false, nil
    }

    fieldsRequiredByDefault = true

    tag := t.Tag.Get(tagName)

    fmt.Println(tag)

    // Check if the field should be ignored
    switch tag {
    case "":
        if !fieldsRequiredByDefault {
            return true, nil
        }
        return false, Error{t.Name, fmt.Errorf("All fields are required to at least have one validation defined"), false}
    case "-":
        return true, nil
    }

    // options := parseTagIntoMap(tag)
    // var customTypeErrors Errors
    // var customTypeValidatorsExist bool
    // for validatorName, customErrorMessage := range options {
    //     fmt.Println(validatorName)
    // }

    //fmt.Println(v.Kind())
    //fmt.Println(options)

    return true, nil
}

func parseTagIntoMap(tag string) tagOptionsMap {
    optionsMap := make(tagOptionsMap)
    options := strings.SplitN(tag, ",", -1)
    for _, option := range options {
        validationOptions := strings.Split(option, "~")
        if !isValidTag(validationOptions[0]) {
            continue
        }
        if len(validationOptions) == 2 {
            optionsMap[validationOptions[0]] = validationOptions[1]
        } else {
            optionsMap[validationOptions[0]] = ""
        }
    }
    return optionsMap
}

func isValidTag(s string) bool {
    if s == "" {
        return false
    }
    for _, c := range s {
        switch {
        case strings.ContainsRune("!#$%&()*+-./:<=>?@[]^_{|}~ ", c):
            // Backslash and quote chars are reserved, but
            // otherwise any punctuation chars are allowed
            // in a tag name.
        default:
            if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
                return false
            }
        }
    }
    return true
}

func main() {
    x := new(exampleStruct2)
    ValidateStruct(x)
}
