package service

import (
	"context"

	pb "github.com/TheDurodola/go-email-service/api/email"
	"github.com/TheDurodola/go-email-service/internal/data/models"
	"github.com/TheDurodola/go-email-service/internal/data/repositories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


type EmailServer struct {
	Repo repositories.EmailRepository
	pb.UnimplementedEmailServiceServer
}

func (s *EmailServer) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	firstname := req.GetFirstname()
	templateName := req.GetTemplateName()
	receipient := req.GetTo()
	emailLog := &models.OutgoingEmail{
		TemplateName: templateName,
		AppName: req.AppName,
		Recipient: receipient,
		FirstName: firstname,
	}

	if err := s.Repo.CreateEmailLog(emailLog); err != nil {
		return nil, status.Error(codes.Internal, "failed to persist email log")
	}

	return &pb.SendEmailResponse{
		IsSent: true,
	}, nil

}


func (s *EmailServer) AddEmailTemplate(ctx context.Context, req *pb.EmailTemplateRequest) (*pb.EmailTemplateResponse, error) {
	

	template := &models.EmailTemplate{
		Type: models.TemplateType(req.GetTemplateType()),
		Name: req.GetTemplateName(),
		Body: req.GetTemplateBody(),
	}


	if err := s.Repo.CreateEmailTemplate(template); err != nil {
		return nil, err
	}

	return &pb.EmailTemplateResponse{
		Message: "Template added successfully",
	}, nil


}