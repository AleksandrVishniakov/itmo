package lib

import (
	"fmt"
	"io"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Schedule []Schedule `hcl:"schedule,block" yaml:"schedule"`
}

type Schedule struct {
	Day    string   `hcl:",label" yaml:"day"`
	Lesson []Lesson `hcl:"lesson,block" yaml:"lesson"`
}

type Lesson struct {
	Name      string `hcl:",label" yaml:"name"`
	Type      string `hcl:"type" yaml:"type"`
	Teacher   string `hcl:"teacher" yaml:"teacher"`
	Room      *int   `hcl:"room" yaml:"room"`
	Address   string `hcl:"address" yaml:"address"`
	TimeStart string `hcl:"time_start" yaml:"time_start"`
	TimeEnd   string `hcl:"time_end" yaml:"time_end"`
	Online    bool   `hcl:"online" yaml:"online"`
}

func ParseSchedule(r io.Reader, w io.Writer) error {
	src, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	parser := hclparse.NewParser()

	file, diags := parser.ParseHCL(src, "schedule.hcl")
	if diags.HasErrors() {
		return fmt.Errorf("parse HCL: %s", diags.Error())
	}

	var config Config

	diags = gohcl.DecodeBody(file.Body, nil, &config)
	if diags.HasErrors() {
		return fmt.Errorf("decode HCL: %s", diags.Error())
	}

	encoder := yaml.NewEncoder(w)
	encoder.SetIndent(2)

	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("encode YAML: %s", err)
	}

	return encoder.Close()
}
