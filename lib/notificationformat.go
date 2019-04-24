package lib

// NotificationFormat is the format of a notification
type NotificationFormat string

// GetContentType returns Content-Type
// associated with NotificationFormat
func (f NotificationFormat) GetContentType() string {
	switch f {
	case Template,
		AppleFormat,
		GcmFormat,
		KindleFormat,
		BaiduFormat:
		return "application/json"
	}

	return "application/xml"
}

// IsValid identifies whether notification format is valid
func (f NotificationFormat) IsValid() bool {
	return f == Template ||
		f == GcmFormat ||
		f == AppleFormat ||
		f == BaiduFormat ||
		f == KindleFormat ||
		f == WindowsFormat ||
		f == WindowsPhoneFormat
}