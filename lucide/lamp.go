package lucide

import x "github.com/plainkit/blox"

// Lamp creates a Lamp Lucide icon.
func Lamp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-lamp", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 12v6"))),
		x.Child(x.Path(x.D("M4.077 10.615A1 1 0 0 0 5 12h14a1 1 0 0 0 .923-1.385l-3.077-7.384A2 2 0 0 0 15 2H9a2 2 0 0 0-1.846 1.23Z"))),
		x.Child(x.Path(x.D("M8 20a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1z"))),
	)
	return x.Svg(svgArgs...)
}
