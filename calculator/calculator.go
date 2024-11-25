package calculator

type calculator struct {
	num1 float64
	num2 float64
}

type memoryCalculator struct {
	memory float64
}

func (m *memoryCalculator) SetMemory(val float64) {
	m.memory = val
	return
}
func (m *memoryCalculator) AddToMemory(val float64) float64 {
	m.memory += val
	return m.memory
}
func (m *memoryCalculator) SubtractFromMemory(val float64) float64 {
	m.memory -= val
	return m.memory
}
func (m *memoryCalculator) Resetmemory() {
	m.memory = 0
	return
}
func AdvanceCalculator() memoryCalculator {
	m := &memoryCalculator{}
	return *m
}

func (c *calculator) add() float64 {
	return c.num1 + c.num2
}
func (c *calculator) subtract() float64 {
	return c.num1 - c.num2
}
func (c *calculator) multiply() float64 {
	return c.num1 * c.num2
}
func (c *calculator) divide() float64 {
	if c.num2 == 0 {
		return 0
	} else {
		return c.num1 / c.num2
	}
}
func Calulate(v1, v2 float64, op string) float64 {
	c := &calculator{v1, v2}
	switch op {
	case "+":
		return c.add()
	case "-":
		return c.subtract()
	case "*":
		return c.multiply()
	case "/":
		return c.divide()
	default:
		return 0
	}
}
