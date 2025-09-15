package lucide

import x "github.com/plainkit/html"

// Ban creates a Ban Lucide icon.
func Ban(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ban", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4.929 4.929 19.07 19.071"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
