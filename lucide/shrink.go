package lucide

import x "github.com/bloxui/blox"

// Shrink creates a Shrink Lucide icon.
func Shrink(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-shrink", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15 15 6 6m-6-6v4.8m0-4.8h4.8"))),
		x.Child(x.Path(x.D("M9 19.8V15m0 0H4.2M9 15l-6 6"))),
		x.Child(x.Path(x.D("M15 4.2V9m0 0h4.8M15 9l6-6"))),
		x.Child(x.Path(x.D("M9 4.2V9m0 0H4.2M9 9 3 3"))),
	)
	return x.Svg(svgArgs...)
}
