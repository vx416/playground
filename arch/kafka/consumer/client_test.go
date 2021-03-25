package consumer

import (
	"testing"
	"time"

	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/suite"
)

func TestClient(t *testing.T) {
	suite.Run(t, new(clientSuite))
}

type clientSuite struct {
	client sarama.Client
	suite.Suite
}

func (s *clientSuite) SetupSuite() {
	//
	// client, err := NewClient([]string{"127.0.0.1:9092"}, nil)
	client, err := NewClient([]string{"10.20.0.164:31409"}, nil)
	s.Require().NoError(err)
	s.client = client
}

func (s *clientSuite) TestGet() {
	co, err := s.client.Coordinator("ems")
	s.Require().NoError(err)
	req := &sarama.OffsetRequest{}
	req.AddBlock("matching.match_result", 0, 0, 10)
	resp, err := co.GetAvailableOffsets(req)
	// resp, err := co.GetMetadata(req)
	s.Require().NoError(err)
	s.T().Log(resp.Blocks["matching.match_result"][0].Offset)
}

func (s *clientSuite) TestPartion() {
	paritions, err := s.client.Partitions("test_topic")
	s.Require().NoError(err)
	s.T().Log(paritions)

	broker, err := s.client.Leader("test_topic", paritions[0])
	s.Require().NoError(err)
	s.T().Log(broker.Addr())

	prodReq := &sarama.ProduceRequest{
		RequiredAcks: sarama.WaitForLocal,
	}
	prodReq.AddMessage("test_topic", paritions[0], &sarama.Message{Value: []byte("Testing")})
	prodresp, err := broker.Produce(prodReq)
	s.Require().NoError(err)
	s.T().Log(prodresp.Blocks["test_topic"][0].Offset)

	req := &sarama.FetchRequest{
		MaxBytes:    1000,
		MaxWaitTime: int32(2 * time.Second),
	}
	req.AddBlock("test_topic", paritions[0], 21, 10000)
	resp, err := broker.Fetch(req)
	s.Require().NoError(err)
	// s.T().Log(resp.Blocks["test_topic"][paritions[2]].RecordsSet[0].MsgSet.Messages[0].Msg.Key)
	s.T().Log(string(resp.Blocks["test_topic"][paritions[0]].RecordsSet[0].MsgSet.Messages[0].Msg.Value))
}
