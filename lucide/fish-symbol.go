package lucide

import x "github.com/bloxui/blox"

// FishSymbol creates a Fish Symbol Lucide icon.
func FishSymbol(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-fish-symbol", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 16s9-15 20-4C11 23 2 8 2 8"))),
	)
	return x.Svg(svgArgs...)
}
