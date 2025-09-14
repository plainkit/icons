package lucide

import x "github.com/bloxui/blox"

// RussianRuble creates a Russian Ruble Lucide icon.
func RussianRuble(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-russian-ruble", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 11h8a4 4 0 0 0 0-8H9v18"))),
		x.Child(x.Path(x.D("M6 15h8"))),
	)
	return x.Svg(svgArgs...)
}
