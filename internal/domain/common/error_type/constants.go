package error_type

type ErrorType string

const (
	PriceAlertIncorrectIndex       ErrorType = "price_alert_incorrect_index"
	IncorrectSubscriptionStorageId ErrorType = "incorrect_subscription_storage_id"
	PriceAlertUpdateError          ErrorType = "price_alert_update_error"
	SubscriptionCreationError      ErrorType = "subscription_creation_error"
	Internal                       ErrorType = "internal"
)
