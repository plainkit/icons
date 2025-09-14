package lucide

import x "github.com/bloxui/blox"

// BadgeSwissFranc creates a Badge Swiss Franc Lucide icon.
func BadgeSwissFranc(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-badge-swiss-franc", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		x.Child(x.Path(x.D("M11 17V8h4"))),
		x.Child(x.Path(x.D("M11 12h3"))),
		x.Child(x.Path(x.D("M9 16h4"))),
	)
	return x.Svg(svgArgs...)
}
