package lucide

import x "github.com/bloxui/blox"

// MousePointerBan creates a Mouse Pointer Ban Lucide icon.
func MousePointerBan(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mouse-pointer-ban", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2.034 2.681a.498.498 0 0 1 .647-.647l9 3.5a.5.5 0 0 1-.033.944L8.204 7.545a1 1 0 0 0-.66.66l-1.066 3.443a.5.5 0 0 1-.944.033z"))),
		x.Child(x.Circle(x.Cx("16"), x.Cy("16"), x.R("6"))),
		x.Child(x.Path(x.D("m11.8 11.8 8.4 8.4"))),
	)
	return x.Svg(svgArgs...)
}
