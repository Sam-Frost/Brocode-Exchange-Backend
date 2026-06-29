package services

import (
	"github.com/Sam-Frost/matching-engine/internal/util"
	"github.com/twmb/franz-go/pkg/kgo"
)

func StartFillProducer(ringBuffer *util.RingBuffer[[]Fill], kakfaClient *kgo.Client) {

}
