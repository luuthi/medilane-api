package order

type StatusOrder int

const (
	Draft StatusOrder = iota
	Confirm
	Confirmed
	Processing
	Packaging
	Delivery
	Delivered
	Received
	Sell
	Sent
	Cancel
)

func (s StatusOrder) String() string {
	return [...]string{"draft", "confirm", "confirmed", "processing", "packaging", "delivery", "delivered", "received", "sell", "sent", "cancel"}[s]
}
