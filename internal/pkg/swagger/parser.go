package swagger

import (
	"fmt"
	"log"

	"github.com/swaggest/openapi-go/openapi3"
)

func Parse(spec []byte) openapi3.Spec {
	var s openapi3.Spec
	if error := s.UnmarshalYAML(spec); error != nil {
		log.Fatal("the schema yaml is invalid")
	}

	fmt.Println(s)

	return s
}
