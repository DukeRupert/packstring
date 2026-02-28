package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/firefly/packstring/internal/data"
	"github.com/firefly/packstring/internal/db"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"github.com/stripe/stripe-go/v81/webhook"
)

// CreateCheckoutSession creates a Stripe Checkout session for a deposit payment.
func CreateCheckoutSession(customerEmail string, amountCents int, tripName, successURL, cancelURL string, inquiryID int64) (checkoutURL, sessionID string, err error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	if stripe.Key == "" {
		return "", "", fmt.Errorf("STRIPE_SECRET_KEY not set")
	}

	params := &stripe.CheckoutSessionParams{
		CustomerEmail: stripe.String(customerEmail),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name:        stripe.String(fmt.Sprintf("Deposit — %s", tripName)),
						Description: stripe.String("Trip deposit for MT Hunt & Fish Outfitters"),
					},
					UnitAmount: stripe.Int64(int64(amountCents)),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(successURL + "?session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String(cancelURL),
		Metadata: map[string]string{
			"inquiry_id": fmt.Sprintf("%d", inquiryID),
		},
	}

	s, err := session.New(params)
	if err != nil {
		return "", "", fmt.Errorf("create checkout session: %w", err)
	}

	return s.URL, s.ID, nil
}

// StripeHandler handles Stripe webhook events.
type StripeHandler struct {
	store *db.Store
}

// NewStripeHandler creates a new Stripe webhook handler.
func NewStripeHandler(store *db.Store) *StripeHandler {
	return &StripeHandler{store: store}
}

// HandleWebhook processes incoming Stripe webhook events.
func (h *StripeHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	const maxBodyBytes = 65536
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("[stripe] read body error: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	if endpointSecret == "" {
		log.Println("[stripe] STRIPE_WEBHOOK_SECRET not set, skipping signature verification")
		// In dev, still process the event
	}

	var event stripe.Event
	if endpointSecret != "" {
		event, err = webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"), endpointSecret)
		if err != nil {
			log.Printf("[stripe] signature verification failed: %v", err)
			http.Error(w, "Invalid signature", http.StatusBadRequest)
			return
		}
	} else {
		if err := json.Unmarshal(payload, &event); err != nil {
			log.Printf("[stripe] unmarshal error: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	}

	switch event.Type {
	case "checkout.session.completed":
		var cs stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &cs); err != nil {
			log.Printf("[stripe] unmarshal session: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		log.Printf("[stripe] checkout.session.completed: %s (payment_intent: %s)", cs.ID, cs.PaymentIntent.ID)

		if err := h.store.UpdatePaymentStatus(cs.ID, "paid", cs.PaymentIntent.ID); err != nil {
			log.Printf("[stripe] update payment status error: %v", err)
		}

	case "checkout.session.expired":
		var cs stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &cs); err != nil {
			log.Printf("[stripe] unmarshal session: %v", err)
			break
		}
		log.Printf("[stripe] checkout.session.expired: %s", cs.ID)
		if err := h.store.UpdatePaymentStatus(cs.ID, "failed", ""); err != nil {
			log.Printf("[stripe] update payment status error: %v", err)
		}

	default:
		log.Printf("[stripe] unhandled event type: %s", event.Type)
	}

	w.WriteHeader(http.StatusOK)
}

// PaymentSuccess renders the public payment success page.
func PaymentSuccess(templates map[string]*template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d := map[string]any{
			"Meta": data.PageMeta{Title: "Payment Received — MT Hunt & Fish Outfitters"},
		}
		if err := templates["payment-success"].ExecuteTemplate(w, "base.html", d); err != nil {
			log.Printf("Error rendering payment success: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// PaymentCancel renders the public payment cancelled page.
func PaymentCancel(templates map[string]*template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d := map[string]any{
			"Meta": data.PageMeta{Title: "Payment Cancelled — MT Hunt & Fish Outfitters"},
		}
		if err := templates["payment-cancel"].ExecuteTemplate(w, "base.html", d); err != nil {
			log.Printf("Error rendering payment cancel: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
