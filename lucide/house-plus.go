package lucide

import x "github.com/plainkit/html"

// HousePlus creates a House Plus Lucide icon.
func HousePlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-house-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.35 21H5a2 2 0 0 1-2-2v-9a2 2 0 0 1 .71-1.53l7-6a2 2 0 0 1 2.58 0l7 6A2 2 0 0 1 21 10v2.35"))),
		x.Child(x.Path(x.D("M14.8 12.4A1 1 0 0 0 14 12h-4a1 1 0 0 0-1 1v8"))),
		x.Child(x.Path(x.D("M15 18h6"))),
		x.Child(x.Path(x.D("M18 15v6"))),
	)
	return x.Svg(svgArgs...)
}
