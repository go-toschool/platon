package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/go-toschool/platon/database"
	"github.com/go-toschool/platon/entities"
	"github.com/go-toschool/platon/service"

	"github.com/go-toschool/platon/talks"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	port := flag.Int64("port", 8004, "listening port")
	postgresDSN := flag.String("postgres-dsn", "postgres://gotoschool:goto1234@localhost:5432/drachma?sslmode=disable", "Postgres DSN")

	flag.Parse()

	pgSvc, err := database.NewPostgres(*postgresDSN)
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	ts := &TalksService{
		Talks: service.NewTalksService(pgSvc),
	}
	srv := grpc.NewServer()
	talks.RegisterTalkingServer(srv, ts)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println(fmt.Sprintf("%s, Listening on: %d", ts.String(), *port))
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// TalksService ...
type TalksService struct {
	Talks entities.Talks
}

func (ts *TalksService) String() string {
	return "GRPC Talks Service"
}

// Get ...
func (ts *TalksService) Get(ctx context.Context, gr *talks.GetRequest) (*talks.GetResponse, error) {
	t, err := ts.Talks.GetByIDAndUserID(gr.GetTalkId(), gr.GetUserId())
	if err != nil {
		return nil, err
	}

	return &talks.GetResponse{
		Talk: t.ToProto(),
	}, nil
}

// Select ...
func (ts *TalksService) Select(ctx context.Context, sr *talks.SelectRequest) (*talks.SelectResponse, error) {
	tt, err := ts.Talks.Select()
	if err != nil {
		return nil, err
	}

	data := make([]*talks.Talk, 0)
	for _, t := range tt {
		data = append(data, t.ToProto())
	}

	return &talks.SelectResponse{
		Talk: data,
	}, nil
}

// Create ...
func (ts *TalksService) Create(ctx context.Context, cr *talks.CreateRequest) (*talks.CreateResponse, error) {
	talk, err := ts.Talks.GetByID(cr.GetTalk().GetId())
	if err != nil {
		t := (&entities.Talk{}).FromProto(cr.Talk)

		if err := ts.Talks.Create(t); err != nil {
			return nil, err
		}

		return &talks.CreateResponse{
			Talk: t.ToProto(),
		}, nil
	}

	return &talks.CreateResponse{
		Talk: talk.ToProto(),
	}, nil
}

// Update ...
func (ts *TalksService) Update(ctx context.Context, ur *talks.UpdateRequest) (*talks.UpdateResponse, error) {
	talk, err := ts.Talks.GetByID(ur.GetData().GetId())
	if err != nil {
		t := (&entities.Talk{}).FromProto(ur.Data)

		if err := ts.Talks.Update(t); err != nil {
			return nil, err
		}

		return &talks.UpdateResponse{
			Data: t.ToProto(),
		}, nil
	}

	return &talks.UpdateResponse{
		Data: talk.ToProto(),
	}, nil
}

// Delete ...
func (ts *TalksService) Delete(ctx context.Context, dr *talks.DeleteRequest) (*talks.DeleteResponse, error) {
	talk, err := ts.Talks.GetByID(dr.GetData().GetId())
	if err != nil {
		return nil, err
	}

	if err := ts.Talks.Update(talk); err != nil {
		return nil, err
	}

	return &talks.DeleteResponse{
		Data: talk.ToProto(),
	}, nil
}
