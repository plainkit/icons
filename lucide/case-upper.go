package lucide

import x "github.com/plainkit/html"

// CaseUpper creates a Case Upper Lucide icon.
func CaseUpper(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-case-upper", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 11h4.5a1 1 0 0 1 0 5h-4a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .5-.5h3a1 1 0 0 1 0 5"))),
		x.Child(x.Path(x.D("m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16"))),
		x.Child(x.Path(x.D("M3.304 13h6.392"))),
	)
	return x.Svg(svgArgs...)
}
