package account

import "hotel-management/address"

type AccountType string

const (
	ServerAccountType       AccountType = "server"
	HouseKeepingAccountType AccountType = "housekeeping"
	ReceptionistAccountType AccountType = "receptionist"
	ManagerAccountType      AccountType = "manager"
	GuestAccountType        AccountType = "guest"
)

type AccountStatus string

const (
	AccountStatusActive      AccountStatus = "active"
	AccountStatusInactive    AccountStatus = "inactive"
	AccountStatusBlackListed AccountStatus = "blacklisted"
)

type Account interface {
	GetID() string
	GetName() string
	GetAddress() address.Address
	GetEmail() string
	GetPhoneNumber() string
	GetAccountType() AccountType
}

type ServerAccount struct {
	ID          string
	Name        string
	Address     address.Address
	Email       string
	Phone       string
	AccountType AccountType
}

func (a *ServerAccount) GetID() string {
	return a.ID
}

func (a *ServerAccount) GetName() string {
	return a.Name
}

func (a *ServerAccount) GetAddress() address.Address {
	return a.Address
}

func (a *ServerAccount) GetEmail() string {
	return a.Email
}

func (a *ServerAccount) GetPhoneNumber() string {
	return a.Phone
}

func (a *ServerAccount) GetAccountType() AccountType {
	return ServerAccountType
}

type HouseKeepingAccount struct {
	ID          string
	Name        string
	Address     address.Address
	Email       string
	Phone       string
	AccountType AccountType
}

func (a *HouseKeepingAccount) GetID() string {
	return a.ID
}

func (a *HouseKeepingAccount) GetName() string {
	return a.Name
}

func (a *HouseKeepingAccount) GetAddress() address.Address {
	return a.Address
}

func (a *HouseKeepingAccount) GetEmail() string {
	return a.Email
}

func (a *HouseKeepingAccount) GetPhoneNumber() string {
	return a.Phone
}

func (a *HouseKeepingAccount) GetAccountType() AccountType {
	return HouseKeepingAccountType
}

type ReceptionistsAccount struct {
	ID          string
	Name        string
	Address     address.Address
	Email       string
	Phone       string
	AccountType AccountType
}

func (a *ReceptionistsAccount) GetID() string {
	return a.ID
}

func (a *ReceptionistsAccount) GetName() string {
	return a.Name
}

func (a *ReceptionistsAccount) GetAddress() address.Address {
	return a.Address
}

func (a *ReceptionistsAccount) GetEmail() string {
	return a.Email
}

func (a *ReceptionistsAccount) GetPhoneNumber() string {
	return a.Phone
}

func (a *ReceptionistsAccount) GetAccountType() AccountType {
	return ReceptionistAccountType
}

type ManagerAccount struct {
	ID          string
	Name        string
	Address     address.Address
	Email       string
	Phone       string
	AccountType AccountType
}

func (a *ManagerAccount) GetID() string {
	return a.ID
}

func (a *ManagerAccount) GetName() string {
	return a.Name
}

func (a *ManagerAccount) GetAddress() address.Address {
	return a.Address
}

func (a *ManagerAccount) GetEmail() string {
	return a.Email
}

func (a *ManagerAccount) GetPhoneNumber() string {
	return a.Phone
}

func (a *ManagerAccount) GetAccountType() AccountType {
	return ManagerAccountType
}

type GuestAccount struct {
	ID          string
	Name        string
	Address     address.Address
	Email       string
	Phone       string
	AccountType AccountType
}

func (a *GuestAccount) GetID() string {
	return a.ID
}

func (a *GuestAccount) GetName() string {
	return a.Name
}

func (a *GuestAccount) GetAddress() address.Address {
	return a.Address
}

func (a *GuestAccount) GetEmail() string {
	return a.Email
}

func (a *GuestAccount) GetPhoneNumber() string {
	return a.Phone
}

func (a *GuestAccount) GetAccountType() AccountType {
	return GuestAccountType
}

func NewAccount(accountType AccountType, id, name, email, phone string, address address.Address) Account {
	switch accountType {
	case ManagerAccountType:
		return &ManagerAccount{ID: id, Name: name, Address: address, Email: email, Phone: phone, AccountType: accountType}
	case HouseKeepingAccountType:
		return &HouseKeepingAccount{ID: id, Name: name, Address: address, Email: email, Phone: phone, AccountType: accountType}
	case ServerAccountType:
		return &ServerAccount{ID: id, Name: name, Address: address, Email: email, Phone: phone, AccountType: accountType}
	case GuestAccountType:
		return &GuestAccount{ID: id, Name: name, Address: address, Email: email, Phone: phone, AccountType: accountType}
	case ReceptionistAccountType:
		return &ReceptionistsAccount{ID: id, Name: name, Address: address, Email: email, Phone: phone, AccountType: accountType}
	default:
		return nil
	}
}
