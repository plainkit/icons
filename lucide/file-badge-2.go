package lucide

import x "github.com/bloxui/blox"

// FileBadge2 creates a File Badge 2 Lucide icon.
func FileBadge2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-badge-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m13.69 12.479 1.29 4.88a.5.5 0 0 1-.697.591l-1.844-.849a1 1 0 0 0-.88.001l-1.846.85a.5.5 0 0 1-.693-.593l1.29-4.88"))),
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7z"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("10"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
