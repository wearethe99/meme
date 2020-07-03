package source

type Codec string

const (
	JPEG = Codec(iota)
	GIF
	PNG
	UNKNOWN
)
