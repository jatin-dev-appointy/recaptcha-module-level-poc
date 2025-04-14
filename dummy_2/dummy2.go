package dummy_2

import (
	"context"
	"dummy_2/pb"
)

type Dummy2Server struct{}

func NewDummy2Server() pb.Dummy2Server {
	return &Dummy2Server{}
}

func (s *Dummy2Server) GetDummy2Method1(ctx context.Context, in *pb.GetDummy2Method1Request) (*pb.Dummy2Method1Object, error) {
	return &pb.Dummy2Method1Object{
		Id:   "dummy_2_id",
		Name: "Dummy 2",
	}, nil
}

func (s *Dummy2Server) GetDummy2Method2(ctx context.Context, in *pb.GetDummy2Method2Request) (*pb.Dummy2Method2Object, error) {
	return &pb.Dummy2Method2Object{
		Id:   "dummy_2_method_2_id",
		Name: "Dummy 2 Method 2",
	}, nil
}
