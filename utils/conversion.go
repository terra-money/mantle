package utils

import (
	"fmt"
	"github.com/iancoleman/strcase"
)

func MsgRouteAndTypeToString(route, t string) string {
	return fmt.Sprintf("%s/Msg%s", route, strcase.ToCamel(t))
}

