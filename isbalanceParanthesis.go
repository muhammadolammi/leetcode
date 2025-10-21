package main

func isBalancedParentensis(input string) bool {
	stack := MakeStack[byte]()
	inputBytes := []byte(input)
	for _, byte := range inputBytes {
		if string(byte) == "(" {
			stack.Push(byte)
		}
		if string(byte) == ")" {
			_, popped := stack.Pop()
			if !popped {
				return false
			}
		}
	}
	return stack.Size() == 0

}
