package usecase

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"github.com/google/uuid"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"web-shop/domain"
)

type ReportUseCase interface {
	SaveReport(ctx context.Context, reports []domain.StatisticsReport) (domain.Report, error)
	GetAllDocuments(ctx context.Context, agentId string) ([]domain.Report, error)
	ExportToPDF(ctx context.Context, reportId string, agentId string) error
}

type reportUseCase struct {

}

func (r reportUseCase) SaveReport(ctx context.Context, reports []domain.StatisticsReport) (domain.Report, error) {
	var report domain.Report
	report.Timestamp = time.Now()
	report.StatisticReport = reports
	report.ReportId = uuid.NewString()

	workingDirectory, _ := os.Getwd()
	startDirectory := workingDirectory
	if !strings.HasSuffix(workingDirectory, "src") {
		firstPart := strings.Split(workingDirectory, "src")
		value := firstPart[0] + "/src/documents"
		workingDirectory = value
		os.Chdir(workingDirectory)
	}
	path1 :=  report.ReportId + ".xml"
	output, err := xml.MarshalIndent(&report, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	err = ioutil.WriteFile(path1, output, 0755)
	os.Chdir(startDirectory)

	httpClient := &http.Client{}
	var existUri string
	if os.Getenv("DOCKER_ENV") == "" {
		existUri = "http://localhost:8095"
	} else {
		existUri = "http://xmldb:8080"
	}
	req, err := http.NewRequest(http.MethodPut, existUri + "/exist/rest/collection/report" + path1, bytes.NewBuffer(output))
	req.SetBasicAuth("admin", "")
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(output)))
	resp, err := httpClient.Do(req)

	if resp.StatusCode != 201 || err != nil {
		return domain.Report{}, fmt.Errorf("error while saving report")
	}

	return report, nil
}

func (r reportUseCase) GetAllDocuments(ctx context.Context, campaignId string) ([]domain.Report, error) {

	panic("err")
}

func (r reportUseCase) ExportToPDF(ctx context.Context, reportId string, agentId string) error {
	httpClient := &http.Client{}
	var existUri string
	if os.Getenv("DOCKER_ENV") == "" {
		existUri = "http://localhost:8095"
	} else {
		existUri = "http://xmldb:8080"
	}
	req, err := http.NewRequest(http.MethodGet, existUri + "/exist/rest/collection/report" + reportId + ".xml", bytes.NewBuffer([]byte("")))
	req.SetBasicAuth("admin", "")
	resp, err := httpClient.Do(req)
	if resp.StatusCode != 200 || err != nil {
		return fmt.Errorf("error while saving report")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var report domain.Report
	err = xml.Unmarshal(body, &report)
	if err != nil {
		return err
	}
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.Row(20, func() {
		m.Col(4, func() {
			m.Text("Campaigns Report", props.Text{
				Top:         10,
				Size:        15,
				Extrapolate: true,

			})
		})
		m.ColSpace(4)
	})

	for _, campaign := range report.StatisticReport {
		m.Row(10, func() {
			m.Col(4, func() {
				m.Text("Advertisement Reaching Stats Per Campaign", props.Text{
					Size:        12,
					Top:         20,
					Extrapolate: true,
					Style:       consts.Bold,
				})
			})
			m.ColSpace(4)
		})

		m.Row(6, func() {
			m.Col(6, func() {
				m.Text("ID: " + campaign.CampaignId, props.Text{
					Size: 12,
					Top:  20,
				})
			})
			m.ColSpace(6)
		})

		m.Row(6, func() {
			m.Col(4, func() {
				m.Text("Type: " + campaign.CampaignType, props.Text{
					Size: 12,
					Top:  20,
				})
			})
			m.ColSpace(4)
		})

		m.Row(6, func() {
			m.Col(4, func() {
				m.Text("Period: " + campaign.CampaignPeriod, props.Text{
					Size: 12,
					Top:  20,
				})
			})
			m.ColSpace(4)
		})

		m.Row(6, func() {
			m.Col(4, func() {
				m.Text("Frequency: " + campaign.AdvertisementFrequency, props.Text{
					Size: 12,
					Top:  20,
				})
			})
			m.ColSpace(4)
		})


		m.Row(10, func() {
			m.Col(4, func() {
				m.Text("Interaction Statistics:", props.Text{
					Size:        12,
					Top:  20,
					Extrapolate: true,
					Style:       consts.Bold,
				})
			})
			m.ColSpace(4)
		})

		header := []string{"Likes", "Dislikes", "Comments", "Visits"}
		contents := [][]string{
			{   strconv.Itoa(campaign.NumOfLikes),
				strconv.Itoa(campaign.NumOfDislikes),
				strconv.Itoa(campaign.NumOfComments),
				strconv.Itoa(campaign.AdvertisingCount.AdvertisedCount),
			},
		}

		m.Row(10, func() {})


		m.Row(30, func() {
			m.TableList(header, contents, props.TableList{
				HeaderProp: props.TableListContent{
					Size: 12,
				},
				ContentProp: props.TableListContent{
					Size: 10,
				},
				Align:                consts.Left,
				AlternatedBackground: &color.Color{Red: 150, Green: 100, Blue: 200},
				HeaderContentSpace:   1,
				Line:                 true,
			})
		})

		m.Row(10, func() {})

		m.Row(10, func() {
			m.Col(4, func() {
				m.Text("Site Visits From Influencer Post:", props.Text{
					Size:        12,
					Top:  20,
					Extrapolate: true,
					Style:       consts.Bold,
				})
			})
			m.ColSpace(6)
		})

		header2 := []string{"Influencer Username", "Visits"}
		contents2 := [][]string{}

		for _, item := range campaign.Clicks {
			var cont []string
			cont = append(cont, item.InfluencerUsername)
			cont = append(cont, strconv.Itoa(item.NumOfClicks))
			contents2 = append(contents2, cont)
		}

		m.Row(10, func() {})

		m.Row(30, func() {
			m.TableList(header2, contents2, props.TableList{
				HeaderProp: props.TableListContent{
					Size: 12,
				},
				ContentProp: props.TableListContent{
					Size: 10,
				},
				Align:                consts.Left,
				AlternatedBackground: &color.Color{Red: 150, Green: 100, Blue: 200},
				HeaderContentSpace:   1,
				Line:                 true,
			})
		})

		m.Row(10, func() {})


		m.Line(10)

	}


	m.Row(10, func() {
		m.Col(4, func() {
			m.Text("Report Generated: " + report.Timestamp.Format("02.04.2006 15:04"), props.Text{
				Size:        12,
				Top:         22,
				Extrapolate: true,
				Style:       consts.Bold,
			})
		})
		m.ColSpace(4)
	})

	workingDirectory, _ := os.Getwd()
	startDirectory := workingDirectory
	if !strings.HasSuffix(workingDirectory, "src") {
		firstPart := strings.Split(workingDirectory, "src")
		value := firstPart[0] + "/src/documents"
		workingDirectory = value
		os.Chdir(workingDirectory)
	}

	err = m.OutputFileAndClose(reportId + ".pdf")
	if err != nil {
		return err
	}
	os.Chdir(startDirectory)
	return nil
}

func NewReportUseCase() ReportUseCase {
	return &reportUseCase{}
}