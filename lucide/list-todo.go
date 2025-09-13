package lucide

import x "github.com/bloxui/blox"

// ListTodo creates a List Todo Lucide icon.
func ListTodo(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-list-todo", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 5h8"))),
		x.Child(x.Path(x.D("M13 12h8"))),
		x.Child(x.Path(x.D("M13 19h8"))),
		x.Child(x.Path(x.D("m3 17 2 2 4-4"))),
		x.Child(x.Rect(x.X("3"), x.Y("4"), x.RectWidth("6"), x.RectHeight("6"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
