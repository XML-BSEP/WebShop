package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"web-shop/domain"
	"web-shop/http/middleware"
	"web-shop/infrastructure/dto"
	"web-shop/infrastructure/persistance/datastore"
	"web-shop/usecase"
)

type CampaignHandler interface {
	SaveToken(c echo.Context) error
	CreateAd(c echo.Context) error
	CreateDisposableCampaign(c echo.Context) error
	CreateMultipleCampaign(c echo.Context) error
	GetAllDisposableCampaigns(c echo.Context) error
	GetAllMultipleCampaigns(c echo.Context) error
	UpdateMultipleCampaign(c echo.Context) error
	DeleteMultipleCampaign(c echo.Context) error
	DeleteDisposableCampaign(c echo.Context) error
	GetAllAdsPerAgent(c echo.Context) error
	GenerateStatisticsReport(c echo.Context) error
	DownloadPdf(c echo.Context) error


}


type campaignHandler struct {
	tokenRepository datastore.TokenRepository
	reportUseCase usecase.ReportUseCase
}

func (c2 campaignHandler) DownloadPdf(c echo.Context) error {
	var token *dto.GeneratePdf

	decoder := json.NewDecoder(c.Request().Body)
	err := decoder.Decode(&token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, invalidJson)
	}

	err = c2.reportUseCase.ExportToPDF(context.Background(), token.ReportId, "")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error")
	}
	return c.JSON(200, "ok")
}

func (c2 campaignHandler) GenerateStatisticsReport(c echo.Context) error {
	userId, _ := middleware.ExtractUserId(c.Request())
	token, err := c2.tokenRepository.FetchToken(context.Background(), userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error")
	}
	var retVal []domain.StatisticsReport
	client := resty.New()

	domain := os.Getenv("NISHTAGRAM_DOMAIN")
	if domain == "" {
		domain = "127.0.0.1"
	}
	if os.Getenv("DOCKER_ENV") == "" {
		resp, _ := client.R().
			SetHeader("Authorization", token).
			EnableTrace().
			Get("https://" + domain + ":8080/generateStatisticReport")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}
		fmt.Println(string(resp.Body()))
		json.Unmarshal(resp.Body(), &retVal)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, invalidJson)
		}
		report, err := c2.reportUseCase.SaveReport(context.Background(), retVal)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}
		return c.JSON(200, report)

	} else {
		resp, _ := client.R().
			SetHeader("Authorization", token).
			EnableTrace().
			Get("http://" + domain + ":8093/ad/generateStatisticReport")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}

		json.Unmarshal(resp.Body(), &retVal)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, invalidJson)
		}
		report, err := c2.reportUseCase.SaveReport(context.Background(), retVal)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}
		return c.JSON(200, report)
	}
}

func (c2 campaignHandler) GetAllAdsPerAgent(c echo.Context) error {
	userId, _ := middleware.ExtractUserId(c.Request())
	token, err := c2.tokenRepository.FetchToken(context.Background(), userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error")
	}
	var retVal []domain.AdPost
	client := resty.New()

	domain := os.Getenv("NISHTAGRAM_DOMAIN")
	if domain == "" {
		domain = "127.0.0.1"
	}
	if os.Getenv("DOCKER_ENV") == "" {
		resp, _ := client.R().
			SetHeader("Authorization", token).
			EnableTrace().
			Get("https://" + domain + ":8080/getAdsByAgent")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}
		fmt.Println(string(resp.Body()))
		json.Unmarshal(resp.Body(), &retVal)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, invalidJson)
		}
		return c.JSON(200, retVal)

	} else {
		resp, _ := client.R().
			SetHeader("Authorization", token).
			EnableTrace().
			Get("http://" + domain + ":8093/ad/getAdsByAgent")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}

		json.Unmarshal(resp.Body(), &retVal)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, invalidJson)
		}
		return c.JSON(200, retVal)
	}
}

