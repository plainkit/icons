package lucide

import x "github.com/plainkit/html"

// Tag creates a Tag Lucide icon.
func Tag(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tag", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.586 2.586A2 2 0 0 0 11.172 2H4a2 2 0 0 0-2 2v7.172a2 2 0 0 0 .586 1.414l8.704 8.704a2.426 2.426 0 0 0 3.42 0l6.58-6.58a2.426 2.426 0 0 0 0-3.42z"))),
		x.Child(x.Circle(x.Cx("7.5"), x.Cy("7.5"), x.R(".5"), x.Fill("currentColor"))),
	)
	return x.Svg(svgArgs...)
}
