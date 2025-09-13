package lucide

import x "github.com/bloxui/blox"

// NotebookText creates a Notebook Text Lucide icon.
func NotebookText(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-notebook-text", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 6h4"))),
		x.Child(x.Path(x.D("M2 10h4"))),
		x.Child(x.Path(x.D("M2 14h4"))),
		x.Child(x.Path(x.D("M2 18h4"))),
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("20"), x.X("4"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Path(x.D("M9.5 8h5"))),
		x.Child(x.Path(x.D("M9.5 12H16"))),
		x.Child(x.Path(x.D("M9.5 16H14"))),
	)
	return x.Svg(svgArgs...)
}
