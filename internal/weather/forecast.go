package weather

import (
	"context"
)

func GetProbabilityOfPrecipitation(appId string, city string) (float64, error) {
	c := DefaultClient(appId)
	fr, err := c.GetForecast(context.Background(), city)
	if err != nil {
		return 0, err
	}
	pop := fr.List[0].Pop
	return pop, nil
}
