package converter

import "fmt"

type NumeralSystem interface {
	Name() string
	Base() any
	Parse(string) (float64, error)
	Format(float64) (string, error)
}

type Converter struct {
	registry map[string]NumeralSystem
}

func New() *Converter {
	return &Converter{
		registry: make(map[string]NumeralSystem),
	}
}

func (c *Converter) Register(id string, ns NumeralSystem) {
	c.registry[id] = ns
}

func (c *Converter) Convert(value string, from, to string) (string, error) {
	fromNS, ok := c.registry[from]
	if !ok {
		return "", fmt.Errorf("numeral system %q not registered", from)
	}

	toNS, ok := c.registry[to]
	if !ok {
		return "", fmt.Errorf("numeral system %q not registered", to)
	}

	n, err := fromNS.Parse(value)
	if err != nil {
		return "", fmt.Errorf("%s: %w", fromNS.Name(), err)
	}

	result, err := toNS.Format(n)
	if err != nil {
		return "", fmt.Errorf("%s: %w", toNS.Name(), err)
	}

	return result, nil
}

func (c *Converter) NumericalSystem(id string) NumeralSystem {
	return c.registry[id]
}
