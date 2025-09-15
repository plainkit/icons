package lucide

import x "github.com/plainkit/html"

// Warehouse creates a Warehouse Lucide icon.
func Warehouse(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-warehouse", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 21V10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v11"))),
		x.Child(x.Path(x.D("M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 1.132-1.803l7.95-3.974a2 2 0 0 1 1.837 0l7.948 3.974A2 2 0 0 1 22 8z"))),
		x.Child(x.Path(x.D("M6 13h12"))),
		x.Child(x.Path(x.D("M6 17h12"))),
	)
	return x.Svg(svgArgs...)
}
