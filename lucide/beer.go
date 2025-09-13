package lucide

import x "github.com/bloxui/blox"

// Beer creates a Beer Lucide icon.
func Beer(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-beer", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 11h1a3 3 0 0 1 0 6h-1"))),
		x.Child(x.Path(x.D("M9 12v6"))),
		x.Child(x.Path(x.D("M13 12v6"))),
		x.Child(x.Path(x.D("M14 7.5c-1 0-1.44.5-3 .5s-2-.5-3-.5-1.72.5-2.5.5a2.5 2.5 0 0 1 0-5c.78 0 1.57.5 2.5.5S9.44 2 11 2s2 1.5 3 1.5 1.72-.5 2.5-.5a2.5 2.5 0 0 1 0 5c-.78 0-1.5-.5-2.5-.5Z"))),
		x.Child(x.Path(x.D("M5 8v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V8"))),
	)
	return x.Svg(svgArgs...)
}
