package lucide

import x "github.com/bloxui/blox"

// TableOfContents creates a Table Of Contents Lucide icon.
func TableOfContents(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-table-of-contents", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 5H3"))),
		x.Child(x.Path(x.D("M16 12H3"))),
		x.Child(x.Path(x.D("M16 19H3"))),
		x.Child(x.Path(x.D("M21 5h.01"))),
		x.Child(x.Path(x.D("M21 12h.01"))),
		x.Child(x.Path(x.D("M21 19h.01"))),
	)
	return x.Svg(svgArgs...)
}
