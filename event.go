package amplitude

type Event struct {
	// A readable ID specified by you. Must have a minimum length of 5 characters. Required unless device_id is present.
	UserID    string  `json:"user_id"`
	// A device-specific identifier, such as the Identifier for Vendor on iOS. Required unless user_id is present. If a device_id is not sent with the event, it will be set to a hashed version of the user_id.
	DeviceID  *string `json:"device_id"`
	// The following event names are reserved for Amplitude use: "[Amplitude] Start Session", "[Amplitude] End Session", "[Amplitude] Revenue", "[Amplitude] Revenue (Verified)", "[Amplitude] Revenue (Unverified)", and "[Amplitude] Merged User".	A unique identifier for your event.
	EventType string  `json:"event_type"`
	// The timestamp of the event in milliseconds since epoch. If time is not sent with the event, it will be set to the request upload time.
	Time      int64   `json:"time"`

	// A dictionary of key-value pairs that represent additional data to be sent along with the event. You can store property values in an array. Date values are transformed into string values. Object depth may not exceed 40 layers.
	EventProperties map[string]interface{} `json:"event_properties"`
	// A dictionary of key-value pairs that represent additional data tied to the user. You can store property values in an array. Date values are transformed into string values. Object depth may not exceed 40 layers.
	UserProperties  map[string]interface{} `json:"user_properties"`
	// This feature is only available to Enterprise customers who have purchased the Accounts add-on. This field adds a dictionary of key-value pairs that represent groups of users to the event as an event-level group. You can only track up to 5 groups.
	Groups          map[string]interface{} `json:"groups"`

	// The current version of your application.
	AppVersion *string `json:"app_version"`
	// Platform of the device.
	Platform   *string `json:"platform"`

	// The name of the mobile operating system or browser that the user is using.
	OsName    *string `json:"os_name"`
	// The version of the mobile operating system or browser the user is using.
	OsVersion *string `json:"os_version"`

	// The device brand that the user is using.
	DeviceBrand        *string `json:"device_brand"`
	// The device manufacturer that the user is using.
	DeviceManufacturer *string `json:"device_manufacturer"`
	// The device model that the user is using.
	DeviceModel        *string `json:"device_model"`

	// The carrier that the user is using.
	Carrier  *string `json:"carrier"`
	// The current country of the user.
	Country  *string `json:"country"`
	// The current region of the user.
	Region   *string `json:"region"`
	// The current city of the user.
	City     *string `json:"city"`
	// The current Designated Market Area of the user.
	Dma      *string `json:"dma"`
	// The language set by the user.
	Language *string `json:"language"`

	// The price of the item purchased. Required for revenue data if the revenue field is not sent. You can use negative values to indicate refunds.
	Price    *float64 `json:"price"`
	// The quantity of the item purchased. Defaults to 1 if not specified.
	Quantity *int64   `json:"quantity"`
	// revneue = price quantity. If you send all 3 fields of price, quantity, and revenue, then (price quantity) will be used as the revenue value. You can use negative values to indicate refunds.
	Revenue  *float64 `json:"revenue"`

	// An identifier for the item purchased. You must send a price and quantity or revenue with this field.
	ProductID   *string `json:"product_id"`
	// The type of revenue for the item purchased. You must send a price and quantity or revenue with this field.
	RevenueType *string `json:"revenue_type"`

	// The current Latitude of the user.
	LocationLat *float32 `json:"location_lat"`
	// The current Longitude of the user.
	LocationLng *float32 `json:"location_lng"`

	// The IP address of the user. Use "$remote" to use the IP address on the upload request. We will use the IP address to reverse lookup a user's location (city, country, region, and DMA). Amplitude has the ability to drop the location and IP address from events once it reaches our servers. You can submit a request to our platform specialist team here to configure this for you.
	Ip *string `json:"ip"`

	// (iOS) Identifier for Advertiser.
	Idfa *string `json:"idfa"`
	// (iOS) Identifier for Vendor.
	Idfv *string `json:"idfv"`
	// (Android) Google Play Services advertising ID
	Adid *string `json:"adid"`

	// (Android) Android ID (not the advertising ID)
	AndroidID *string `json:"android_id"`
	// (Optional) An incrementing counter to distinguish events with the same user_id and timestamp from each other. We recommend you send an event_id, increasing over time, especially if you expect events to occur simultanenously.
	EventID   *int64  `json:"event_id"`
	// (Optional) The start time of the session in milliseconds since epoch (Unix Timestamp), necessary if you want to associate events with a particular system. A session_id of -1 is the same as no session_id specified.
	SessionID *int64  `json:"session_id"`
	// (Optional) A unique identifier for the event. We will deduplicate subsequent events sent with an insert_id we have already seen before within the past 7 days. We recommend generation a UUID or using some combination of device_id, user_id, event_type, event_id, and time.
	InsertID  *int64  `json:"insert_id"`
}