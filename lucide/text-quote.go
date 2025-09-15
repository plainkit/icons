package lucide

import x "github.com/plainkit/html"

// TextQuote creates a Text Quote Lucide icon.
func TextQuote(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-text-quote", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 5H3"))),
		x.Child(x.Path(x.D("M21 12H8"))),
		x.Child(x.Path(x.D("M21 19H8"))),
		x.Child(x.Path(x.D("M3 12v7"))),
	)
	return x.Svg(svgArgs...)
}
