package processor

type Processor interface {
	Process(input string) (string, error)
}
