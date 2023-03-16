package futures

import (
	"context"
	"fmt"
	"net/http"
)

// BuySellVolumeService list buy and sell volume
type BuySellVolumeService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *BuySellVolumeService) Symbol(symbol string) *BuySellVolumeService {
	s.symbol = symbol
	return s
}

// Period set period
func (s *BuySellVolumeService) Period(period string) *BuySellVolumeService {
	s.period = period
	return s
}

// Limit set limit
func (s *BuySellVolumeService) Limit(limit int) *BuySellVolumeService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *BuySellVolumeService) StartTime(startTime int64) *BuySellVolumeService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *BuySellVolumeService) EndTime(endTime int64) *BuySellVolumeService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *BuySellVolumeService) Do(ctx context.Context, opts ...RequestOption) (res []*BuySellVolume, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/takerlongshortRatio",
	}
	r.setParam("symbol", s.symbol)
	r.setParam("period", s.period)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*BuySellVolume{}, err
	}
	j, err := newJSON(data)
	if err != nil {
		return []*BuySellVolume{}, err
	}

	num := len(j.MustArray())
	res = make([]*BuySellVolume, num)
	for i := 0; i < num; i++ {
		item := j.GetIndex(i)
		if len(item.MustMap()) < 4 {
			err = fmt.Errorf("invalid BuySellVolume response")
			return []*BuySellVolume{}, err
		}
		res[i] = &BuySellVolume{
			buySellRatio: item.Get("buySellRatio").MustString(),
			buyVol:       item.Get("buyVol").MustString(),
			sellVol:      item.Get("sellVol").MustString(),
			timestamp:    item.Get("timestamp").MustInt64(),
		}
	}

	return res, nil
}

// BuySellVolume defines BuySellVolume info
type BuySellVolume struct {
	buySellRatio string `json:"buySellRatio"`
	buyVol       string `json:"buyVol"`
	sellVol      string `json:"sellVol"`
	timestamp    int64  `json:"timestamp"`
}
