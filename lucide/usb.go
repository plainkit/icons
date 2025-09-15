package lucide

import x "github.com/plainkit/html"

// Usb creates a Usb Lucide icon.
func Usb(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-usb", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("10"), x.Cy("7"), x.R("1"))),
		x.Child(x.Circle(x.Cx("4"), x.Cy("20"), x.R("1"))),
		x.Child(x.Path(x.D("M4.7 19.3 19 5"))),
		x.Child(x.Path(x.D("m21 3-3 1 2 2Z"))),
		x.Child(x.Path(x.D("M9.26 7.68 5 12l2 5"))),
		x.Child(x.Path(x.D("m10 14 5 2 3.5-3.5"))),
		x.Child(x.Path(x.D("m18 12 1-1 1 1-1 1Z"))),
	)
	return x.Svg(svgArgs...)
}
