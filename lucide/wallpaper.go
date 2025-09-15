package lucide

import x "github.com/plainkit/blox"

// Wallpaper creates a Wallpaper Lucide icon.
func Wallpaper(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wallpaper", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 17v4"))),
		x.Child(x.Path(x.D("M8 21h8"))),
		x.Child(x.Path(x.D("m9 17 6.1-6.1a2 2 0 0 1 2.81.01L22 15"))),
		x.Child(x.Circle(x.Cx("8"), x.Cy("9"), x.R("2"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
