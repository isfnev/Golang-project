package types

type UserGroup string

const (
	GroupAdmin               UserGroup = "GroupAdmin"
	GroupDealerFinance       UserGroup = "GroupDealerFinance"
	GroupSalesManagement     UserGroup = "GroupSalesManagement"
	GroupEveryOne            UserGroup = "GroupEveryOne"
	GroupDb                  UserGroup = "GroupDb"
	GroupAdminDealer         UserGroup = "GroupAdminDealer"
	GroupAdminAccounts       UserGroup = "GroupAdminAccounts"
	GroupAdminDealerAccounts UserGroup = "GroupAdminDealerAccounts"
	GroupDealerAccounts      UserGroup = "GroupDealerAccounts"
	GroupDatEveryOne         UserGroup = "GroupDatEveryOne"
	GroupDatAdmin            UserGroup = "GroupDatAdmin"
	GroupDatAdminAccounts    UserGroup = "GroupDatAdminAccounts"
	GroupRevenueAdminDealer  UserGroup = "GroupRevenueAdminDealer"
)