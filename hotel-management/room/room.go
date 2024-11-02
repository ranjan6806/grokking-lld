package room

type RoomType string

const (
	RoomTypeStandard      RoomType = "standard"
	RoomTypeDeluxe        RoomType = "deluxe"
	RoomTypeFamilySuite   RoomType = "family-suite"
	RoomTypeBusinessSuite RoomType = "business-suite"
)

type RoomStatus string

const (
	RoomStatusAvailable     RoomStatus = "available"
	RoomStatusReserved      RoomStatus = "reserved"
	RoomStatusOccupied      RoomStatus = "occupied"
	RoomStatusBeingServiced RoomStatus = "being-serviced"
)

type Room interface {
	GetRoomNumber() string
	IsAvailable() bool
	Checkin() error
	Checkout() error
}

type RoomStandard struct {
	RoomNumber   string
	RoomType     RoomType
	RoomStatus   RoomStatus
	bookingPrice float64
}

func (r *RoomStandard) GetRoomNumber() string {
	return r.RoomNumber
}

func (r *RoomStandard) IsAvailable() bool {
	return r.RoomStatus == RoomStatusAvailable
}

func (r *RoomStandard) Checkin() error {
	return nil
}

func (r *RoomStandard) Checkout() error {
	return nil
}

type RoomDeluxe struct {
	RoomNumber   string
	RoomType     RoomType
	RoomStatus   RoomStatus
	bookingPrice float64
}

func (r *RoomDeluxe) GetRoomNumber() string {
	return r.RoomNumber
}

func (r *RoomDeluxe) IsAvailable() bool {
	return r.RoomStatus == RoomStatusAvailable
}

func (r *RoomDeluxe) Checkin() error {
	return nil
}

func (r *RoomDeluxe) Checkout() error {
	return nil
}

type RoomFamilySuite struct {
	RoomNumber   string
	RoomType     RoomType
	RoomStatus   RoomStatus
	bookingPrice float64
}

func (r *RoomFamilySuite) GetRoomNumber() string {
	return r.RoomNumber
}

func (r *RoomFamilySuite) IsAvailable() bool {
	return r.RoomStatus == RoomStatusAvailable
}

func (r *RoomFamilySuite) Checkin() error {
	return nil
}

func (r *RoomFamilySuite) Checkout() error {
	return nil
}

type RoomDeluxeSuite struct {
	RoomNumber   string
	RoomType     RoomType
	RoomStatus   RoomStatus
	bookingPrice float64
}

func (r *RoomDeluxeSuite) GetRoomNumber() string {
	return r.RoomNumber
}

func (r *RoomDeluxeSuite) IsAvailable() bool {
	return r.RoomStatus == RoomStatusAvailable
}

func (r *RoomDeluxeSuite) Checkin() error {
	return nil
}

func (r *RoomDeluxeSuite) Checkout() error {
	return nil
}

func NewRoom(roomType RoomType, roomNumber string, roomStatus RoomStatus, bookingPrice float64) Room {
	return &RoomStandard{
		RoomNumber:   roomNumber,
		bookingPrice: bookingPrice,
	}
}
