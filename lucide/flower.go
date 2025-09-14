package lucide

import x "github.com/bloxui/blox"

// Flower creates a Flower Lucide icon.
func Flower(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flower", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("3"))),
		x.Child(x.Path(x.D("M12 16.5A4.5 4.5 0 1 1 7.5 12 4.5 4.5 0 1 1 12 7.5a4.5 4.5 0 1 1 4.5 4.5 4.5 4.5 0 1 1-4.5 4.5"))),
		x.Child(x.Path(x.D("M12 7.5V9"))),
		x.Child(x.Path(x.D("M7.5 12H9"))),
		x.Child(x.Path(x.D("M16.5 12H15"))),
		x.Child(x.Path(x.D("M12 16.5V15"))),
		x.Child(x.Path(x.D("m8 8 1.88 1.88"))),
		x.Child(x.Path(x.D("M14.12 9.88 16 8"))),
		x.Child(x.Path(x.D("m8 16 1.88-1.88"))),
		x.Child(x.Path(x.D("M14.12 14.12 16 16"))),
	)
	return x.Svg(svgArgs...)
}
