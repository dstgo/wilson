// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package is provides a list of commonly used string vax rules.
package is

import (
	"regexp"
	"unicode"

	"github.com/asaskevich/govalidator"
	"github.com/dstgo/wilson/app/core/vax"
)

var (
	// ErrEmail is the error that returns in case of an invalid email.
	ErrEmail = vax.NewError("validate.string.email", "must be a valid email address")
	// ErrURL is the error that returns in case of an invalid URL.
	ErrURL = vax.NewError("validate.string.url", "must be a valid URL")
	// ErrRequestURL is the error that returns in case of an invalid request URL.
	ErrRequestURL = vax.NewError("validate.string.requrl", "must be a valid request URL")
	// ErrRequestURI is the error that returns in case of an invalid request URI.
	ErrRequestURI = vax.NewError("validate.string.requri", "must be a valid request URI")
	// ErrAlpha is the error that returns in case of an invalid alpha value.
	ErrAlpha = vax.NewError("validate.string.alpha", "must contain English letters only")
	// ErrDigit is the error that returns in case of an invalid digit value.
	ErrDigit = vax.NewError("validate.string.digit", "must contain digits only")
	// ErrAlphanumeric is the error that returns in case of an invalid alphanumeric value.
	ErrAlphanumeric = vax.NewError("validate.string.alphanum", "must contain English letters and digits only")
	// ErrUTFLetter is the error that returns in case of an invalid utf letter value.
	ErrUTFLetter = vax.NewError("validate.string.unicode", "must contain unicode letter characters only")
	// ErrUTFDigit is the error that returns in case of an invalid utf digit value.
	ErrUTFDigit = vax.NewError("validate.string.unicodenum", "must contain unicode decimal digits only")
	// ErrUTFLetterNumeric is the error that returns in case of an invalid utf numeric or letter value.
	ErrUTFLetterNumeric = vax.NewError("validate.string.unicodeletternum", "must contain unicode letters and numbers only")
	// ErrUTFNumeric is the error that returns in case of an invalid utf numeric value.
	ErrUTFNumeric = vax.NewError("validate.string.unicodenumch", "must contain unicode number characters only")
	// ErrLowerCase is the error that returns in case of an invalid lower case value.
	ErrLowerCase = vax.NewError("validate.string.lower", "must be in lower case")
	// ErrUpperCase is the error that returns in case of an invalid upper case value.
	ErrUpperCase = vax.NewError("validate.string.upper", "must be in upper case")
	// ErrHexadecimal is the error that returns in case of an invalid hexadecimal number.
	ErrHexadecimal = vax.NewError("validate.string.hex", "must be a valid hexadecimal number")
	// ErrHexColor is the error that returns in case of an invalid hexadecimal color code.
	ErrHexColor = vax.NewError("validate.string.hexcolor", "must be a valid hexadecimal color code")
	// ErrRGBColor is the error that returns in case of an invalid RGB color code.
	ErrRGBColor = vax.NewError("validate.string.rgb", "must be a valid RGB color code")
	// ErrInt is the error that returns in case of an invalid integer value.
	ErrInt = vax.NewError("validate.string.int", "must be an integer number")
	// ErrFloat is the error that returns in case of an invalid float value.
	ErrFloat = vax.NewError("validate.string.float", "must be a floating point number")
	// ErrUUIDv3 is the error that returns in case of an invalid UUIDv3 value.
	ErrUUIDv3 = vax.NewError("validate.string.uuidv3", "must be a valid UUID v3")
	// ErrUUIDv4 is the error that returns in case of an invalid UUIDv4 value.
	ErrUUIDv4 = vax.NewError("validate.string.uuidv4", "must be a valid UUID v4")
	// ErrUUIDv5 is the error that returns in case of an invalid UUIDv5 value.
	ErrUUIDv5 = vax.NewError("validate.string.uuidv5", "must be a valid UUID v5")
	// ErrUUID is the error that returns in case of an invalid UUID value.
	ErrUUID = vax.NewError("validate.string.uuid", "must be a valid UUID")
	// ErrCreditCard is the error that returns in case of an invalid credit card number.
	ErrCreditCard = vax.NewError("validate.string.creditcard", "must be a valid credit card number")
	// ErrISBN10 is the error that returns in case of an invalid ISBN-10 value.
	ErrISBN10 = vax.NewError("validate.string.isbn10", "must be a valid ISBN-10")
	// ErrISBN13 is the error that returns in case of an invalid ISBN-13 value.
	ErrISBN13 = vax.NewError("validate.string.isbn13", "must be a valid ISBN-13")
	// ErrISBN is the error that returns in case of an invalid ISBN value.
	ErrISBN = vax.NewError("validate.string.isbn", "must be a valid ISBN")
	// ErrJSON is the error that returns in case of an invalid JSON.
	ErrJSON = vax.NewError("validate.string.json", "must be in valid JSON format")
	// ErrASCII is the error that returns in case of an invalid ASCII.
	ErrASCII = vax.NewError("validate.string.ascii", "must contain ASCII characters only")
	// ErrPrintableASCII is the error that returns in case of an invalid printable ASCII value.
	ErrPrintableASCII = vax.NewError("validate.string.printable", "must contain printable ASCII characters only")
	// ErrMultibyte is the error that returns in case of an invalid multibyte value.
	ErrMultibyte = vax.NewError("validate.string.multibyte", "must contain multibyte characters")
	// ErrFullWidth is the error that returns in case of an invalid full-width value.
	ErrFullWidth = vax.NewError("validate.string.fullwidth", "must contain full-width characters")
	// ErrHalfWidth is the error that returns in case of an invalid half-width value.
	ErrHalfWidth = vax.NewError("validate.string.halfwidth", "must contain half-width characters")
	// ErrVariableWidth is the error that returns in case of an invalid variable width value.
	ErrVariableWidth = vax.NewError("validate.string.fullhalfwidth", "must contain both full-width and half-width characters")
	// ErrBase64 is the error that returns in case of an invalid base54 value.
	ErrBase64 = vax.NewError("validate.string.base64", "must be encoded in Base64")
	// ErrDataURI is the error that returns in case of an invalid data URI.
	ErrDataURI = vax.NewError("validate.string.base64uri", "must be a Base64-encoded data URI")
	// ErrE164 is the error that returns in case of an invalid e165.
	ErrE164 = vax.NewError("validate.string.e164", "must be a valid E164 number")
	// ErrCountryCode2 is the error that returns in case of an invalid two-letter country code.
	ErrCountryCode2 = vax.NewError("validate.string.2code", "must be a valid two-letter country code")
	// ErrCountryCode3 is the error that returns in case of an invalid three-letter country code.
	ErrCountryCode3 = vax.NewError("validate.string.3code", "must be a valid three-letter country code")
	// ErrCurrencyCode is the error that returns in case of an invalid currency code.
	ErrCurrencyCode = vax.NewError("validate.string.iso4217code", "must be valid ISO 4217 currency code")
	// ErrDialString is the error that returns in case of an invalid string.
	ErrDialString = vax.NewError("validate.string.dail", "must be a valid dial string")
	// ErrMac is the error that returns in case of an invalid mac address.
	ErrMac = vax.NewError("validate.string.mac", "must be a valid MAC address")
	// ErrIP is the error that returns in case of an invalid IP.
	ErrIP = vax.NewError("validate.string.ip", "must be a valid IP address")
	// ErrIPv4 is the error that returns in case of an invalid IPv4.
	ErrIPv4 = vax.NewError("validate.string.ipv4", "must be a valid IPv4 address")
	// ErrIPv6 is the error that returns in case of an invalid IPv6.
	ErrIPv6 = vax.NewError("validate.string.ipv6", "must be a valid IPv6 address")
	// ErrSubdomain is the error that returns in case of an invalid subdomain.
	ErrSubdomain = vax.NewError("validate.string.subdomain", "must be a valid subdomain")
	// ErrDomain is the error that returns in case of an invalid domain.
	ErrDomain = vax.NewError("validate.string.domain", "must be a valid domain")
	// ErrDNSName is the error that returns in case of an invalid DNS name.
	ErrDNSName = vax.NewError("validate.string.dns", "must be a valid DNS name")
	// ErrHost is the error that returns in case of an invalid host.
	ErrHost = vax.NewError("validate.string.host", "must be a valid IP address or DNS name")
	// ErrPort is the error that returns in case of an invalid port.
	ErrPort = vax.NewError("validate.string.port", "must be a valid port number")
	// ErrMongoID is the error that returns in case of an invalid MongoID.
	ErrMongoID = vax.NewError("validate.string.mongoid", "must be a valid hex-encoded MongoDB ObjectId")
	// ErrLatitude is the error that returns in case of an invalid latitude.
	ErrLatitude = vax.NewError("validate.string.latitude", "must be a valid latitude")
	// ErrLongitude is the error that returns in case of an invalid longitude.
	ErrLongitude = vax.NewError("validate.string.longitude", "must be a valid longitude")
	// ErrSSN is the error that returns in case of an invalid SSN.
	ErrSSN = vax.NewError("validate.string.ssn", "must be a valid social security number")
	// ErrSemver is the error that returns in case of an invalid semver.
	ErrSemver = vax.NewError("validate.string.version", "must be a valid semantic version")
)