func (c2 campaignHandler) DeleteDisposableCampaign(c echo.Context) error {

	userId, _ := middleware.ExtractUserId(c.Request())
	token, err := c2.tokenRepository.FetchToken(context.Background(), userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error")
	}

	client := resty.New()

	domain := os.Getenv("NISHTAGRAM_DOMAIN")
	if domain == "" {
		domain = "127.0.0.1"
	}
	if os.Getenv("DOCKER_ENV") == "" {
		resp, _ := client.R().
			SetBody(c.Request().Body).
			SetHeader("Authorization", token).
			EnableTrace().
			Post("https://" + domain + ":8080/deleteDisposableCampaign")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}

		return c.JSON(200, "ok")

	} else {
		resp, _ := client.R().
			SetBody(c.Request().Body).
			SetHeader("Authorization", token).
			EnableTrace().
			Post("http://" + domain + ":8093/ad/deleteDisposableCampaign")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}

		return c.JSON(200, "ok")
	}
}

func (c2 campaignHandler) GetAllDisposableCampaigns(c echo.Context) error {

	userId, _ := middleware.ExtractUserId(c.Request())
	token, err := c2.tokenRepository.FetchToken(context.Background(), userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error")
	}

	client := resty.New()
	var retVal []domain.DisposableCampaign
	domain := os.Getenv("NISHTAGRAM_DOMAIN")
	if domain == "" {
		domain = "127.0.0.1"
	}
	if os.Getenv("DOCKER_ENV") == "" {
		resp, err := client.R().
			SetHeader("Authorization", token).
			EnableTrace().
			Get("https://" + domain + ":8080/getAllDisposableCampaigns")
		fmt.Println(err)

		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}
		json.Unmarshal(resp.Body(), &retVal)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, invalidJson)
		}
		return c.JSON(200, retVal)

	} else {
		resp, _ := client.R().
			SetHeader("Authorization", token).
			EnableTrace().
			Get("http://" + domain + ":8093/ad/getAllDisposableCampaigns")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}
		json.Unmarshal(resp.Body(), &retVal)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, invalidJson)
		}
		return c.JSON(200, retVal)
	}
}

func (c2 campaignHandler) GetAllMultipleCampaigns(c echo.Context) error {

	userId, _ := middleware.ExtractUserId(c.Request())
	token, err := c2.tokenRepository.FetchToken(context.Background(), userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error")
	}
	var retVal []domain.MultipleCampaign
	client := resty.New()

	domain := os.Getenv("NISHTAGRAM_DOMAIN")
	if domain == "" {
		domain = "127.0.0.1"
	}
	if os.Getenv("DOCKER_ENV") == "" {
		resp, err := client.R().
			SetHeader("Authorization", token).
			EnableTrace().
			Get("https://" + domain + ":8080/getAllMultipleCampaigns")
		fmt.Println(err)


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}
		json.Unmarshal(resp.Body(), &retVal)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, invalidJson)
		}
		return c.JSON(200, retVal)

	} else {
		resp, _ := client.R().
			SetHeader("Authorization", token).
			EnableTrace().
			Get("http://" + domain + ":8093/ad/getAllMultipleCampaigns")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}
		json.Unmarshal(resp.Body(), &retVal)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, invalidJson)
		}
		return c.JSON(200, retVal)
	}
}

func (c2 campaignHandler) UpdateMultipleCampaign(c echo.Context) error {
	userId, _ := middleware.ExtractUserId(c.Request())
	token, err := c2.tokenRepository.FetchToken(context.Background(), userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error")
	}

	client := resty.New()

	domain := os.Getenv("NISHTAGRAM_DOMAIN")
	if domain == "" {
		domain = "127.0.0.1"
	}
	if os.Getenv("DOCKER_ENV") == "" {
		resp, _ := client.R().
			SetBody(c.Request().Body).
			SetHeader("Authorization", token).
			EnableTrace().
			Post("https://" + domain + ":8080/updateMultipleCampaign")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}

		return c.JSON(200, "ok")

	} else {
		resp, _ := client.R().
			SetBody(c.Request().Body).
			SetHeader("Authorization", token).
			EnableTrace().
			Post("http://" + domain + ":8093/ad/updateMultipleCampaign")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}

		return c.JSON(200, "ok")
	}
}

