package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	pb "github.com/TheDurodola/go-email-service/api/email"
	"github.com/TheDurodola/go-email-service/internal/data/models"
	"github.com/TheDurodola/go-email-service/internal/data/repositories"
	"github.com/joho/godotenv"
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

	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using system environment variables")
    }


	BREVO_API_KEY := os.Getenv("BREVO_API_KEY")
	url := os.Getenv("BREVO_URL")


	payload := strings.NewReader("{\n  \"htmlContent\": \"<html><head></head><body><p>Hello,</p>This is my first transactional email sent from Brevo.</p></body></html>\",\n  \"sender\": {\n    \"email\": \"" + os.Getenv("SENDER_EMAIL") + "\",\n    \"name\": \"Bolaji from YRSD\"\n  },\n  \"subject\": \"Hello from YRSD!\",\n  \"to\": [\n    {\n      \"email\": \"" + receipient + "\",\n      \"name\": \"" + firstname + "\"\n    }\n  ]\n}")
	
	req2, _ := http.NewRequest("POST", url, payload)
	req2.Header.Add("api-key", BREVO_API_KEY)
	req2.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req2)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to send email")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	
	fmt.Println(res)
	fmt.Println(string(body))

	if err := s.Repo.CreateLog(emailLog); err != nil {
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


	if err := s.Repo.CreateTemplate(template); err != nil {
		return nil, status.Error(codes.Internal, "failed to persist email template")
	}

	return &pb.EmailTemplateResponse{
		Message: "Template added successfully",
	}, nil


}