package lucide

import x "github.com/plainkit/blox"

// PenTool creates a Pen Tool Lucide icon.
func PenTool(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pen-tool", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15.707 21.293a1 1 0 0 1-1.414 0l-1.586-1.586a1 1 0 0 1 0-1.414l5.586-5.586a1 1 0 0 1 1.414 0l1.586 1.586a1 1 0 0 1 0 1.414z"))),
		x.Child(x.Path(x.D("m18 13-1.375-6.874a1 1 0 0 0-.746-.776L3.235 2.028a1 1 0 0 0-1.207 1.207L5.35 15.879a1 1 0 0 0 .776.746L13 18"))),
		x.Child(x.Path(x.D("m2.3 2.3 7.286 7.286"))),
		x.Child(x.Circle(x.Cx("11"), x.Cy("11"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
