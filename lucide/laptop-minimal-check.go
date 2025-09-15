package lucide

import x "github.com/plainkit/blox"

// LaptopMinimalCheck creates a Laptop Minimal Check Lucide icon.
func LaptopMinimalCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-laptop-minimal-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 20h20"))),
		x.Child(x.Path(x.D("m9 10 2 2 4-4"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("12"), x.X("3"), x.Y("4"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
