package decode

type Decoder interface {
	Decode(gifFile string) string
}
