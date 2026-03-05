package service

import (
	"context"
	"html/template"

	pb "github.com/TheDurodola/go-email-service/api/email"
	"gorm.io/gorm"
)


type EmailServer struct {
	DB *gorm.DB
	pb.UnimplementedEmailServiceServer
}

func (s *EmailServer) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	template := req.TemplateName
	receipitent := req.To
	

	
}


func (s *EmailServer) AddEmailTemplate(ctx context.Context, req *pb.EmailTemplateRequest) (*pb.EmailTemplateResponse, error) {

}