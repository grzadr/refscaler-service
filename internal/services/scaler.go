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
	log.Printf("received items to scale\n")
	reader := strings.NewReader(enlistment_query)
	enlistment, err := refscaler.NewEnlistment(
		reader,
		units.EmbeddedUnitRegistry,
	)
	if err != nil {
		log.Printf("failed to create enlistments: %s", err)
		return nil, ErrEnlistmentCreate
	}
	log.Printf("loaded %d items", enlistment.Length())

	scale, err := enlistment.MakeMeasureValue(scale_query)
	if err != nil {
		log.Printf("failed to convert '%s': %s", scale_query, err)
		return nil, ErrScaleConvert
	}
	log.Printf("scaling to %s", scale_query)

	return enlistment.GetScaled(scale).ToString(3), nil
}
