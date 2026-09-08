package operations

import (
	"encoding/json"
	"strconv"
	"strings"

	"storage-api/shared/events"
	"storage-api/shared/logs"
)

type Pipeline struct {
	Name      string
	steps     []Step
	results   []StepResult
	publisher *events.Publisher
}

type Step struct {
	Name   string
	Action func() error
}

type StepStatus string

const (
	StepStatusOK     StepStatus = "OK"
	StepStatusFailed StepStatus = "FAILED"
)

type StepResult struct {
	Step   uint       `json:"step"`
	Name   string     `json:"name"`
	Status StepStatus `json:"status"`
}

func NewPipeline(name string) *Pipeline {
	publisher := events.NewPublisher()

	publisher.Subscribe(
		logs.EventName,
		logs.Log,
	)

	return &Pipeline{
		Name:      name,
		publisher: publisher,
	}
}

func (p *Pipeline) Step(
	name string,
	action func() error,
) *Pipeline {

	p.steps = append(
		p.steps,
		Step{
			Name:   name,
			Action: action,
		},
	)

	return p
}

func (p *Pipeline) Execute() error {
	p.results = nil

	p.logPipelineStart()

	for i, step := range p.steps {
		stepNumber := uint(i + 1)

		if err := p.run(step, stepNumber); err != nil {
			p.logSummary(StepStatusFailed)

			return err
		}
	}

	p.logSummary(StepStatusOK)

	return nil
}

func (p *Pipeline) ExecuteAlways() []error {
	p.results = nil

	var errors []error

	p.logPipelineStart()

	for i, step := range p.steps {
		stepNumber := uint(i + 1)

		if err := p.run(step, stepNumber); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		p.logSummary(StepStatusFailed)
	} else {
		p.logSummary(StepStatusOK)
	}

	return errors
}

func (p *Pipeline) run(
	step Step,
	stepNumber uint,
) error {

	var message strings.Builder

	message.WriteString("[STEP ")
	message.WriteString(
		strconv.FormatUint(
			uint64(stepNumber),
			10,
		),
	)
	message.WriteString("]: [")
	message.WriteString(step.Name)
	message.WriteString("] Starting")

	p.log(
		message.String(),
		logs.LogTypeInfo,
	)

	err := step.Action()

	if err != nil {

		p.results = append(
			p.results,
			StepResult{
				Step:   stepNumber,
				Name:   step.Name,
				Status: StepStatusFailed,
			},
		)

		message.Reset()

		message.WriteString("[STEP ")
		message.WriteString(
			strconv.FormatUint(
				uint64(stepNumber),
				10,
			),
		)
		message.WriteString("]: [")
		message.WriteString(step.Name)
		message.WriteString("] Failed: ")
		message.WriteString(err.Error())

		p.log(
			message.String(),
			logs.LogTypeError,
		)

		return err
	}

	p.results = append(
		p.results,
		StepResult{
			Step:   stepNumber,
			Name:   step.Name,
			Status: StepStatusOK,
		},
	)

	message.Reset()

	message.WriteString("[STEP ")
	message.WriteString(
		strconv.FormatUint(
			uint64(stepNumber),
			10,
		),
	)
	message.WriteString("]: [")
	message.WriteString(step.Name)
	message.WriteString("] Success")

	p.log(
		message.String(),
		logs.LogTypeInfo,
	)

	return nil
}

func (p *Pipeline) logPipelineStart() {
	var message strings.Builder

	message.WriteString("[PIPELINE]: [")
	message.WriteString(p.Name)
	message.WriteString("] Starting")

	p.log(
		message.String(),
		logs.LogTypeInfo,
	)
}

func (p *Pipeline) logSummary(status StepStatus) {
	resultJSON, err := json.Marshal(p.results)

	if err != nil {
		p.log(
			"[PIPELINE]: Failed to generate execution summary",
			logs.LogTypeError,
		)

		return
	}

	var message strings.Builder

	message.WriteString("[PIPELINE]: [")
	message.WriteString(p.Name)
	message.WriteString("] ")
	message.WriteString(string(status))
	message.WriteString(" | ")
	message.Write(resultJSON)

	logType := logs.LogTypeInfo

	if status == StepStatusFailed {
		logType = logs.LogTypeError
	}

	p.log(
		message.String(),
		logType,
	)
}

func (p *Pipeline) log(
	message string,
	logType logs.LogType,
) {
	p.publisher.Publish(
		logs.EventName,
		logs.NewLogEvent(
			message,
			logType,
		),
	)
}