var (
	// Email validates if a string is an email or not. It also checks if the MX record exists for the email domain.
	Email = vax.String(govalidator.IsExistingEmail, ErrEmail)
	// EmailFormat validates if a string is an email or not. Note that it does NOT check if the MX record exists or not.
	EmailFormat = vax.String(govalidator.IsEmail, ErrEmail)
	// URL validates if a string is a valid URL
	URL = vax.String(govalidator.IsURL, ErrURL)
	// RequestURL validates if a string is a valid request URL
	RequestURL = vax.String(govalidator.IsRequestURL, ErrRequestURL)
	// RequestURI validates if a string is a valid request URI
	RequestURI = vax.String(govalidator.IsRequestURI, ErrRequestURI)
	// Alpha validates if a string contains English letters only (a-zA-Z)
	Alpha = vax.String(govalidator.IsAlpha, ErrAlpha)
	// Digit validates if a string contains digits only (0-9)
	Digit = vax.String(isDigit, ErrDigit)
	// Alphanumeric validates if a string contains English letters and digits only (a-zA-Z0-9)
	Alphanumeric = vax.String(govalidator.IsAlphanumeric, ErrAlphanumeric)
	// UTFLetter validates if a string contains unicode letters only
	UTFLetter = vax.String(govalidator.IsUTFLetter, ErrUTFLetter)
	// UTFDigit validates if a string contains unicode decimal digits only
	UTFDigit = vax.String(govalidator.IsUTFDigit, ErrUTFDigit)
	// UTFLetterNumeric validates if a string contains unicode letters and numbers only
	UTFLetterNumeric = vax.String(govalidator.IsUTFLetterNumeric, ErrUTFLetterNumeric)
	// UTFNumeric validates if a string contains unicode number characters (category N) only
	UTFNumeric = vax.String(isUTFNumeric, ErrUTFNumeric)
	// LowerCase validates if a string contains lower case unicode letters only
	LowerCase = vax.String(govalidator.IsLowerCase, ErrLowerCase)
	// UpperCase validates if a string contains upper case unicode letters only
	UpperCase = vax.String(govalidator.IsUpperCase, ErrUpperCase)
	// Hexadecimal validates if a string is a valid hexadecimal number
	Hexadecimal = vax.String(govalidator.IsHexadecimal, ErrHexadecimal)
	// HexColor validates if a string is a valid hexadecimal color code
	HexColor = vax.String(govalidator.IsHexcolor, ErrHexColor)
	// RGBColor validates if a string is a valid RGB color in the form of rgb(R, G, B)
	RGBColor = vax.String(govalidator.IsRGBcolor, ErrRGBColor)
	// Int validates if a string is a valid integer number
	Int = vax.String(govalidator.IsInt, ErrInt)
	// Float validates if a string is a floating point number
	Float = vax.String(govalidator.IsFloat, ErrFloat)
	// UUIDv3 validates if a string is a valid version 3 UUID
	UUIDv3 = vax.String(govalidator.IsUUIDv3, ErrUUIDv3)
	// UUIDv4 validates if a string is a valid version 4 UUID
	UUIDv4 = vax.String(govalidator.IsUUIDv4, ErrUUIDv4)
	// UUIDv5 validates if a string is a valid version 5 UUID
	UUIDv5 = vax.String(govalidator.IsUUIDv5, ErrUUIDv5)
	// UUID validates if a string is a valid UUID
	UUID = vax.String(govalidator.IsUUID, ErrUUID)
	// CreditCard validates if a string is a valid credit card number
	CreditCard = vax.String(govalidator.IsCreditCard, ErrCreditCard)
	// ISBN10 validates if a string is an ISBN version 10
	ISBN10 = vax.String(govalidator.IsISBN10, ErrISBN10)
	// ISBN13 validates if a string is an ISBN version 13
	ISBN13 = vax.String(govalidator.IsISBN13, ErrISBN13)
	// ISBN validates if a string is an ISBN (either version 10 or 13)
	ISBN = vax.String(isISBN, ErrISBN)
	// JSON validates if a string is in valid JSON format
	JSON = vax.String(govalidator.IsJSON, ErrJSON)
	// ASCII validates if a string contains ASCII characters only
	ASCII = vax.String(govalidator.IsASCII, ErrASCII)
	// PrintableASCII validates if a string contains printable ASCII characters only
	PrintableASCII = vax.String(govalidator.IsPrintableASCII, ErrPrintableASCII)
	// Multibyte validates if a string contains multibyte characters
	Multibyte = vax.String(govalidator.IsMultibyte, ErrMultibyte)
	// FullWidth validates if a string contains full-width characters
	FullWidth = vax.String(govalidator.IsFullWidth, ErrFullWidth)
	// HalfWidth validates if a string contains half-width characters
	HalfWidth = vax.String(govalidator.IsHalfWidth, ErrHalfWidth)
	// VariableWidth validates if a string contains both full-width and half-width characters
	VariableWidth = vax.String(govalidator.IsVariableWidth, ErrVariableWidth)
	// Base64 validates if a string is encoded in Base64
	Base64 = vax.String(govalidator.IsBase64, ErrBase64)
	// DataURI validates if a string is a valid base64-encoded data URI
	DataURI = vax.String(govalidator.IsDataURI, ErrDataURI)
	// E164 validates if a string is a valid ISO3166 Alpha 2 country code
	E164 = vax.String(isE164Number, ErrE164)
	// CountryCode2 validates if a string is a valid ISO3166 Alpha 2 country code
	CountryCode2 = vax.String(govalidator.IsISO3166Alpha2, ErrCountryCode2)
	// CountryCode3 validates if a string is a valid ISO3166 Alpha 3 country code
	CountryCode3 = vax.String(govalidator.IsISO3166Alpha3, ErrCountryCode3)
	// CurrencyCode validates if a string is a valid IsISO4217 currency code.
	CurrencyCode = vax.String(govalidator.IsISO4217, ErrCurrencyCode)
	// DialString validates if a string is a valid dial string that can be passed to Dial()
	DialString = vax.String(govalidator.IsDialString, ErrDialString)
	// MAC validates if a string is a MAC address
	MAC = vax.String(govalidator.IsMAC, ErrMac)
	// IP validates if a string is a valid IP address (either version 4 or 6)
	IP = vax.String(govalidator.IsIP, ErrIP)
	// IPv4 validates if a string is a valid version 4 IP address
	IPv4 = vax.String(govalidator.IsIPv4, ErrIPv4)
	// IPv6 validates if a string is a valid version 6 IP address
	IPv6 = vax.String(govalidator.IsIPv6, ErrIPv6)
	// Subdomain validates if a string is valid subdomain
	Subdomain = vax.String(isSubdomain, ErrSubdomain)
	// Domain validates if a string is valid domain
	Domain = vax.String(isDomain, ErrDomain)
	// DNSName validates if a string is valid DNS name
	DNSName = vax.String(govalidator.IsDNSName, ErrDNSName)
	// Host validates if a string is a valid IP (both v4 and v6) or a valid DNS name
	Host = vax.String(govalidator.IsHost, ErrHost)
	// Port validates if a string is a valid port number
	Port = vax.String(govalidator.IsPort, ErrPort)
	// MongoID validates if a string is a valid Mongo ID
	MongoID = vax.String(govalidator.IsMongoID, ErrMongoID)
	// Latitude validates if a string is a valid latitude
	Latitude = vax.String(govalidator.IsLatitude, ErrLatitude)
	// Longitude validates if a string is a valid longitude
	Longitude = vax.String(govalidator.IsLongitude, ErrLongitude)
	// SSN validates if a string is a social security number (SSN)
	SSN = vax.String(govalidator.IsSSN, ErrSSN)
	// Semver validates if a string is a valid semantic version
	Semver = vax.String(govalidator.IsSemver, ErrSemver)
)

