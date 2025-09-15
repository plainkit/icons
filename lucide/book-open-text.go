package lucide

import x "github.com/plainkit/blox"

// BookOpenText creates a Book Open Text Lucide icon.
func BookOpenText(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-book-open-text", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 7v14"))),
		x.Child(x.Path(x.D("M16 12h2"))),
		x.Child(x.Path(x.D("M16 8h2"))),
		x.Child(x.Path(x.D("M3 18a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h5a4 4 0 0 1 4 4 4 4 0 0 1 4-4h5a1 1 0 0 1 1 1v13a1 1 0 0 1-1 1h-6a3 3 0 0 0-3 3 3 3 0 0 0-3-3z"))),
		x.Child(x.Path(x.D("M6 12h2"))),
		x.Child(x.Path(x.D("M6 8h2"))),
	)
	return x.Svg(svgArgs...)
}
