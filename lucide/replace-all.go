package lucide

import x "github.com/bloxui/blox"

// ReplaceAll creates a Replace All Lucide icon.
func ReplaceAll(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-replace-all", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 14a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("M14 4a2 2 0 0 1 2-2"))),
		x.Child(x.Path(x.D("M16 10a2 2 0 0 1-2-2"))),
		x.Child(x.Path(x.D("M20 14a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("M20 2a2 2 0 0 1 2 2"))),
		x.Child(x.Path(x.D("M22 8a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("m3 7 3 3 3-3"))),
		x.Child(x.Path(x.D("M6 10V5a 3 3 0 0 1 3-3h1"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("8"), x.X("2"), x.Y("14"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
