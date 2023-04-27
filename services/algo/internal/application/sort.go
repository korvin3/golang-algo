package application

import (
	"golang.org/x/exp/constraints"
	"github.com/korvin3/golang-algo/services/algo/internal/domain/sort"
	"github.com/korvin3/golang-algo/services/algo/internal/domain/sort/constants"
	"github.com/korvin3/golang-algo/services/algo/internal/domain/calc_result/model"
    "github.com/google/uuid"

	"time"
	"encoding/json"
)

type Command[K constraints.Ordered] struct {
	input []K
	order constants.Order
}

func SortHandler(c Command) (model.Calc_Result, error) {
	startTime := time.Now()
	result := sort.MergeSort(c.input, c.order)
	endTime := time.Now()

	jsonResult, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return model.Calc_Result{
			uuid: uuid.New(), 
			json: jsonResult, 
			started_at: startTime, 
			completed_at: endTime,
		}, 
		nil
}