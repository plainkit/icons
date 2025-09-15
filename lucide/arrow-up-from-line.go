package lucide

import x "github.com/plainkit/html"

// ArrowUpFromLine creates a Arrow Up From Line Lucide icon.
func ArrowUpFromLine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-up-from-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m18 9-6-6-6 6"))),
		x.Child(x.Path(x.D("M12 3v14"))),
		x.Child(x.Path(x.D("M5 21h14"))),
	)
	return x.Svg(svgArgs...)
}
