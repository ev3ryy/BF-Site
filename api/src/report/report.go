package report

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"BF/src/database"
	"BF/src/types"
)

// GET handler
func GetReportsHandler(w http.ResponseWriter, r *http.Request) {
	var reports []types.Report
	if err := database.DB.Preload("Role").Find(&reports).Error; err != nil {
		http.Error(w, "failed to fetch reports", http.StatusInternalServerError)
		return
	}

	var response []types.ReportResponse
	for _, report := range reports {
		response = append(response, types.ReportResponse{
			ID:          report.ID,
			AdminName:   report.AdminName,
			RoleName:    report.Role.Name,
			PunishDate:  report.PunishDate.Format("2006-01-02"),
			Description: report.Description,
			Evidence:    report.Evidence,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// POST handler
func ReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("this method (%s) is unsupported", r.Method), http.StatusMethodNotAllowed)
		return
	}

	var input types.ReportInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request input", http.StatusBadRequest)
		return
	}

	punishDate, err := time.Parse("2006-01-02", input.PunishDate)
	if err != nil {
		http.Error(w, "invalid data format. use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	var role types.Role
	if err := database.DB.FirstOrCreate(&role, types.Role{
		Name: input.RoleName,
	}).Error; err != nil {
		http.Error(w, "failed to find or create role", http.StatusInternalServerError)
		return
	}

	report := types.Report{
		AdminName:   input.AdminName,
		RoleID:      role.ID,
		PunishDate:  punishDate,
		Description: input.Description,
		Evidence:    input.Evidence,
	}

	if err := database.DB.Create(&report).Error; err != nil {
		http.Error(w, "failed to create report", http.StatusInternalServerError)
		return
	}

	reportOutput := types.ReportOutput{
		ID:      report.ID,
		Message: "report submitted",
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reportOutput)
}
