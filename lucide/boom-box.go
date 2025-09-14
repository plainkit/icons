package lucide

import x "github.com/bloxui/blox"

// BoomBox creates a Boom Box Lucide icon.
func BoomBox(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-boom-box", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 9V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v4"))),
		x.Child(x.Path(x.D("M8 8v1"))),
		x.Child(x.Path(x.D("M12 8v1"))),
		x.Child(x.Path(x.D("M16 8v1"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("12"), x.X("2"), x.Y("9"), x.Rx("2"))),
		x.Child(x.Circle(x.Cx("8"), x.Cy("15"), x.R("2"))),
		x.Child(x.Circle(x.Cx("16"), x.Cy("15"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
