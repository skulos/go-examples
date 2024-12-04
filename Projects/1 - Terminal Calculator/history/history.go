package history

type Operation uint8

const (
	ADD Operation = iota
	SUBTRACT
	MULTIPLY
	DIVIDE
	MODULUS
)

type HistoryType struct {
	Op     Operation
	First  float64
	Second float64
}

var history = []HistoryType{}

func Add(op Operation, first, second float64) {
	history = append(history, HistoryType{
		Op:     op,
		First:  first,
		Second: second,
	})
}

func List() []HistoryType {
	return history
}
