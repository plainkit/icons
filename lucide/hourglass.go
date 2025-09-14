package lucide

import x "github.com/bloxui/blox"

// Hourglass creates a Hourglass Lucide icon.
func Hourglass(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-hourglass", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 22h14"))),
		x.Child(x.Path(x.D("M5 2h14"))),
		x.Child(x.Path(x.D("M17 22v-4.172a2 2 0 0 0-.586-1.414L12 12l-4.414 4.414A2 2 0 0 0 7 17.828V22"))),
		x.Child(x.Path(x.D("M7 2v4.172a2 2 0 0 0 .586 1.414L12 12l4.414-4.414A2 2 0 0 0 17 6.172V2"))),
	)
	return x.Svg(svgArgs...)
}
