package services

import (
	"errors"
	"log"
	"strings"

	"github.com/grzadr/refscaler/refscaler"
	"github.com/grzadr/refscaler/units"
)

var (
	ErrEnlistmentCreate = errors.New("failed to create enlistments")
	ErrScaleConvert     = errors.New("failed to convert scale")
)

func GetScaled(enlistment_query string, scale_query string) ([]string, error) {
	reader := strings.NewReader(enlistment_query)
	enlistment, err := refscaler.NewEnlistment(
		reader,
		units.EmbeddedUnitRegistry,
	)
	if err != nil {
		log.Fatalf("failed to create enlistments: %s", err)
		return nil, ErrEnlistmentCreate
	}

	scale, err := enlistment.MakeMeasureValue(scale_query)
	if err != nil {
		log.Fatalf("failed to convert '%s': %s", scale_query, err)
		return nil, ErrScaleConvert
	}

	return enlistment.GetScaled(scale).ToString(3), nil
}
