package dummy_1

import (
	"context"
	"dummy_1/pb"
)

type Dummy1Server struct{}

func NewDummy1Server() pb.Dummy1Server {
	return &Dummy1Server{}
}

func (s *Dummy1Server) GetDummy1Method1(ctx context.Context, in *pb.GetDummy1Method1Request) (*pb.Dummy1Method1Object, error) {
	return &pb.Dummy1Method1Object{
		Id:   "dummy_1_method_1_id",
		Name: "Dummy 1 Method 1",
	}, nil
}

func (s *Dummy1Server) GetDummy1Method2(ctx context.Context, in *pb.GetDummy1Method2Request) (*pb.Dummy1Method2Object, error) {
	return &pb.Dummy1Method2Object{
		Id:   "dummy_1_method_2_id",
		Name: "Dummy 1 Method 2",
	}, nil
}
