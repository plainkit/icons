package lucide

import x "github.com/bloxui/blox"

// Eraser creates a Eraser Lucide icon.
func Eraser(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-eraser", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 21H8a2 2 0 0 1-1.42-.587l-3.994-3.999a2 2 0 0 1 0-2.828l10-10a2 2 0 0 1 2.829 0l5.999 6a2 2 0 0 1 0 2.828L12.834 21"))),
		x.Child(x.Path(x.D("m5.082 11.09 8.828 8.828"))),
	)
	return x.Svg(svgArgs...)
}