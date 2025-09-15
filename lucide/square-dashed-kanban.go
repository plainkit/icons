package lucide

import x "github.com/plainkit/html"

// SquareDashedKanban creates a Square Dashed Kanban Lucide icon.
func SquareDashedKanban(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-dashed-kanban", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 7v7"))),
		x.Child(x.Path(x.D("M12 7v4"))),
		x.Child(x.Path(x.D("M16 7v9"))),
		x.Child(x.Path(x.D("M5 3a2 2 0 0 0-2 2"))),
		x.Child(x.Path(x.D("M9 3h1"))),
		x.Child(x.Path(x.D("M14 3h1"))),
		x.Child(x.Path(x.D("M19 3a2 2 0 0 1 2 2"))),
		x.Child(x.Path(x.D("M21 9v1"))),
		x.Child(x.Path(x.D("M21 14v1"))),
		x.Child(x.Path(x.D("M21 19a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("M14 21h1"))),
		x.Child(x.Path(x.D("M9 21h1"))),
		x.Child(x.Path(x.D("M5 21a2 2 0 0 1-2-2"))),
		x.Child(x.Path(x.D("M3 14v1"))),
		x.Child(x.Path(x.D("M3 9v1"))),
	)
	return x.Svg(svgArgs...)
}
