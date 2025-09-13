package lucide

import x "github.com/bloxui/blox"

// Kanban creates a Kanban Lucide icon.
func Kanban(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-kanban", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 3v14"))),
		x.Child(x.Path(x.D("M12 3v8"))),
		x.Child(x.Path(x.D("M19 3v18"))),
	)
	return x.Svg(svgArgs...)
}
