package question_bundle
type BundleType string
type Status string
const (
    BundleType_Manual BundleType = "Manual"
    BundleType_ByCategory BundleType = "ByCategory"
    BundleType_ByTag BundleType = "ByTag"
    Status_Concept Status = "Concept"
    Status_Approved Status = "Approved"
    Status_Published Status = "Published"
)