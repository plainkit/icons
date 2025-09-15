package lucide

import x "github.com/plainkit/blox"

// Image creates a Image Lucide icon.
func Image(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-image", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("9"), x.R("2"))),
		x.Child(x.Path(x.D("m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21"))),
	)
	return x.Svg(svgArgs...)
}
