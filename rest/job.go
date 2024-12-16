package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/gommon/log"
	infrastructureJob "watchdog-go.com/internal/infrastructure/grpc"
	pb "watchdog-go.com/internal/infrastructure/grpc/generated"
)

type MakeProposalRequest struct {
	JobId    string `json:"job_id"`
	LawyerId string `json:"lawyer_id"`
	Value    string `json:"value"`
}

func HandleJob(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		makeProposal(w, r)
	}
}

func makeProposal(w http.ResponseWriter, r *http.Request) {

	log.Print("it's here")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	var proposal MakeProposalRequest

	err = json.Unmarshal(body, &proposal)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	valueFloatNumber, _ := strconv.ParseFloat(proposal.Value, 64)
	req := &pb.ProposalRequest{
		ContractorId: "32",
		JobId:        proposal.JobId,
		LawyerId:     proposal.LawyerId,
		Value:        valueFloatNumber,
	}
	err = infrastructureJob.MakeAnOffer(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

}