func (c2 campaignHandler) DeleteMultipleCampaign(c echo.Context) error {
	userId, _ := middleware.ExtractUserId(c.Request())
	token, err := c2.tokenRepository.FetchToken(context.Background(), userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error")
	}

	client := resty.New()

	domain := os.Getenv("NISHTAGRAM_DOMAIN")
	if domain == "" {
		domain = "127.0.0.1"
	}
	if os.Getenv("DOCKER_ENV") == "" {
		resp, _ := client.R().
			SetBody(c.Request().Body).
			SetHeader("Authorization", token).
			EnableTrace().
			Post("https://" + domain + ":8080/deleteMultipleCampaign")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}

		return c.JSON(200, "ok")

	} else {
		resp, _ := client.R().
			SetBody(c.Request().Body).
			SetHeader("Authorization", token).
			EnableTrace().
			Post("http://" + domain + ":8093/ad/deleteMultipleCampaign")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}

		return c.JSON(200, "ok")
	}
}

func (c2 campaignHandler) SaveToken(c echo.Context) error {
	var token *domain.AdminToken

	decoder := json.NewDecoder(c.Request().Body)
	err := decoder.Decode(&token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, invalidJson)
	}
	token.UserId, _ = middleware.ExtractUserId(c.Request())
	err = c2.tokenRepository.SaveToken(context.Background(), token.Token, token.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, invalidCreds)
	}
	return c.JSON(http.StatusOK, "ok")

}

func (c2 campaignHandler) CreateAd(c echo.Context) error {

	userId, _ := middleware.ExtractUserId(c.Request())
	token, err := c2.tokenRepository.FetchToken(context.Background(), userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error")
	}

	client := resty.New()

	domain := os.Getenv("NISHTAGRAM_DOMAIN")
	if domain == "" {
		domain = "127.0.0.1"
	}
	if os.Getenv("DOCKER_ENV") == "" {
		resp, _ := client.R().
			SetBody(c.Request().Body).
			SetHeader("Authorization", token).
			EnableTrace().
			Post("https://" + domain + ":8080/createAd")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}

		return c.JSON(200, "ok")

	} else {
		resp, _ := client.R().
			SetBody(c.Request().Body).
			SetHeader("Authorization", token).
			EnableTrace().
			Post("http://" + domain + ":8093/ad/createAd")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}

		return c.JSON(200, "ok")
	}
}

func (c2 campaignHandler) CreateDisposableCampaign(c echo.Context) error {

	userId, _ := middleware.ExtractUserId(c.Request())
	token, err := c2.tokenRepository.FetchToken(context.Background(), userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error")
	}

	client := resty.New()

	domain := os.Getenv("NISHTAGRAM_DOMAIN")
	if domain == "" {
		domain = "127.0.0.1"
	}
	if os.Getenv("DOCKER_ENV") == "" {
		resp, _ := client.R().
			SetBody(c.Request().Body).
			SetHeader("Authorization", token).
			EnableTrace().
			Post("https://" + domain + ":8080/createDisposableCampaign")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}

		return c.JSON(200, "ok")

	} else {
		resp, _ := client.R().
			SetBody(c.Request().Body).
			SetHeader("Authorization", token).
			EnableTrace().
			Post("http://" + domain + ":8093/ad/createDisposableCampaign")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}

		return c.JSON(200, "ok")
	}
}

func (c2 campaignHandler) CreateMultipleCampaign(c echo.Context) error {

	userId, _ := middleware.ExtractUserId(c.Request())
	token, err := c2.tokenRepository.FetchToken(context.Background(), userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error")
	}

	client := resty.New()

	domain := os.Getenv("NISHTAGRAM_DOMAIN")
	if domain == "" {
		domain = "127.0.0.1"
	}
	if os.Getenv("DOCKER_ENV") == "" {
		resp, _ := client.R().
			SetBody(c.Request().Body).
			SetHeader("Authorization", token).
			EnableTrace().
			Post("https://" + domain + ":8080/createMultipleCampaign")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}

		return c.JSON(200, "ok")

	} else {
		resp, _ := client.R().
			SetBody(c.Request().Body).
			SetHeader("Authorization", token).
			EnableTrace().
			Post("http://" + domain + ":8093/ad/createMultipleCampaign")


		if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusInternalServerError, "error")
		}

		return c.JSON(200, "ok")
	}
}

func NewCampaignHandler(tokenRepo datastore.TokenRepository, reportUsecase usecase.ReportUseCase) CampaignHandler {
	return &campaignHandler{tokenRepository: tokenRepo, reportUseCase: reportUsecase}
}