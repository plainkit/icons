package lucide

import x "github.com/bloxui/blox"

// LaptopMinimal creates a Laptop Minimal Lucide icon.
func LaptopMinimal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-laptop-minimal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("12"), x.X("3"), x.Y("4"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Line(x.X1("2"), x.X2("22"), x.Y1("20"), x.Y2("20"))),
	)
	return x.Svg(svgArgs...)
}
