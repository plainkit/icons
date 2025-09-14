package lucide

import x "github.com/bloxui/blox"

// FileQuestionMark creates a File Question Mark Lucide icon.
func FileQuestionMark(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-question-mark", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 17h.01"))),
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7z"))),
		x.Child(x.Path(x.D("M9.1 9a3 3 0 0 1 5.82 1c0 2-3 3-3 3"))),
	)
	return x.Svg(svgArgs...)
}
