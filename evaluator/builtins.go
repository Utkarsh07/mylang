package evaluator

import (
	"mylang/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Function: func(arguments ...object.Object) object.Object {
			if len(arguments) != 1 {
				return newError("Wrong number of arguments. Got=%d, Expected=1", len(arguments))
			}
			switch argument := arguments[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(argument.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(argument.Elements))}
			default:
				return newError("argument to `len` not supported, Got=%s",
					arguments[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Function: func(arguments ...object.Object) object.Object {
			if len(arguments) != 1 {
				return newError("Wrong number of arguments. Got=%d, Expected=1", len(arguments))
			}
			if arguments[0].Type() != object.ARRAY_OBJ {
				return newError("arguments to `first` must be ARRAY, GOT=%s", arguments[0].Type())
			}
			array := arguments[0].(*object.Array)
			if len(array.Elements) > 0 {
				return array.Elements[0]
			}
			return NULL
		},
	},
	"last": &object.Builtin{
		Function: func(arguments ...object.Object) object.Object {
			if len(arguments) != 1 {
				return newError("Wrong number of arguments. Got=%d, Expected=1", len(arguments))
			}
			if arguments[0].Type() != object.ARRAY_OBJ {
				return newError("arguments to `last` must be ARRAY, Got=%s", arguments[0].Type())
			}
			array := arguments[0].(*object.Array)
			if len(array.Elements) > 0 {
				return array.Elements[len(array.Elements)-1]
			}
			return NULL
		},
	},
	"rest": &object.Builtin{
		Function: func(arguments ...object.Object) object.Object {
			if len(arguments) != 1 {
				return newError("Wrong number of arguments. Got=%d, Expected=1", len(arguments))
			}
			if arguments[0].Type() != object.ARRAY_OBJ {
				return newError("arguments to `rest` must be ARRAY, Got=%s", arguments[0].Type())
			}
			array := arguments[0].(*object.Array)
			length := len(array.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, array.Elements[1:length])
				return &object.Array{Elements: newElements}
			}
			return NULL
		},
	},
	"push": &object.Builtin{
		Function: func(arguments ...object.Object) object.Object {
			if len(arguments) != 2 {
				return newError("Wrong number of arguments. Got=%d, Expected=2", len(arguments))
			}
			if arguments[0].Type() != object.ARRAY_OBJ {
				return newError("arguments to `push` must be ARRAY, Got=%s", arguments[0].Type())
			}
			array := arguments[0].(*object.Array)
			length := len(array.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, array.Elements)
			newElements[length] = arguments[1]

			return &object.Array{Elements: newElements}
		},
	},
}
