package uuid

import (
	"fmt"
	"math/rand"
	"regexp"
)

func randomRange(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

// Generate will generate a new v4 UUID as a string
func Generate() string {
	return fmt.Sprintf(
		"%04x%04x-%04x-%04x-%04x-%04x%04x%04x",
		randomRange(0, 0xffff), randomRange(0, 0xffff),
		randomRange(0, 0xffff),
		randomRange(0, 0x0fff)|0x4000,
		randomRange(0, 0x3fff)|0x8000,
		randomRange(0, 0xffff), randomRange(0, 0xffff), randomRange(0, 0xffff),
	)
}

// Validate will validate the supplied string is a valid v4 UUID
func Validate(uuid string) bool {
	match, _ := regexp.MatchString(`^{?[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}?$`, uuid)
	return match
}
