package price

import (
	"context"
	"github.com/gookit/goutil/testutil/assert"
	"microservice-test/proto/price"
	"testing"
)

func TestGetPrice(t *testing.T) {
	ps := &Service{}

	// Test case 1: Special Day
	request1 := &price.PriceRequest{Date: "2023-11-20"}
	response1, err1 := ps.GetPrice(context.Background(), request1)

	assert.Nil(t, err1)
	assert.Equal(t, response1.Price, int64(200000))

	// Test case 2: Normal Day
	request2 := &price.PriceRequest{Date: "2023-11-11"}
	response2, err2 := ps.GetPrice(context.Background(), request2)

	assert.Nil(t, err2)
	assert.Equal(t, response2.Price, int64(100000))

	// Test case 3: Invalid Date
	request3 := &price.PriceRequest{Date: "2023-11-11-11"}
	response3, err3 := ps.GetPrice(context.Background(), request3)

	assert.NotNil(t, err3)
	assert.Nil(t, response3, nil)

	// Test case 4: Empty Date
	request4 := &price.PriceRequest{Date: ""}
	response4, err4 := ps.GetPrice(context.Background(), request4)

	assert.NotNil(t, err4)
	assert.Nil(t, response4, nil)
}
