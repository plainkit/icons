package lucide

import x "github.com/bloxui/blox"

// Ligature creates a Ligature Lucide icon.
func Ligature(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-ligature", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 12h2v8"))),
		x.Child(x.Path(x.D("M14 20h4"))),
		x.Child(x.Path(x.D("M6 12h4"))),
		x.Child(x.Path(x.D("M6 20h4"))),
		x.Child(x.Path(x.D("M8 20V8a4 4 0 0 1 7.464-2"))),
	)
	return x.Svg(svgArgs...)
}
