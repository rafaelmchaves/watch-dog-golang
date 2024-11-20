package grpc

import (
	"context"
	"log"
	"time"

	pb "watchdog-go.com/internal/infrastructure/grpc/generated"
)

func MakeAnOffer(proposalRequest *pb.ProposalRequest) error {

	connection := GetClientConn()
	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := pb.NewContractorClient(connection)

	_, err := client.CreateJobProposal(ctx, proposalRequest)
	if err != nil {
		log.Printf("Error calling CreateJobProposal: %v", err)
		return err
	}
	// Print the response
	log.Println("Received MakeAnOffer response ok")
	return nil
}
