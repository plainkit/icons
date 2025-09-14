package lucide

import x "github.com/bloxui/blox"

// Settings creates a Settings Lucide icon.
func Settings(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-settings", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9.671 4.136a2.34 2.34 0 0 1 4.659 0 2.34 2.34 0 0 0 3.319 1.915 2.34 2.34 0 0 1 2.33 4.033 2.34 2.34 0 0 0 0 3.831 2.34 2.34 0 0 1-2.33 4.033 2.34 2.34 0 0 0-3.319 1.915 2.34 2.34 0 0 1-4.659 0 2.34 2.34 0 0 0-3.32-1.915 2.34 2.34 0 0 1-2.33-4.033 2.34 2.34 0 0 0 0-3.831A2.34 2.34 0 0 1 6.35 6.051a2.34 2.34 0 0 0 3.319-1.915"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
