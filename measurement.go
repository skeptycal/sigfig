package sigfig

type (
	measurement struct {
		value   Floater
		sigfigs int
		unit    string
	}
)
