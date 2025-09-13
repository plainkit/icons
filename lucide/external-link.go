package lucide

import x "github.com/bloxui/blox"

// ExternalLink creates a External Link Lucide icon.
func ExternalLink(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-external-link", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 3h6v6"))),
		x.Child(x.Path(x.D("M10 14 21 3"))),
		x.Child(x.Path(x.D("M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6"))),
	)
	return x.Svg(svgArgs...)
}
