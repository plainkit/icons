package lucide

import x "github.com/bloxui/blox"

// Logs creates a Logs Lucide icon.
func Logs(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-logs", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 5h1"))),
		x.Child(x.Path(x.D("M3 12h1"))),
		x.Child(x.Path(x.D("M3 19h1"))),
		x.Child(x.Path(x.D("M8 5h1"))),
		x.Child(x.Path(x.D("M8 12h1"))),
		x.Child(x.Path(x.D("M8 19h1"))),
		x.Child(x.Path(x.D("M13 5h8"))),
		x.Child(x.Path(x.D("M13 12h8"))),
		x.Child(x.Path(x.D("M13 19h8"))),
	)
	return x.Svg(svgArgs...)
}
