package lucide

import x "github.com/bloxui/blox"

// BrickWall creates a Brick Wall Lucide icon.
func BrickWall(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-brick-wall", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M12 9v6"))),
		x.Child(x.Path(x.D("M16 15v6"))),
		x.Child(x.Path(x.D("M16 3v6"))),
		x.Child(x.Path(x.D("M3 15h18"))),
		x.Child(x.Path(x.D("M3 9h18"))),
		x.Child(x.Path(x.D("M8 15v6"))),
		x.Child(x.Path(x.D("M8 3v6"))),
	)
	return x.Svg(svgArgs...)
}
