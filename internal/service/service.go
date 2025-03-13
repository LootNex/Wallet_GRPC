package service

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"github.com/LootNex/Wallet_GRPC.git/pkg/api"
)

type Server struct {
	api.UnimplementedWalletsServer
	mu sync.Mutex
	connection *sql.DB
}

func NewServer(db *sql.DB) *Server{
	return &Server{
		connection: db,
	}
}

func (s *Server) CreateWallet(ctx context.Context, req *api.CreateWalletRequest) (*api.CreateWalletResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_,err := s.connection.Exec("INSERT INTO wallet(Id, Balance) VALUES ($1, $2)", req.Id, req.Balance)
	if err != nil{
		log.Println("Data wasnot added to the database")
	}
	return &api.CreateWalletResponse{
		Id: req.Id,
		Balance: &api.Wallet_Balance{Balance: req.Balance},
	},nil
}

func (s *Server) GetBalance(ctx context.Context, req *api.GetBalanceRequest) (*api.GetBalanceResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var balance int
	s.connection.QueryRow("SELECT Balance FROM wallet WHERE Id = $1", req.Id).Scan(&balance)
	return &api.GetBalanceResponse{
		Balance: &api.Wallet_Balance{Balance: int64(balance)},
	}, nil
}

func (s *Server) UpdateBalance(ctx context.Context, req *api.UpdateBalaceRequest) (*api.UpdateBalaceResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_,err := s.connection.Exec("UPDATE wallet SET Balance = $1 WHERE Id = $2", req.Balance, req.Id)
	if err != nil{
		log.Println("Data wasnot added to the database")
	}

	return &api.UpdateBalaceResponse{
		Balance: &api.Wallet_Balance{Balance: req.Balance},
	},nil
}
