package lucide

import x "github.com/plainkit/html"

// FileDigit creates a File Digit Lucide icon.
func FileDigit(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-digit", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Rect(x.RectWidth("4"), x.RectHeight("6"), x.X("2"), x.Y("12"), x.Rx("2"))),
		x.Child(x.Path(x.D("M10 12h2v6"))),
		x.Child(x.Path(x.D("M10 18h4"))),
	)
	return x.Svg(svgArgs...)
}
