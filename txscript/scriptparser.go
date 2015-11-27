package txscript

type Op struct {
	Opcode string
	Data   []byte
}

func ParseScript(script []byte) ([]Op, error) {
	opcodes, err := parseScript(script)
	if nil != err {
		return nil, err
	}

	parsedOps := []Op{}
	for _, pop := range opcodes {
		opcodeName := pop.opcode.name
		op := Op{
			Opcode: opcodeName,
			Data:   pop.data,
		}
		parsedOps = append(parsedOps, op)
	}

	return parsedOps, nil
}
