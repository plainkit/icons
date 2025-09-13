package lucide

import x "github.com/bloxui/blox"

// Archive creates an Archive Lucide icon.
func Archive(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-archive", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("5"), x.X("2"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Path(x.D("M4 8v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8"))),
		x.Child(x.Path(x.D("M10 12h4"))),
	)
	return x.Svg(svgArgs...)
}
