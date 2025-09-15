package lucide

import x "github.com/plainkit/html"

// NotebookTabs creates a Notebook Tabs Lucide icon.
func NotebookTabs(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-notebook-tabs", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 6h4"))),
		x.Child(x.Path(x.D("M2 10h4"))),
		x.Child(x.Path(x.D("M2 14h4"))),
		x.Child(x.Path(x.D("M2 18h4"))),
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("20"), x.X("4"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Path(x.D("M15 2v20"))),
		x.Child(x.Path(x.D("M15 7h5"))),
		x.Child(x.Path(x.D("M15 12h5"))),
		x.Child(x.Path(x.D("M15 17h5"))),
	)
	return x.Svg(svgArgs...)
}
