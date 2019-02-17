package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/go-toschool/platon/entities"
	"github.com/go-toschool/platon/postgres"
	"github.com/jmoiron/sqlx"

	"github.com/go-toschool/platon/talks"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	port := flag.Int64("port", 8004, "listening port")
	postgresDSN := flag.String("postgres-dsn", "postgres://gotoschool:goto1234@localhost:5432/drachma?sslmode=disable", "Postgres DSN")

	flag.Parse()

	db, err := sqlx.Connect("postgres", *postgresDSN)
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	ts := &TalksService{
		Talks: &postgres.TalksService{
			Store: db,
		},
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
	t, err := ts.Talks.Get(&entities.TalksQuery{
		ID:     gr.GetTalkId(),
		UserID: gr.GetUserId(),
	})
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
	ot, err := ts.Talks.Get(&entities.TalksQuery{
		ID: cr.Talk.Id,
	})
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
		Talk: ot.ToProto(),
	}, nil
}

// Update ...
func (ts *TalksService) Update(ctx context.Context, ur *talks.UpdateRequest) (*talks.UpdateResponse, error) {
	ot, err := ts.Talks.Get(&entities.TalksQuery{
		ID: ur.Data.Id,
	})
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
		Data: ot.ToProto(),
	}, nil
}

// Delete ...
func (ts *TalksService) Delete(ctx context.Context, dr *talks.DeleteRequest) (*talks.DeleteResponse, error) {
	t, err := ts.Talks.Get(&entities.TalksQuery{
		ID: dr.Data.Id,
	})
	if err != nil {
		return nil, err
	}

	if err := ts.Talks.Update(t); err != nil {
		return nil, err
	}

	return &talks.DeleteResponse{
		Data: t.ToProto(),
	}, nil
}
