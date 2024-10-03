package purchase
type Status string
const (
    Status_Created Status = "Created"
    Status_Processing Status = "Processing"
    Status_Failed Status = "Failed"
    Status_PaymentProcessing Status = "PaymentProcessing"
    Status_Complete Status = "Complete"
)