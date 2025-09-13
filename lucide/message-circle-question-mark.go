package lucide

import x "github.com/bloxui/blox"

// MessageCircleQuestionMark creates a Message Circle Question Mark Lucide icon.
func MessageCircleQuestionMark(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-message-circle-question-mark", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719"))),
		x.Child(x.Path(x.D("M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"))),
		x.Child(x.Path(x.D("M12 17h.01"))),
	)
	return x.Svg(svgArgs...)
}
