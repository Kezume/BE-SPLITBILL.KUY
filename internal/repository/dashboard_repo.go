package repository

type DashboardRepository interface {
	// Akan diisi query database nanti (Get Summary, Active Groups, dll)
}

type dashboardRepo struct{}

func NewDashboardRepository() DashboardRepository {
	return &dashboardRepo{}
}
