package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gitlab.com/alexalreadytaken/notify-service/internal/models/rest"
	"gitlab.com/alexalreadytaken/notify-service/internal/repos"
	"gitlab.com/alexalreadytaken/notify-service/internal/services"
	"gitlab.com/alexalreadytaken/notify-service/internal/utils"
)

type NotifyerControllerTestSuite struct {
	suite.Suite
	gin *gin.Engine
}

func TestNotifyController(t *testing.T) {
	suite.Run(t, &NotifyerControllerTestSuite{})
}

func (s *NotifyerControllerTestSuite) SetupSuite() {
	gin := gin.Default()
	cnf := utils.LoadAppConfigFromEnv()
	pgDb, err := repos.NewDb(cnf)
	if err != nil {
		log.Fatalf("can`t open db connection=%s", err.Error())
	}
	notifyerRepo, err := repos.NewPgNotifyerRepo(pgDb)
	if err != nil {
		log.Fatalf("can`t create notifyer repo=%s", err.Error())
	}
	scheduler := services.NewSchedulerService(notifyerRepo, cnf)
	notifyerController := NewNotifyerController(notifyerRepo, scheduler)
	apiGroup := gin.Group("/api")
	AddNotifyerRoutes(apiGroup, notifyerController)
	s.gin = gin
}

func (s *NotifyerControllerTestSuite) TestPositiveClientCreating() {
	client := newValidClient()
	resp := s.newClientRequest(client)
	s.Equal(200, resp.Code)
}

func (s *NotifyerControllerTestSuite) TestCreateClientInvalidPhoneNumber() {
	client := newValidClient()
	client.PhoneNumber = "xxxxxxxxxxx"
	resp := s.newClientRequest(client)
	s.Equal(400, resp.Code)
	res := s.restResultToStrust(resp)
	s.Equal("phone number must be a digit", res.Msg)
	client.PhoneNumber = "11111111111"
	resp = s.newClientRequest(client)
	s.Equal(400, resp.Code)
	res = s.restResultToStrust(resp)
	s.Equal("phone number must starts at 7", res.Msg)
}

func (s *NotifyerControllerTestSuite) TestCreateClientInvalidOperatorCode() {
	client := newValidClient()
	client.MobileOperatorCode = "amoma"
	resp := s.newClientRequest(client)
	s.Equal(400, resp.Code)
	res := s.restResultToStrust(resp)
	s.Equal("mobile operator code must be a digit", res)
}

func (s *NotifyerControllerTestSuite) TestCreateClientInvalidTimezone() {
	client := newValidClient()
	client.Timezone = "aaaaaaaaaaa"
	resp := s.newClientRequest(client)
	s.Equal(400, resp.Code)
}

func (s *NotifyerControllerTestSuite) newClientRequest(
	client *rest.Client) *httptest.ResponseRecorder {
	resp := httptest.NewRecorder()
	body, err := bodyToReader(client)
	s.NoError(err, "error while creating client")
	s.NotNil(client, "created user is nil")
	req, err := http.NewRequest("POST", "/api/client/", body)
	req.Header.Add("Content-Type", "application/json")
	s.NoError(err, "error while creating client request")
	s.gin.ServeHTTP(resp, req)
	return resp
}

func (s *NotifyerControllerTestSuite) restResultToStrust(
	resp *httptest.ResponseRecorder) *rest.Result {
	var res rest.Result
	err := json.Unmarshal(resp.Body.Bytes(), &res)
	s.NoError(err)
	return &res
}

func newValidClient() *rest.Client {
	return &rest.Client{
		PhoneNumber:        "79999999999",
		MobileOperatorCode: "111",
		Tag:                "manager",
		Timezone:           "Asia/Yekaterinburg",
	}
}

func bodyToReader(body interface{}) (io.Reader, error) {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(bodyJson), nil
}
