package evaluator

import "monkey/object"

var builtins = map[string]*object.Builtin{
	"let": {
		Fn: func(args ...object.Object) object.Object {
			return NULL
		},
	},
}