var (
	reDigit = regexp.MustCompile("^[0-9]+$")
	// Subdomain regex source: https://stackoverflow.com/a/7933253
	reSubdomain = regexp.MustCompile(`^[A-Za-z0-9](?:[A-Za-z0-9\-]{0,61}[A-Za-z0-9])?$`)
	// E164 regex source: https://stackoverflow.com/a/23299989
	reE164 = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	// Domain regex source: https://stackoverflow.com/a/7933253
	// Slightly modified: Removed 255 max length vax since Go regex does not
	// support lookarounds. More info: https://stackoverflow.com/a/38935027
	reDomain = regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-z0-9])?\.)+(?:[a-zA-Z]{1,63}| xn--[a-z0-9]{1,59})$`)
)

func isISBN(value string) bool {
	return govalidator.IsISBN(value, 10) || govalidator.IsISBN(value, 13)
}

func isDigit(value string) bool {
	return reDigit.MatchString(value)
}

func isE164Number(value string) bool {
	return reE164.MatchString(value)
}

func isSubdomain(value string) bool {
	return reSubdomain.MatchString(value)
}

func isDomain(value string) bool {
	if len(value) > 255 {
		return false
	}

	return reDomain.MatchString(value)
}

func isUTFNumeric(value string) bool {
	for _, c := range value {
		if unicode.IsNumber(c) == false {
			return false
		}
	}
	return true
}
