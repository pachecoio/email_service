package _interface

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/pachecoio/email_service/adapters/mailgun"
	"github.com/pachecoio/email_service/domain"
	"github.com/pachecoio/email_service/services"
	"net/http"
	"os"
)

func SendEmailHandler(w http.ResponseWriter, r *http.Request) {
	//	Load environment variables
	err := godotenv.Load(".env")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	domainValue := os.Getenv("MAILGUN_DOMAIN")
	apiKeyValue := os.Getenv("MAILGUN_API_KEY")

	//	Send mail with mailgun
	c := mailgun.NewClient(domainValue, apiKeyValue)
	var payload domain.EmailPayload

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid payload: "+err.Error(), http.StatusBadRequest)
	}
	res := services.SendEmail(c, &payload)
	if res.Failed {
		http.Error(w, res.Message, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func SendBatchEmailHandler(w http.ResponseWriter, r *http.Request) {
	//	Load environment variables
	err := godotenv.Load(".env")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	domainValue := os.Getenv("MAILGUN_DOMAIN")
	apiKeyValue := os.Getenv("MAILGUN_API_KEY")

	//	Send mail with mailgun
	c := mailgun.NewClient(domainValue, apiKeyValue)

	var payload domain.EmailBatchPayload

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid payload: "+err.Error(), http.StatusBadRequest)
	}
	res := services.SendBatch(c, &payload)

	encoded, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(encoded)
}
