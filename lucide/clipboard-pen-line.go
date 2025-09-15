package lucide

import x "github.com/plainkit/html"

// ClipboardPenLine creates a Clipboard Pen Line Lucide icon.
func ClipboardPenLine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clipboard-pen-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("4"), x.X("8"), x.Y("2"), x.Rx("1"))),
		x.Child(x.Path(x.D("M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-.5"))),
		x.Child(x.Path(x.D("M16 4h2a2 2 0 0 1 1.73 1"))),
		x.Child(x.Path(x.D("M8 18h1"))),
		x.Child(x.Path(x.D("M21.378 12.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"))),
	)
	return x.Svg(svgArgs...)
}
