package lucide

import x "github.com/bloxui/blox"

// Coffee creates a Coffee Lucide icon.
func Coffee(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-coffee", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 2v2"))),
		x.Child(x.Path(x.D("M14 2v2"))),
		x.Child(x.Path(x.D("M16 8a1 1 0 0 1 1 1v8a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4V9a1 1 0 0 1 1-1h14a4 4 0 1 1 0 8h-1"))),
		x.Child(x.Path(x.D("M6 2v2"))),
	)
	return x.Svg(svgArgs...)
}