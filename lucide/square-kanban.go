package lucide

import x "github.com/plainkit/blox"

// SquareKanban creates a Square Kanban Lucide icon.
func SquareKanban(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-kanban", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M8 7v7"))),
		x.Child(x.Path(x.D("M12 7v4"))),
		x.Child(x.Path(x.D("M16 7v9"))),
	)
	return x.Svg(svgArgs...)
}
